package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"goyo.in/gpstracker/models"
	"goyo.in/gpstracker/utils"
)

type FuelController struct {
	beego.Controller
}

func (this *FuelController) Add() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	var ob map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

	//this.Ctx.Input.Bind(&imei, "imei")
	//fmt.Println(imei)
	result := models.SaveFuelEntry(ob)
	this.Data["json"] = utils.CreateWrap("200", result)
	this.ServeJSON()
}

func (this *FuelController) Get() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	var ob map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	//this.Ctx.Input.Bind(&imei, "imei")
	//fmt.Println(imei)
	ob["flag"] = "all"
	result := models.GetFuelEntry(ob)
	this.Data["json"] = utils.CreateWrap("200", result)
	this.ServeJSON()
}

func (this *FuelController) GetEdit() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	var ob map[string]interface{}
	json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
	ob["flag"] = "edit"
	result := models.GetFuelEntry(ob)
	this.Data["json"] = utils.CreateWrap("200", result)
	this.ServeJSON()
}
