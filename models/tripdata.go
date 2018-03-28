package models

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/datamodel"
	patterns "goyo.in/gpstracker/patterns"
	"goyo.in/gpstracker/socketios"
)

// type GeoJson struct {
// 	Type        string    `json:"-"`
// 	Coordinates []float64 `json:"coordinates"`
// }

type (
	// ParamsTripdata represents the structure of our resource
	ParamsTripdata struct {
		Vhids []string `json:"vhids"`
	}
)

func init() {}

//GetLastStatus of vehicles
func GetLastStatus(trpparams ParamsTripdata) (ret []datamodel.Vehicles, err error) {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	var result []datamodel.Vehicles
	c := col(_sn, patterns.ColVhcls)

	//_vh := trpparams.Vhids //strings.Split(trpparams.Vhids, ",")
	err = c.Find(bson.M{"vhid": bson.M{"$in": trpparams.Vhids}}).All(&result)

	// if err != nil {
	// 	panic(err)
	// }
	return result, err
}

//UpdateData
func UpdateData(d interface{}, vhid string, f string) (err error) {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	//fmt.Println(vhid, d)
	c := col(_sn, patterns.ColVhcls)
	if dberr := c.Update(bson.M{"vhid": vhid}, bson.M{"$set": d}); dberr != nil {
		fmt.Println(dberr)
		if dberr.Error() == "not found" {
			_, err = c.UpsertId(bson.M{"vhid": vhid}, bson.M{"$set": d})
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	//fmt.Println(err)

	// send in history
	if (f == "reg") || (f == "loc") {
		// insert in history table
		ch := col(_sn, patterns.ColVhtrps)
		err = ch.Insert(d)

	}

	// send to socket server
	go func() {
		socket := socketios.GetSocketIO()
		socket.BroadcastTo(vhid, "msgd", bson.M{"evt": "data", "data": d})
	}()

	return err
}

type (
	// ParamsTripdata represents the structure of our resource
	ParamsTripHistorydata struct {
		Vhid   string   `json:"vhid"`
		FromDt string   `json:"frmdt"`
		ToDt   string   `json:"todt"`
		OFlag  string   `json:"flag"`
		Vhids  []string `json:"vhids"`
	}
)

///Over Speed data

func GetOverSpeedDataCount(trpparams ParamsTripHistorydata) (count int, err error) {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	fd, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", trpparams.FromDt))
	tod, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", trpparams.ToDt))
	// fmt.Println(tod)
	//_vh := trpparams.Vhids //strings.Split(trpparams.Vhids, ",")

	c := col(_sn, patterns.ColVhtrps)
	count, err = c.Find(bson.M{"vhid": bson.M{"$in": trpparams.Vhids},
		"sertm": bson.M{"$gte": fd, "$lte": tod},
		"isp":   true}).Count()
	//process on data

	return count, err
}

//
//Get OverSpeed Actual Data

// func GetOverSpeedDataData(trpparams ParamsTripHistorydata) (count int, err error) {
// 	_sn := getDBSession().Copy()
// 	defer _sn.Close()

// 	fd, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", trpparams.FromDt))
// 	tod, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", trpparams.ToDt))

// 	//_vh := trpparams.Vhids //strings.Split(trpparams.Vhids, ",")

// 	c := col(_sn, patterns.ColVhtrps)
// 	count, err = c.Find(bson.M{"vhid": bson.M{"$in": trpparams.Vhids},
// 		"date": bson.M{"$gte": fd, "$lte": tod},
// 		"isp":  true}).Count()
// 	//process on data

// 	return count, err
// }
