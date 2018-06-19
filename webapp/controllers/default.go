package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	this.Layout = "index.html"
	this.Data["Website"] = "My Website"
	this.Data["Email"] = "your.email.address@example.com"
	this.Data["EmailName"] = "Your Name"
	this.TplName = "test.tpl"

}
