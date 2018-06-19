package routers

import (
	"github.com/astaxie/beego"
	"goyo.in/gpstracker/controllers"
	ctrl "goyo.in/gpstracker/webapp/controllers"
)

func init() {

	// namespaces
	var namespaces []string = []string{"goyoapi", "another"}

	//
	beego.AddNamespace(beego.NewNamespace("/"+namespaces[0],
		beego.NSNamespace("/tripapi/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
	))

	//
	beego.AddNamespace(beego.NewNamespace("/"+namespaces[0],
		beego.NSNamespace("/tripapi/logout",
			beego.NSInclude(
				&controllers.LogoutController{},
			),
		),
	))
	//
	beego.AddNamespace(beego.NewNamespace("/"+namespaces[0],
		beego.NSNamespace("/tripapi/loginsession",
			beego.NSInclude(
				&controllers.LoginSessionController{},
			),
		),
	))
	// controllers
	//var ctrllers []string = []string{"login", "other"}

	beego.AddNamespace(beego.NewNamespace("/"+namespaces[0],
		beego.NSNamespace("/tripapi/register",
			beego.NSInclude(
				&controllers.RegisterController{},
			),
		),
	))
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

	//tripapi/getReports
	beego.AddNamespace(beego.NewNamespace("/"+namespaces[0],
		beego.NSNamespace("/tripapi/report",
			beego.NSInclude(
				&controllers.ReportController{},
			),
		),
	))

	//tripapi/createGeoFence
	beego.AddNamespace(beego.NewNamespace("/"+namespaces[0],
		beego.NSNamespace("/tripapi/createGeoFence",
			beego.NSInclude(
				&controllers.CreateGeoFenceController{},
			),
		),
	))

	//tripapi/getGeoFence
	beego.AddNamespace(beego.NewNamespace("/"+namespaces[0],
		beego.NSNamespace("/tripapi/getGeoFence",
			beego.NSInclude(
				&controllers.GetGeoFenceController{},
			),
		),
	))

	//tripapi/deleteGeoFence
	beego.AddNamespace(beego.NewNamespace("/"+namespaces[0],
		beego.NSNamespace("/tripapi/deleteGeoFence",
			beego.NSInclude(
				&controllers.DeletGeoFenceController{},
			),
		),
	))

	//tripapi/deleteGeoFence
	beego.AddNamespace(beego.NewNamespace("/"+namespaces[0],
		beego.NSNamespace("/tripapi/vehicle",
			beego.NSInclude(
				&controllers.VehicleController{},
			),
		),
	))

	//OverSpeed
	//tripapi/deleteGeoFence
	beego.AddNamespace(beego.NewNamespace("/"+namespaces[0],
		beego.NSNamespace("/tripapi/getOverSpeedCount",
			beego.NSInclude(
				&controllers.OverSpeedCountController{},
			),
		),
	))

	/*Web site*/
	//tripapi/deleteGeoFence
	beego.Router("/hello-world", &ctrl.MainController{})
	beego.Router("view/device:imei", &ctrl.DeviceController{})
}
