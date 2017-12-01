package dataprocess

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/models"
)

func GetHistoryData(param models.ParamsTripHistorydata, _sn *mgo.Session) (result datamodel.SegmentWrapper, err error) {

	data, _ := models.GetStoredHistory(param.Vhid, param.FromDt, _sn)
	if len(data.Segments) > 0 {
		return data, nil
	}

	return FungetHistoryData(param, _sn)
}

//
func FungetHistoryData(param models.ParamsTripHistorydata, _sn *mgo.Session) (result datamodel.SegmentWrapper, err error) {

	//"358735070927263", "2017-11-06T00:00:00+05:30"
	res := models.GetDateData(param.Vhid, param.FromDt, _sn)

	var polyGonPoints [][]float64

	var lastLat []float64
	var dis float64
	var maxSpeed int
	var maxSpeedLoc []float64
	var maxSpeedTM time.Time
	var movingDuration time.Duration

	var total_MaxSpeed int = 0
	var total_distance float64 = 0
	var starttm time.Time
	var lasttm time.Time
	var datacount = (len(res) - 1)
	forDate, _ := time.Parse(time.RFC3339, param.FromDt)
	// var _loc []locdt

	//non traveled variable
	//var _nont_startTim time.Time
	var _nont_end_Lat []float64
	var switcher string = ""
	//////////////////////////////////////////////////////

	var _segment []datamodel.SegmentArr
	for index, res1 := range res {

		if res1.Actvt == "loc" {

			if lastLat == nil {
				lastLat = res1.Loc
				//frmt, _ := time.Parse(time.RFC3339, param.FromDt)
				starttm = res1.Sertm
				lasttm = res1.Sertm

			}
			if maxSpeed < res1.Speed {
				maxSpeed = res1.Speed
				maxSpeedLoc = res1.Loc
				maxSpeedTM = res1.Sertm

				if maxSpeed > total_MaxSpeed {
					total_MaxSpeed = maxSpeed
				}
			}
			//fmt.Println(res1.Sertm, lasttm)
			isTMseg := isDifferentSegmentTime(res1.Sertm, lasttm)
			isDistseg := isDifferentSegmentDist(lastLat, res1.Loc)
			//fmt.Println("dt ", res1.Sertm)
			//fmt.Println("gpstm ", l.Time)
			// fmt.Println(datacount, index)
			if isDistseg || isTMseg || index == datacount {

				if isDistseg && !isTMseg {

				} else {
					if switcher == "" {
						if len(polyGonPoints) > 2 {
							switcher = "solid"
						}
					}

					if dis/1000 < 1 {
						switcher = "dashed"
					}

					if switcher == "solid" {
						//if len(polyGonPoints) > 1 {
						//Travelled Data
						dur := lasttm.Sub(starttm)
						movingDuration += dur
						out := time.Time{}.Add(dur)
						ploy := EncodeCoords(polyGonPoints)
						sg := datamodel.SegmentArr{
							Distance: dis / 1000, Duration: out.Format("15:04:05"),
							EncodPoly:   fmt.Sprintf("%s", ploy),
							StartTm:     starttm,
							EndTm:       lasttm,
							MaxSpeed:    maxSpeed,
							MaxSpeedLoc: maxSpeedLoc,
							MaxSpeedTM:  maxSpeedTM,
							TrackType:   "solid",
						}
						_segment = append(_segment, sg)
						switcher = "dashed"
					}
					//}
					////////////////////////////////////////////////////////////////////////////////////
					//if len(_segment) >= 1 {
					//Non Traveled Data
					if switcher == "dashed" {
						_nont_end_Lat = []float64{res1.Loc[1], res1.Loc[0]}

						_dur := res1.Sertm.Sub(lasttm)
						_out := time.Time{}.Add(_dur)

						_latlon := [][]float64{
							[]float64{lastLat[1], lastLat[0]},
							_nont_end_Lat,
						}
						//fmt.Println("seg ", lasttm)
						_ploy := EncodeCoords(_latlon)
						_sg_notrav := datamodel.SegmentArr{
							Distance: 0,
							Duration: _out.Format("15:04:05"),
							//Loc:      _loc,
							EncodPoly:   fmt.Sprintf("%s", _ploy),
							StartTm:     lasttm,
							EndTm:       res1.Sertm,
							MaxSpeed:    0,
							MaxSpeedLoc: maxSpeedLoc,
							MaxSpeedTM:  maxSpeedTM,
							TrackType:   "dashed",
						}
						_segment = append(_segment, _sg_notrav)
						switcher = "solid"
					}
					//}
					total_distance += dis
					dis = 0
					maxSpeed = 0
					lastLat = nil
					//fmt.Println("reset date ", res1.Sertm)
					//_loc = []locdt{}
					polyGonPoints = [][]float64{{}}
				}
			}

			polyGonPoints = append(polyGonPoints, []float64{res1.Loc[1], res1.Loc[0]})
			//_nont_end_Lat = []float64{res1.Loc[1], res1.Loc[0]}
			//_loc = append(_loc, l) //append locations
			//fmt.Println(res1.Sertm)

			if lastLat != nil {
				dis += getDistance(res1.Loc, lastLat)
				lastLat = res1.Loc
				lasttm = res1.Sertm

				//fmt.Println(lasttm)
			}

		}
	}
	out1 := time.Time{}.Add(movingDuration)

	avgSped := int((total_distance / 1000) / movingDuration.Hours())
	//fmt.Println(out1.Format("15:04:05"))
	_final_wrap := datamodel.SegmentWrapper{MaxSpeed: total_MaxSpeed,
		Segments:     _segment,
		ToalDistance: total_distance / 1000,
		TravelTime:   out1.Format("15:04:05"),
		AvgSpeed:     avgSped,
		Vhid:         param.Vhid,
		Date:         forDate,
	}
	return _final_wrap, nil
}

//GetVehicles ...
func GetVehicles(search interface{}, _sn *mgo.Session) (result []models.Vhdata, err error) {

	res := models.GetVehicles(search, _sn)

	return res, nil
}

//UpdateVehiclesHistoryDate ...
func UpdateVehiclesHistoryDate(vhid string, histm time.Time, _sn *mgo.Session) {

	models.UpdateVehiclesHistoryDate(vhid, histm, _sn)
}

//AddHistoryData ...
func AddHistoryData(histrywrp datamodel.SegmentWrapper, _sn *mgo.Session) {
	models.AddVehiclesHistoryDate(histrywrp, _sn)
}

// //GetStoredHistory ...
// func GetStoredHistory(vhid string, histm string, _sn *mgo.Session) (result datamodel.SegmentWrapper, err error) {

// 	return models.GetStoredHistory(vhid, histm, _sn)
// }
