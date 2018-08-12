package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

type SocketController struct {
	beego.Controller
}

func (this *SocketController) Teset() {
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	this.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	fmt.Println("test Get method")
}
