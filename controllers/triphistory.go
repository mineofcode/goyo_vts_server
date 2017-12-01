package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"goyo.in/gpstracker/dataprocess"
	"goyo.in/gpstracker/models"
	"goyo.in/gpstracker/utils"
)

// Operations about login
type TripHistoryDataController struct {
	beego.Controller
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router / [post]
func (o *TripHistoryDataController) Post() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	var ob models.ParamsTripHistorydata
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	result, err := dataprocess.GetHistoryData(ob, nil)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = utils.CreateWrap("200", result)
	}
	o.ServeJSON()
}
