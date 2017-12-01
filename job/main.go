package job

import (
	"time"

	"github.com/jasonlvhit/gocron"
	"goyo.in/gpstracker/dataprocess"
	"goyo.in/gpstracker/models"
)

func StartJob() {

	gocron.Every(1).Day().At("00:10").Do(dailyHistroy, nil)

	gocron.Start()

}

func DailyData(search interface{}) {
	dailyHistroy(search)
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
			frmt, _ := time.Parse(time.RFC3339, "2017-11-08T00:00:00+05:30")
			vhs.Histtm = frmt
		}

		// increase by day one
		vhs.Histtm = vhs.Histtm.AddDate(0, 0, 1)

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

		if data.Segments != nil {
			// Add history data to mongodb
			dataprocess.AddHistoryData(data, _sn)
		}
		// update last updated date to vehicle collection
		dataprocess.UpdateVehiclesHistoryDate(vhs.VhId, vhs.Histtm, _sn)

		// fmt.Println(vhs.Histtm)
		// fmt.Println(data)
	}

}
