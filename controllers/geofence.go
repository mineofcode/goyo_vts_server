package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/models"
	"goyo.in/gpstracker/utils"
)

// Operations about login
type CreateGeoFenceController struct {
	beego.Controller
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router / [post]
func (o *CreateGeoFenceController) Post() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	var ob datamodel.GeoFenceModelAr
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	result, err := models.CreateGeoFence(ob)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = utils.CreateWrap("200", result)
	}
	o.ServeJSON()
}

// Operations about login
type GetGeoFenceController struct {
	beego.Controller
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router / [post]
func (o *GetGeoFenceController) Post() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	//var ob datamodel.GetGeoFenceParams
	var ob interface{}
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	result, err := models.GetGeoFence(ob)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = utils.CreateWrap("200", result)
	}
	o.ServeJSON()
}

// Operations about login
type DeletGeoFenceController struct {
	beego.Controller
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router / [post]
func (o *DeletGeoFenceController) Post() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	var ob datamodel.GetGeoFenceParams
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	result, err := models.DeleteGeoFence(ob)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = utils.CreateWrap("200", result)
	}
	o.ServeJSON()
}
