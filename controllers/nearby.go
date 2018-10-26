package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"goyo.in/gpstracker/models"
	"goyo.in/gpstracker/utils"
)

type LocationsController struct {
	beego.Controller
}

func (this *LocationsController) GetNearByVehicles() {

	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	var ob map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	actres := models.GetNearByVehicles(ob)
	this.Data["json"] = utils.CreateWrap("200", actres)
	this.ServeJSON()
}
