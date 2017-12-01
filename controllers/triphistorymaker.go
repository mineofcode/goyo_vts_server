package controllers

import (
	"encoding/json"

	"goyo.in/gpstracker/job"

	"github.com/astaxie/beego"
	"goyo.in/gpstracker/utils"
)

// Operations about login
type TripHistoryDataMakerController struct {
	beego.Controller
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router / [post]
func (o *TripHistoryDataMakerController) Post() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	var ob interface{}
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	job.DailyData(ob)
	o.Data["json"] = utils.CreateWrap("200", "Updated")
	o.ServeJSON()
}
