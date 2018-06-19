package models

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/db"
)

type TimeWiseData struct {
	Loc   []float64 `bson:"loc"`
	Gpstm string    `bson:"gpstm"`
	Actvt string    `bson:"actvt"`
	Speed int       `bson:"speed"`
	Sertm time.Time `bson:"sertm"`
}

func GetDateData(vehid string, date string, _sn *mgo.Session) []TimeWiseData {

	if _sn == nil {
		_sn = getDBSession().Copy()
		defer _sn.Close()
	}

	c := col(_sn, db.ColVhtrps)

	//Parameters
	//"2017-11-07T00:00:00+05:30"
	frmt, err := time.Parse(time.RFC3339, date)
	//fmt.Println(frmt)
	tot := frmt.AddDate(0, 0, 1)
	//fmt.Println(tot)
	if err != nil {
		fmt.Println("error")
	}

	var dResult1 []TimeWiseData
	// fmt.Println("this ", frmt)

	_ = c.Find(bson.M{"vhid": vehid, "sertm": bson.M{"$gt": frmt, "$lt": tot}, "actvt": "loc"}).Sort("_id").All(&dResult1)
	return dResult1
}
