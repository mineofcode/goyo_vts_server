package controllers

import (
	"github.com/astaxie/beego"
	"goyo.in/gpstracker/models"
)

type DeviceController struct {
	beego.Controller
}

type RequestParams struct {
	imei string `bson:"imei"`
}

func (this *DeviceController) Get() {
	// this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	// this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	// this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	this.Layout = "mobindex.html"

	var imei string
	this.Ctx.Input.Bind(&imei, "imei")
	result := models.GetVehicle(imei, nil)

	this.Data["device"] = result
	this.TplName = "subviews/devicedetails.html"

}
