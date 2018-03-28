package job

import (
	"fmt"
	"time"

	"github.com/jasonlvhit/gocron"
	"goyo.in/gpstracker/dataprocess"
	"goyo.in/gpstracker/models"
	// "github.com/kellydunn/golang-geo"
)

func StartJob() {
	//geofence := NewGeofence([][]*geo.Point{polygon, holes}, int64(20))

	gocron.Every(1).Day().At("00:10").Do(dailyHistroy_SCH)
	<-gocron.Start()

}

func DailyData(search interface{}) {
	dailyHistroy(search)
}

func dailyHistroy_SCH() {
	dailyHistroy(nil)
}

func dailyHistroy(search interface{}) {

	//get single mongodb session
	_sn := models.GetSession().Copy()
	defer _sn.Close()

	// get vehicle list
	res, _ := dataprocess.GetVehicles(search, _sn)
	// loop for vehicles
	//create history
	for _, vhs := range res {

		// checking for hostory
		if vhs.Histtm.Year() == 1 {
			frmt, _ := time.Parse(time.RFC3339, "2017-01-08T00:00:00+05:30")
			vhs.Histtm = frmt
		}

		// increase by day one
		//

		// set d to starting date and keep adding 1 day to it as long as month doesn't change
		//return
		for vhs.Histtm = vhs.Histtm.AddDate(0, 0, 1); vhs.Histtm.Before(time.Now().AddDate(0, 0, -1)); vhs.Histtm = vhs.Histtm.AddDate(0, 0, 1) {
			// do stuff with d
			fmt.Println(vhs.Histtm)

			if vhs.Histtm.Equal(time.Now()) || vhs.Histtm.After(time.Now()) {
				continue
			}

			//setting layout of date to get history record
			const layout = "2006-01-02T00:00:00"
			params := models.ParamsTripHistorydata{
				FromDt: vhs.Histtm.Format(layout) + "+05:30",
				Vhid:   vhs.VhId,
			}
			// get date wise history data
			data, _ := dataprocess.FungetHistoryData(params, _sn)

			//if data.Segments != nil {
			// Add history data to mongodb
			dataprocess.AddHistoryData(data, _sn)
			//}
			// update last updated date to vehicle collection
			dataprocess.UpdateVehiclesHistoryDate(vhs.VhId, vhs.Histtm, _sn)
		}
		// fmt.Println(vhs.Histtm)
		// fmt.Println(data)
	}

}
