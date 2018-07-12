package models

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/db"
	"goyo.in/gpstracker/utils"
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

	c := col(_sn, db.ColVhcls)
	var dResult1 []Vhdata

	_ = c.Find(search).All(&dResult1)

	return dResult1
}

func GetVehicle(vhid string, _sn *mgo.Session) []datamodel.Vehicles {
	if _sn == nil {
		_sn = getDBSession().Copy()
		defer _sn.Close()
	}

	c := col(_sn, db.ColVhcls)
	var dResult1 []datamodel.Vehicles

	_ = c.Find(bson.M{"vhid": vhid}).Limit(1).All(&dResult1)
	return dResult1
}

func UpdateVehiclesHistoryDate(vhid string, histm time.Time, _sn *mgo.Session) {
	if _sn == nil {
		_sn = getDBSession().Copy()
		defer _sn.Close()
	}

	c := col(_sn, db.ColVhcls)

	if dberr := c.Update(bson.M{"vhid": vhid}, bson.M{"$set": bson.M{"histtm": histm}}); dberr != nil {
		fmt.Println(dberr)
	}
}

func AddVehiclesHistoryDate(histrywrp datamodel.SegmentWrapper, _sn *mgo.Session) {
	if _sn == nil {
		_sn = getDBSession().Copy()
		defer _sn.Close()
	}

	c := col(_sn, db.ColHistory)
	if dberr, _ := c.Upsert(bson.M{"vhid": histrywrp.Vhid, "date": histrywrp.Date}, histrywrp); dberr != nil {
		fmt.Println(dberr)
	}
}

func GetStoredHistory(vhid string, histm string, _sn *mgo.Session) (result datamodel.SegmentWrapper, err error) {
	if _sn == nil {
		_sn = getDBSession().Copy()
		defer _sn.Close()
	}

	c := col(_sn, db.ColHistory)

	frmt, err := time.Parse(time.RFC3339, histm)

	if err != nil {
		fmt.Println("error")
	}

	var dR datamodel.SegmentWrapper
	err = c.Find(bson.M{"vhid": vhid, "date": frmt}).One(&dR)

	return dR, err
}

// Get Vehicle By User ID

type GetVehicelListByUID struct {
	UID    string      `bson:"uid" json:"uid"`
	VtsId  int         `bson:"vtsid" json:"vtsid"`
	Vhid   string      `bson:"vhid" json:"vhid"`
	VhNm   string      `bson:"vhname" json:"vno"`
	VhIcon string      `bson:"vhtyp" json:"ico"`
	Vhd    interface{} `bson:"vhd" json:"vhd"`
}

func GetVehicleByUID(uid string) (result []GetVehicelListByUID, err error) {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	c := col(_sn, db.ColVhcls)

	if err != nil {
		fmt.Println("error")
	}

	var dR []GetVehicelListByUID

	err = c.Find(bson.M{"uid": uid}).All(&dR)
	return dR, err
}

func GetVehicleByID(params map[string]interface{}) (result datamodel.Vehicles, err error) {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	c := col(_sn, db.ColVhcls)

	if err != nil {
		fmt.Println("error")
	}

	var dR datamodel.Vehicles

	err = c.Find(bson.M{"vtsid": params["vtsid"]}).One(&dR)
	return dR, err
}

// Activate Vehicle Data

func ActivateVehicleData(d map[string]interface{}, vhid string) (utils.Response, string) {
	_sn := getDBSession().Copy()
	defer _sn.Close()
	var err error

	resp := CheckDeviceActivate(vhid)

	if !resp.Status && d["vtsid"] == 0 {
		return resp, ""
	}

	c := col(_sn, db.ColVhcls)

	if dberr := c.Update(bson.M{"vhid": vhid}, bson.M{"$set": d}); dberr != nil {
		resp.Error = dberr.Error()
		resp.Status = false

		if dberr.Error() == "not found" {
			d["vtsid"] = GetNextSequence(_sn, SEQVehicleID)
			fmt.Println(d)
			err = c.Insert(d)

			if err != nil {
				resp.Error = err.Error()
				resp.Status = false
				return resp, ""
			}

			resp.Message = "Added Successfully"
			resp.Status = true
		}
	} else {
		var vh Vhdata

		c.Find(bson.M{"vhid": vhid}).One(&vh)

		resp.Message = "Updated Successfully"
		resp.Status = true

		return resp, vh.Ip
	}

	// var vh ActiVateResult
	// c.Find(bson.M{"vhid": vhid}).One(&vh)

	return resp, ""
}

//update Insert vehicledata

func UpdateVehicleData(d map[string]interface{}, vhid interface{}) (result string, ipaddr string) {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	var err error
	var ip string

	c := col(_sn, db.ColVhcls)

	if dberr := c.Update(bson.M{"vhid": vhid}, bson.M{"$set": d}); dberr != nil {
		result = dberr.Error()

		if dberr.Error() == "not found" {
			d["vtsid"] = GetNextSequence(_sn, SEQVehicleID)

			fmt.Println(d)
			err = c.Insert(d)
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

//update Insert vehicledata

func GetVehicleIP(vhid interface{}) (ipaddr string) {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	var ip string

	c := col(_sn, db.ColVhcls)
	var vh Vhdata

	c.Find(bson.M{"vhid": vhid}).One(&vh)

	ip = vh.Ip

	return ip
}

type VhLoginData struct {
	AllowSpd int       `bson:"alwspeed"`
	VhNm     string    `bson:"vhname"`
	LstSpdtm time.Time `bson:"lstspdtm"`
	VtsID    int       `bson:"vtsid"`
	PClients []string  `bson:"pushcl"`
	AC       int       `bson:"D1"`
	ACC      int       `bson:"acc"`
	SerTm    time.Time `bson:"sertm"`
	Loc      []float64 `bson:"loc"`
}

func GetVehiclesData(vhid string) VhLoginData {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	c := col(_sn, db.ColVhcls)
	var dResult1 VhLoginData

	_ = c.Find(bson.M{"vhid": vhid}).One(&dResult1)
	return dResult1
}
