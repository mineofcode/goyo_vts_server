package models

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/datamodel"
	patterns "goyo.in/gpstracker/patterns"
)

type Vhdata struct {
	VhId   string    `bson:"vhid"`
	Histtm time.Time `bson:"histtm"`
	Ip     string    `bson:"ip"`
}

func GetVehicles(search interface{}, _sn *mgo.Session) []Vhdata {
	if _sn == nil {
		_sn = getDBSession().Copy()
		defer _sn.Close()
	}

	c := col(_sn, patterns.ColVhcls)
	var dResult1 []Vhdata
	// fmt.Println("this ", frmt)

	_ = c.Find(search).All(&dResult1)
	return dResult1
}

func UpdateVehiclesHistoryDate(vhid string, histm time.Time, _sn *mgo.Session) {
	if _sn == nil {
		_sn = getDBSession().Copy()
		defer _sn.Close()
	}

	c := col(_sn, patterns.ColVhcls)
	if dberr := c.Update(bson.M{"vhid": vhid}, bson.M{"$set": bson.M{"histtm": histm}}); dberr != nil {
		fmt.Println(dberr)
	}
}

func AddVehiclesHistoryDate(histrywrp datamodel.SegmentWrapper, _sn *mgo.Session) {
	if _sn == nil {
		_sn = getDBSession().Copy()
		defer _sn.Close()
	}

	c := col(_sn, patterns.ColHistory)
	if dberr, _ := c.Upsert(bson.M{"vhid": histrywrp.Vhid, "date": histrywrp.Date}, histrywrp); dberr != nil {
		fmt.Println(dberr)
	}
}

func GetStoredHistory(vhid string, histm string, _sn *mgo.Session) (result datamodel.SegmentWrapper, err error) {
	if _sn == nil {
		_sn = getDBSession().Copy()
		defer _sn.Close()
	}

	c := col(_sn, patterns.ColHistory)

	//Parameters
	//"2017-11-07T00:00:00+05:30"
	frmt, err := time.Parse(time.RFC3339, histm)
	//fmt.Println(frmt)
	//fmt.Println(tot)
	if err != nil {
		fmt.Println("error")
	}

	var dR datamodel.SegmentWrapper
	// fmt.Println("this ", frmt)

	err = c.Find(bson.M{"vhid": vhid, "date": frmt}).One(&dR)
	return dR, err
}

//update Insert vehicledata
func UpdateVehicleData(d interface{}, vhid interface{}) (result string, ipaddr string) {
	_sn := getDBSession().Copy()
	defer _sn.Close()
	var err error
	var ip string
	//fmt.Println(vhid, d)
	c := col(_sn, patterns.ColVhcls)
	if dberr := c.Update(bson.M{"vhid": vhid}, bson.M{"$set": d}); dberr != nil {
		//fmt.Println(dberr)
		result = dberr.Error()
		if dberr.Error() == "not found" {
			_, err = c.UpsertId(bson.M{"vhid": vhid}, bson.M{"$set": d})
			if err != nil {
				result = err.Error()
			}
			result = "Created Successfully"
		}
	} else {

		var vh Vhdata

		c.Find(bson.M{"vhid": vhid}).One(&vh)

		ip = vh.Ip

		result = "Updated Successfully"
	}
	return result, ip
}

type VhLoginData struct {
	AllowSpd int    `bson:"alwspeed"`
	VhNm     string `bson:"vhname"`
}

func GetVehiclesData(vhid string) VhLoginData {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	c := col(_sn, patterns.ColVhcls)
	var dResult1 VhLoginData
	// fmt.Println("this ", frmt)

	_ = c.Find(bson.M{"vhid": vhid}).One(&dResult1)
	return dResult1
}
