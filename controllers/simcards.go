package controllers

import (
	"encoding/json"
	"fmt"

	"goyo.in/gpstracker/datamodel"

	"github.com/astaxie/beego"
	"goyo.in/gpstracker/models"
	"goyo.in/gpstracker/utils"
)

// Operations about login
type SIMCardController struct {
	beego.Controller
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router / [post]
func (o *SIMCardController) GetSIM() {
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
func (this *SIMCardController) CheckSIM() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	var imei = this.Ctx.Input.Param(":imei")
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
func (o *SIMCardController) AddSIM() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	var ob datamodel.SIMMaster
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	result := models.CreateSIMEntry(ob)
	o.Data["json"] = utils.CreateWrap("200", result)
	o.ServeJSON()
}
