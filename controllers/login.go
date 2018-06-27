package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"goyo.in/gpstracker/models"
)

// Operations about login
type LoginController struct {
	beego.Controller
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router / [post]
func (o *LoginController) Login() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	var ob models.Login
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	result, err := models.VerifyLogin(ob)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = result
	}
	o.ServeJSON()
}

// Operations about login
type LogoutController struct {
	beego.Controller
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router / [post]
func (o *LogoutController) Logout() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	var ob models.Session
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	result, err := models.Logout(ob)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = result
	}
	o.ServeJSON()
}

// Operations about login
type LoginSessionController struct {
	beego.Controller
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router / [post]
func (o *LoginSessionController) LoginSession() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	var ob models.Session
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	result, err := models.VerifySession(ob)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = result
	}
	o.ServeJSON()
}
