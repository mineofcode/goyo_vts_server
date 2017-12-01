package routers

import (
	"github.com/astaxie/beego"
	"goyo.in/gpstracker/controllers"
)

func init() {

	// namespaces
	var namespaces []string = []string{"goyoapi", "another"}

	// controllers
	//var ctrllers []string = []string{"login", "other"}

	// restfulRouter := beego.NewNamespace("/"+namespaces[0],
	// 	beego.NSNamespace("/"+ctrllers[0],
	// 		beego.NSInclude(
	// 			&controllers.LoginController{},
	// 		),
	// 	),
	// )
	// beego.AddNamespace(restfulRouter)

	//tripapi/getvahicleupdates
	beego.AddNamespace(beego.NewNamespace("/"+namespaces[0],
		beego.NSNamespace("/tripapi/getvahicleupdates",
			beego.NSInclude(
				&controllers.TripDataController{},
			),
		),
	))

	//tripapi/getHistory
	beego.AddNamespace(beego.NewNamespace("/"+namespaces[0],
		beego.NSNamespace("/tripapi/getHistory",
			beego.NSInclude(
				&controllers.TripHistoryDataController{},
			),
		),
	))

	//tripapi/getHistory
	beego.AddNamespace(beego.NewNamespace("/"+namespaces[0],
		beego.NSNamespace("/tripapi/createHistory",
			beego.NSInclude(
				&controllers.TripHistoryDataMakerController{},
			),
		),
	))
}
