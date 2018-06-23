package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/db"
	"goyo.in/gpstracker/reqprops"
	"goyo.in/gpstracker/utils"
)

func AddDevice() {

}

//DeviceCommands model
type deviceDetails struct {
	Imei         string    `bson:"imei" 	json:"imei"`
	SimNumber    string    `bson:"sim" 	json:"sim"`
	UID          string    `bson:"uid" 		json:"uid"`
	Ucode        string    `bson:"ucode" 	json:"ucode"`
	Platform     string    `bson:"platform" json:"platform"`
	Src          string    `bson:"src" json:"src"`
	IP           string    `bson:"ip" 		json:"ip"`
	Time         time.Time `bson:"time" 	json:"time"`
	Resultraw    string    `bson:"resultraw" json:"resultraw"`
	Stringresult string    `bson:"result" 	json:"result"`
	DeviceID     string    `bson:"devid" 	json:"devid"`
}

//check device exists and activated
func CheckDeviceActivate(imei string) utils.Response {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	resp := utils.Response{}

	//Check imei in inventory
	ColDeviceEnv := col(_sn, db.ColDeviceEnv)
	var deviceMaster datamodel.DeviceMaster
	ColDeviceEnv.Find(bson.M{"imei": imei}).One(&deviceMaster)

	// if imei not found
	if deviceMaster.Imei == "" {
		resp.Error = "Invalid Imei Number!!! Please contact administrator."
		resp.Status = false
		resp.StausCode = 0
		return resp
	}
	//Check imei in vehicles
	ColVhcls := col(_sn, db.ColVhcls)
	var vehicles datamodel.Vehicles
	ColVhcls.Find(bson.M{"vhid": imei}).One(&vehicles)

	if vehicles.Vhid != "" && vehicles.UID != "" {
		resp.Error = "This device is already registered!!! Please contact administrator."
		resp.Status = false
		resp.StausCode = 0
		return resp
	}
	resp.Message = "Device is available to register."
	resp.Error = ""
	resp.Status = true
	resp.Data = vehicles
	resp.StausCode = 0
	return resp

}

func DeviceActivation(data reqprops.DeviceActivationProp) utils.Response {
	result := CheckDeviceActivate(data.IMEI)
	resp := utils.Response{}
	_sn := getDBSession().Copy()
	defer _sn.Close()

	if result.Status {
		resp.Status = true

		return resp
	}
	return result

}

//check device exists and activated
func CreateDeviceEntry(data datamodel.DeviceMaster) utils.Response {
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
	ColDeviceEnv := col(_sn, db.ColDeviceEnv)

	err := ColDeviceEnv.Insert(data)

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

	resp.Message = "Device addedd successfully"
	resp.Status = true
	resp.StausCode = 0
	return resp

}

func CreateDeviceBulkEntry(data []datamodel.DeviceMaster) utils.Response {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	validerr := GetValidator().Struct(data)

	resp := utils.Response{}

	if validerr != nil {
		resp.Error = validerr.Error()
		resp.Status = false
		resp.StausCode = 0
		return resp
	}

	//Check imei in inventory
	ColDeviceEnv := col(_sn, db.ColDeviceEnv)

	err := ColDeviceEnv.Insert(data)

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

	resp.Message = "Device addedd successfully"
	resp.Status = true
	resp.StausCode = 0
	return resp

}
