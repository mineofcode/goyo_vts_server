package models

import (
	"fmt"
	"reflect"
	"time"

	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/db"
)

type MilageReport struct {
	travel_tm      string `bson:"travel_tm" json:"travel_tm"`
	total_distance string `bson:"total_distance" json:"total_distance"`
	date           string `bson:"date" json:"date"`
}

func GetReport(trpparams datamodel.ReportParams) (ret []interface{}, err error) {

	// if err != nil {
	// 	panic(err)
	// }
	switch trpparams.Reporttyp {
	case "milege":
		return getMilageReport(trpparams.Params)
	case "speed":
		return getSpeedReport(trpparams.Params)
	}

	return nil, nil
}

type milageField struct {
	vhids string `bson:"vhid" json:"vhid"`
}

func getMilageReport(params interface{}) (ret []interface{}, err error) {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	var result []interface{}
	c := col(_sn, db.ColHistory)
	param := reflect.ValueOf(params)

	//fmt.Println(param.MapIndex(reflect.ValueOf("abc")).IsValid())

	// if val, ok := params["foo"]; ok {
	// 	//do something here
	// 	fmt.Println(val)
	// }
	vhid := param.MapIndex(reflect.ValueOf("vhid")).Interface()
	frmdt := param.MapIndex(reflect.ValueOf("frmdt")).Interface()
	todt := param.MapIndex(reflect.ValueOf("todate")).Interface()
	prtype_map := param.MapIndex(reflect.ValueOf("type"))
	var prtype interface{} = ""
	if prtype_map.IsValid() {
		prtype = prtype_map.Interface()
	}

	fd, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", frmdt))
	tod, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", todt))
	// fmt.Println(frmdt)
	// fmt.Println(todt)

	//_vh := trpparams.Vhids //strings.Split(trpparams.Vhids, ",")
	//err = c.Find(bson.M{"vhid": vhid, "$group": bson.M{}}).All(&result)
	//group by
	var groupby interface{}
	if prtype == "datewise" {
		groupby = bson.M{
			"vhid": "$vhid",
			"date": "$date",
		}
	} else {
		groupby = bson.M{
			"vhid": "$vhid",
		}
	}

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{"vhid": bson.M{"$in": vhid},
				"date": bson.M{"$gte": fd, "$lte": tod}},
		},
		bson.M{
			"$group": bson.M{
				"_id":     groupby,
				"avgspd":  bson.M{"$avg": "$avg_spd"},
				"maxspd":  bson.M{"$max": "$mx_spd"},
				"frmdate": bson.M{"$first": "$date"},
				"todate":  bson.M{"$last": "$date"},
				"milege":  bson.M{"$sum": "$total_distance"},
			},
		},

		bson.M{
			"$lookup": bson.M{
				"from":         db.ColVhcls,
				"localField":   "_id.vhid",
				"foreignField": "vhid",
				"as":           "vhname",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":           1,
				"avgspd":        1,
				"maxspd":        1,
				"frmdate":       1,
				"todate":        1,
				"milege":        1,
				"vhname.vhname": 1,
			},
		},
		bson.M{
			"$sort": bson.M{
				"_id":     1,
				"frmdate": 1,
			},
		},
	}
	err = c.Pipe(pipeline).All(&result)
	//process on data

	return result, err
}

func getSpeedReport(params interface{}) (ret []interface{}, err error) {

	param := reflect.ValueOf(params)
	prtypemap := param.MapIndex(reflect.ValueOf("type"))
	var prtype interface{} = ""
	if prtypemap.IsValid() {
		prtype = prtypemap.Interface()
	}

	if prtype == "details" {

		return getSpeedDetailsReport(params)
	}

	_sn := getDBSession().Copy()
	defer _sn.Close()

	var result []interface{}
	c := col(_sn, db.ColVhtrps)

	//fmt.Println(param.MapIndex(reflect.ValueOf("abc")).IsValid())

	// if val, ok := params["foo"]; ok {
	// 	//do something here
	// 	fmt.Println(val)
	// }
	vhid := param.MapIndex(reflect.ValueOf("vhid")).Interface()
	frmdt := param.MapIndex(reflect.ValueOf("frmdt")).Interface()
	todt := param.MapIndex(reflect.ValueOf("todate")).Interface()

	fd, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", frmdt))
	tod, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", todt))
	// fmt.Println(frmdt)
	// fmt.Println(todt)

	//_vh := trpparams.Vhids //strings.Split(trpparams.Vhids, ",")
	//err = c.Find(bson.M{"vhid": vhid, "$group": bson.M{}}).All(&result)
	//group by
	var groupby interface{}
	groupby = bson.M{
		"vhid": "$vhid",
		"date": bson.M{"$dateToString": bson.M{"format": "%Y-%m-%d", "date": "$sertm"}},
	}

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{"vhid": bson.M{"$in": vhid},
				"sertm": bson.M{"$gte": fd, "$lte": tod},
				"isp":   true,
			},
		},
		bson.M{
			"$group": bson.M{
				"_id":   groupby,
				"count": bson.M{"$sum": 1},
			},
		},

		bson.M{
			"$lookup": bson.M{
				"from":         db.ColVhcls,
				"localField":   "_id.vhid",
				"foreignField": "vhid",
				"as":           "vhname",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":           1,
				"count":         1,
				"vhname.vhname": 1,
			},
		},
		bson.M{
			"$sort": bson.M{
				"_id.vhid": 1,
				"sertm":    1,
			},
		},
	}
	err = c.Pipe(pipeline).All(&result)
	//process on data

	return result, err
}

//Get vehicle speed report
func getSpeedDetailsReport(params interface{}) (ret []interface{}, err error) {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	var result []interface{}
	c := col(_sn, db.ColVhtrps)
	param := reflect.ValueOf(params)

	//fmt.Println(param.MapIndex(reflect.ValueOf("abc")).IsValid())

	// if val, ok := params["foo"]; ok {
	// 	//do something here
	// 	fmt.Println(val)
	// }
	vhid := param.MapIndex(reflect.ValueOf("vhid")).Interface()
	frmdt := param.MapIndex(reflect.ValueOf("frmdt")).Interface()
	todt := param.MapIndex(reflect.ValueOf("todate")).Interface()

	fd, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", frmdt))
	tod, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", todt))

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{"vhid": bson.M{"$in": vhid},
				"sertm": bson.M{"$gte": fd, "$lte": tod},
				"isp":   true,
			},
		},
		// bson.M{
		// 	"$lookup": bson.M{
		// 		"from":         db.ColVhcls,
		// 		"localField":   "_id.vhid",
		// 		"foreignField": "vhid",
		// 		"as":           "vhname",
		// 	},
		// },
		bson.M{
			"$project": bson.M{
				"_id":      0,
				"vhid":     1,
				"alwspeed": 1,
				"speed":    1,
				"sertm":    bson.M{"$dateToString": bson.M{"format": "%Y-%m-%d %H:%M:%S", "date": "$sertm"}},
			},
		},
		bson.M{
			"$sort": bson.M{
				"sertm": 1,
			},
		},
	}
	err = c.Pipe(pipeline).All(&result)
	//process on data

	return result, err
}
