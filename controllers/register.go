package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/models"
)

// Operations about login
type RegisterController struct {
	beego.Controller
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router / [post]
func (o *RegisterController) Register() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	var ob datamodel.User
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	result, err := models.RegisterUser(ob)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = result
	}
	o.ServeJSON()
}

func (o *RegisterController) VerifyOtp() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	var ob datamodel.User
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	result, err := models.VerifyOtp(ob)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = result
	}
	o.ServeJSON()
}

func (o *RegisterController) ResendOtp() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	var ob datamodel.User
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	result := models.ResendOtp(ob)
	o.Data["json"] = result
	o.ServeJSON()
}
