package controllers

import (
	"encoding/json"
	"fmt"
	"reflect"

	"goyo.in/gpstracker/protocal"

	"github.com/astaxie/beego"
	"goyo.in/gpstracker/models"
	"goyo.in/gpstracker/utils"
)

// Operations about login
type VehicleController struct {
	beego.Controller
}

type vhmod struct {
	Vhid     string `json:"vhid"`
	Alwspeed int    `json:"alwspeed"`
}

// @Title Create
// @Description login
// @Param	body  body  models.Login  true  "The object content"
// @Success 200 {login} models.Login
// @Failure 403 body is empty
// @router / [post]
func (o *VehicleController) Post() {
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

	var ob interface{}
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	//req := vhmod{}

	//var err error

	var response vhmod
	// response.Count = 123
	_ = json.Unmarshal(o.Ctx.Input.RequestBody, &response)

	fmt.Println(response)

	param := reflect.ValueOf(ob)

	fmt.Println(param.MapIndex(reflect.ValueOf("alwspeed")).Kind())

	//alwspeed := param.MapIndex(reflect.ValueOf("alwspeed")).Interface()

	result, ip := models.UpdateVehicleData(ob, response.Vhid)

	if ip != "" {
		if response.Alwspeed != 0 {

			protocalHandler.UpdateAllowSpeed(response.Alwspeed, ip)
		}

	}

	o.Data["json"] = utils.CreateWrap("200", result)

	o.ServeJSON()
}
