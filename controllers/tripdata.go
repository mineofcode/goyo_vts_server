package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"goyo.in/gpstracker/models"
	"goyo.in/gpstracker/utils"
)

// Operations about login
type TripDataController struct {
	beego.Controller
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router /:username/:passwd [get]
func (o *TripDataController) Get() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	//var ob models.ParamsTripdata
	// ob.Vhids = o.Ctx.Input.Param(":vhids")

	// result, err := models.GetLastStatus(ob)
	// if err != nil {
	// 	o.Data["json"] = err.Error()
	// } else {
	// 	o.Data["json"] = result
	// }
	// o.ServeJSON()
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router / [post]
func (o *TripDataController) Post() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	var ob models.ParamsTripdata
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	result, err := models.GetLastStatus(ob)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = utils.CreateWrap("200", result)
	}
	o.ServeJSON()
}

func (o *TripDataController) Options() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	//var ob models.ParamsTripdata
	// ob.Vhids = o.Ctx.Input.Param(":vhids")

	// result, err := models.GetLastStatus(ob)
	// if err != nil {
	// 	o.Data["json"] = err.Error()
	// } else {
	// 	o.Data["json"] = result
	// }
	// o.ServeJSON()
}
