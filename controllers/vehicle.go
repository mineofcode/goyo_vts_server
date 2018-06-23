package controllers

import (
	"goyo.in/gpstracker/datamodel"
	"encoding/json"

	"goyo.in/gpstracker/protocal"

	"github.com/astaxie/beego"
	"goyo.in/gpstracker/models"
	"goyo.in/gpstracker/utils"
)

// Operations about login

type VehicleController struct {
	beego.Controller
}

type vhmod struct {
	Vhid     string `json:"vhid"`
	Alwspeed int    `json:"alwspeed"`
	UID string `json:"uid"`
}

// Post Vehicle

func (o *VehicleController) Post() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	var ob interface{}
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	var response datamodel.Vehicles
	_ = json.Unmarshal(o.Ctx.Input.RequestBody, &response)

	result, ip := models.UpdateVehicleData(ob, response.Vhid)

	if ip != "" {
		if response.Alwspeed != 0 {
			protocalHandler.UpdateAllowSpeed(response.Alwspeed, ip)
		}
	}

	o.Data["json"] = utils.CreateWrap("200", result)

	o.ServeJSON()
}

// Activate Vehicle

func (o *VehicleController) ActivateVehicle() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	var ob interface{}
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	var response vhmod

	_ = json.Unmarshal(o.Ctx.Input.RequestBody, &response)
	actres := models.ActivateVehicleData(ob, response.Vhid)

	o.Data["json"] = utils.CreateWrap("200", actres)

	o.ServeJSON()
}

// Get Vehicle By User ID

func (o *VehicleController) GetVehicleByUID() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	var ob interface{}
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	var response vhmod

	_ = json.Unmarshal(o.Ctx.Input.RequestBody, &response)
	actres, _:= models.GetVehicleByUID(response.UID)

	o.Data["json"] = utils.CreateWrap("200", actres)

	o.ServeJSON()
}
