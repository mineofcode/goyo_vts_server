package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/models"
	"goyo.in/gpstracker/utils"
)

// Operations about login
type DeviceMasterController struct {
	beego.Controller
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router / [post]
func (o *DeviceMasterController) GetDevices() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	// var ob datamodel.DeviceMaster
	// json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	// result, err := models.CreateGeoFence(ob)
	// if err != nil {
	// 	o.Data["json"] = err.Error()
	// } else {
	// 	o.Data["json"] = utils.CreateWrap("200", result)
	// }
	o.ServeJSON()
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router / [post]
func (this *DeviceMasterController) CheckDevice() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	var imei = ""
	this.Ctx.Input.Bind(&imei, "imei")
	fmt.Println(imei)
	result := models.CheckDeviceActivate(imei)
	this.Data["json"] = utils.CreateWrap("200", result)
	this.ServeJSON()
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router / [post]
func (o *DeviceMasterController) AddDevices() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	var ob datamodel.DeviceMaster
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	result := models.CreateDeviceEntry(ob)
	o.Data["json"] = utils.CreateWrap("200", result)
	o.ServeJSON()
}
