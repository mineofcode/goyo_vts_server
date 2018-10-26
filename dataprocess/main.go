package dataprocess

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/models"
)

func GetHistoryData(param models.ParamsTripHistorydata, _sn *mgo.Session) (result datamodel.SegmentWrapper, err error) {

	// data, _ := models.GetStoredHistory(param.Vhid, param.FromDt, _sn)
	// if len(data.Segments) > 0 {
	// 	return data, nil
	// }

	return FungetHistoryData(param, _sn)
}

var _acc int = -1
var _acc_lastTime time.Time
var _accsegment []datamodel.AccSegmentArr

//
func FungetHistoryData(param models.ParamsTripHistorydata, _sn *mgo.Session) (result datamodel.SegmentWrapper, err error) {
	_acc = -1
	_acc_lastTime = time.Now()
	_accsegment = nil
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
	forDate = forDate.UTC()
	// var _loc []locdt

	//non traveled variable
	//var _nont_startTim time.Time
	var _nont_end_Lat []float64
	var switcher string = ""
	//////////////////////////////////////////////////////
	// acc details

	var _segment []datamodel.SegmentArr

	for index, res1 := range res {

		serTm := res1.Sertm.UTC()

		//fmt.Println(serTm)
		if res1.Actvt == "loc" {

			if res1.Flag == "acc" {
				generateAccSegment(res1, forDate, false)
				continue
			}

			if lastLat == nil {
				lastLat = res1.Loc
				//frmt, _ := time.Parse(time.RFC3339, param.FromDt)
				starttm = serTm
				lasttm = serTm

			}
			if maxSpeed < res1.Speed {
				maxSpeed = res1.Speed
				maxSpeedLoc = res1.Loc
				maxSpeedTM = serTm

				if maxSpeed > total_MaxSpeed {
					total_MaxSpeed = maxSpeed
				}
			}
			//fmt.Println(serTm, lasttm)
			isTMseg := isDifferentSegmentTime(serTm, lasttm)
			isDistseg := isDifferentSegmentDist(lastLat, res1.Loc)
			//fmt.Println("dt ", serTm)
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
							Distance:    dis / 1000,
							Duration:    out.Format("15:04:05"),
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

						_dur := serTm.Sub(lasttm)
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
							EndTm:       serTm,
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
					//fmt.Println("reset date ", serTm)
					//_loc = []locdt{}
					polyGonPoints = [][]float64{{}}
				}
			}

			polyGonPoints = append(polyGonPoints, []float64{res1.Loc[1], res1.Loc[0]})
			//_nont_end_Lat = []float64{res1.Loc[1], res1.Loc[0]}
			//_loc = append(_loc, l) //append locations
			//fmt.Println(serTm)

			if lastLat != nil {
				dis += getDistance(res1.Loc, lastLat)
				lastLat = res1.Loc
				lasttm = serTm

				//fmt.Println(lasttm)
			}

		}

	}
	if datacount > 0 {
		generateAccSegment(res[datacount], forDate, true)
	}
	out1 := time.Time{}.Add(movingDuration)

	avgSped := int((total_distance / 1000) / movingDuration.Hours())
	//fmt.Println(out1.Format("15:04:05"))
	_final_wrap := datamodel.SegmentWrapper{MaxSpeed: total_MaxSpeed,
		Segments:     _segment,
		ACCSegments:  _accsegment,
		ToalDistance: total_distance / 1000,
		TravelTime:   out1.Format("15:04:05"),
		AvgSpeed:     avgSped,
		Vhid:         param.Vhid,
		Date:         forDate,
		AccAvail:     len(_accsegment) > 0,
	}
	return _final_wrap, nil
}

func generateAccSegment(res1 models.TimeWiseData, forDate time.Time, isLast bool) {
	// fmt.Println("acc")

	typ := "stop"
	in_acc := 1
	iscontinue := false
	if _acc == -1 {
		_acc_lastTime = forDate
	}
	if isLast {
		if res1.Acc == 0 {
			in_acc = 1
		} else {
			in_acc = 0
		}

	} else {
		in_acc = res1.Acc
	}

	serverTime := res1.Sertm.UTC()
	// check if this is last record of array then set the data
	if isLast == true {
		if len(_accsegment) > 0 {

			// 	serverEndTime := forDate.Add(time.Hour*time.Duration(23) +
			// 		time.Minute*time.Duration(59) +
			// 		time.Second*time.Duration(59))
			// setting last segment value
			lastsegment := _accsegment[len(_accsegment)-1]

			serverTime = forDate.Add(time.Hour*time.Duration(23) +
				time.Minute*time.Duration(59) +
				time.Second*time.Duration(59))
			if lastsegment.Type == "start" {
				if time.Now().Before(serverTime) {
					serverTime = time.Now().UTC()
				} else {

					iscontinue = true
				}
			} else {
				durl := serverTime.Sub(lastsegment.StartTm)
				outl := time.Time{}.Add(durl)

				lastsegment.Duration = outl.Format("15:04:05")
				lastsegment.EndTm = serverTime
				_accsegment[len(_accsegment)-1] = lastsegment
				return
			}
		}

		// 	durl := serverEndTime.Sub(lastsegment.StartTm)
		// 	outl := time.Time{}.Add(durl)

		// 	lastsegment.Duration = outl.Format("15:04:05")
		// 	lastsegment.EndTm = serverEndTime

		// 	_accsegment[len(_accsegment)-1] = lastsegment
		// } else {

	}
	//return

	if in_acc == 1 {
		typ = "stop"
	} else {
		typ = "start"
	}

	//fmt.Println(lastsegment.Type)
	// data will start from this day morning 12:00 am to till res.Time
	dur := serverTime.Sub(_acc_lastTime)
	outx := time.Time{}.Add(dur)

	sg := datamodel.AccSegmentArr{
		Duration:   outx.Format("15:04:05"),
		StartTm:    _acc_lastTime,
		EndTm:      serverTime,
		Point:      res1.Loc,
		Type:       typ,
		IsContinue: iscontinue,
	}
	_acc_lastTime = serverTime
	_acc = in_acc
	_accsegment = append(_accsegment, sg)

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
