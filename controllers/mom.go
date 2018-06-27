package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"goyo.in/gpstracker/models"
)

// Operations about login
type MoMController struct {
	beego.Controller
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router / [post]
func (o *MoMController) Save() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	var ob map[string]interface{}
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	result := models.SaveMoM(ob)
	o.Data["json"] = result
	o.ServeJSON()
}

func (o *MoMController) Get() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	var ob map[string]interface{}
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	result := models.GetMoM(ob)
	o.Data["json"] = result
	o.ServeJSON()

}
