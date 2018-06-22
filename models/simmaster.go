package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/db"
	"goyo.in/gpstracker/utils"
)

//DeviceCommands model
type simDetails struct {
	Imei         string    `bson:"imei" 	json:"imei"`
	SimNumber    string    `bson:"sim" 		json:"sim"`
	UID          string    `bson:"uid" 		json:"uid"`
	Ucode        string    `bson:"ucode" 	json:"ucode"`
	Platform     string    `bson:"platform" json:"platform"`
	Src          string    `bson:"src" 		json:"src"`
	IP           string    `bson:"ip" 		json:"ip"`
	Time         time.Time `bson:"time" 	json:"time"`
	Resultraw    string    `bson:"resultraw" json:"resultraw"`
	Stringresult string    `bson:"result" 	json:"result"`
	DeviceID     string    `bson:"devid" 	json:"devid"`
}

//check device exists and activated
func CheckSIMActivate(mobno string) utils.Response {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	resp := utils.Response{}

	//Check imei in inventory
	ColSimEnv := col(_sn, db.ColSimEnv)
	var sIMMaster datamodel.SIMMaster
	ColSimEnv.Find(bson.M{"mobno": mobno}).One(&sIMMaster)

	// if imei not found
	if sIMMaster.MobNo == "" {
		resp.Error = "Invalid Sim Number!!! Please contact administrator."
		resp.Status = false
		resp.StausCode = 0
		return resp
	}
	//Check imei in vehicles
	// ColVhcls := col(_sn, db.ColVhcls)
	// var vehicles datamodel.Vehicles
	// ColVhcls.Find(bson.M{"imei": imei}).One(&vehicles)

	// if vehicles.Vhid != "" && vehicles.UID != "" {
	// 	resp.Error = "This device is already registered!!! Please contact administrator."
	// 	resp.Status = false
	// 	resp.StausCode = 0
	// 	return resp
	// }

	resp.Error = "This device is already registered!!! Please contact administrator."
	resp.Status = false
	resp.StausCode = 0
	return resp

}

//check device exists and activated
func CreateSIMEntry(data datamodel.SIMMaster) utils.Response {
	_sn := getDBSession().Copy()
	defer _sn.Close()
	data.CreateOn = time.Now()

	validerr := GetValidator().Struct(data)

	resp := utils.Response{}

	if validerr != nil {
		resp.Error = validerr.Error()
		resp.Status = false
		resp.StausCode = 0
		return resp
	}

	//Check imei in inventory
	ColSimEnv := col(_sn, db.ColSimEnv)

	err := ColSimEnv.Insert(data)

	// if imei not found
	if err != nil {
		resp.Error = err.Error()
		resp.Status = false
		resp.StausCode = 0
		return resp
	}
	//Check imei in vehicles
	// ColVhcls := col(_sn, db.ColVhcls)
	// var vehicles datamodel.Vehicles
	// ColVhcls.Find(bson.M{"imei": imei}).One(&vehicles)

	// if vehicles.Vhid != "" && vehicles.UID != "" {
	// 	resp.Error = "This device is already registered!!! Please contact administrator."
	// 	resp.Status = false
	// 	resp.StausCode = 0
	// 	return resp
	// }

	resp.Message = "SIM addedd successfully"
	resp.Status = true
	resp.StausCode = 0
	return resp

}

func GetSimDetails(mobno string) utils.Response {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	resp := utils.Response{}

	//Check imei in inventory
	ColSimEnv := col(_sn, db.ColSimEnv)
	var sIMMaster datamodel.SIMMaster
	ColSimEnv.Find(bson.M{"mobno": mobno}).One(&sIMMaster)

	// if imei not found
	if sIMMaster.MobNo == "" {
		resp.Error = "Invalid Sim Number!!! Please contact administrator."
		resp.Status = false
		resp.StausCode = 0
		return resp
	}

	resp.Data = sIMMaster

	resp.Status = true
	resp.StausCode = 0
	return resp

}
