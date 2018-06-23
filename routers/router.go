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
		beego.NSNamespace("/tripapi/logout",
			beego.NSInclude(
				&controllers.LogoutController{},
			),
		),

		beego.NSNamespace("/tripapi/loginsession",
			beego.NSInclude(
				&controllers.LoginSessionController{},
			),
		),
		beego.NSNamespace("/tripapi/register",
			beego.NSInclude(
				&controllers.RegisterController{},
			),
		),
		beego.NSNamespace("/tripapi/getvahicleupdates",
			beego.NSInclude(
				&controllers.TripDataController{},
			),
		),
		beego.NSNamespace("/tripapi/getHistory",
			beego.NSInclude(
				&controllers.TripHistoryDataController{},
			),
		),
		beego.NSNamespace("/tripapi/createHistory",
			beego.NSInclude(
				&controllers.TripHistoryDataMakerController{},
			),
		),

		//tripapi/getReports

		beego.NSNamespace("/tripapi/report",
			beego.NSInclude(
				&controllers.ReportController{},
			),
		),

		//tripapi/createGeoFence

		beego.NSNamespace("/tripapi/createGeoFence",
			beego.NSInclude(
				&controllers.CreateGeoFenceController{},
			),
		),
		//tripapi/getGeoFence

		beego.NSNamespace("/tripapi/getGeoFence",
			beego.NSInclude(
				&controllers.GetGeoFenceController{},
			),
		),

		//tripapi/deleteGeoFence

		beego.NSNamespace("/tripapi/deleteGeoFence",
			beego.NSInclude(
				&controllers.DeletGeoFenceController{},
			),
		),

		//tripapi/deleteGeoFence

		beego.NSNamespace("/tripapi/vehicle",
			beego.NSInclude(
				&controllers.VehicleController{},
			),
			beego.NSRouter("/getVehicleByUID", &controllers.VehicleController{}, "post:GetVehicleByUID"),
		),
		

		beego.NSNamespace("/tripapi/getOverSpeedCount",
			beego.NSInclude(
				&controllers.OverSpeedCountController{},
			),
		),
		beego.NSNamespace("/tripapi/device",
			beego.NSRouter("/check", &controllers.DeviceMasterController{}, "get,post:CheckDevice"),
			beego.NSRouter("/:id", &controllers.DeviceMasterController{}, "get:GetDevices"),
			beego.NSRouter("/add", &controllers.DeviceMasterController{}, "post:AddDevices"),
			beego.NSRouter("/activate", &controllers.VehicleController{}, "post:ActivateVehicle"),
		),
		beego.NSNamespace("/tripapi/sim",
			beego.NSRouter("/add", &controllers.SIMCardController{}, "post:AddSIM"),
		),
	))

	//OverSpeed
	//tripapi/deleteGeoFence

	/*Web site*/
	//tripapi/deleteGeoFence
	beego.Router("/hello-world", &ctrl.MainController{})
	beego.Router("view/device:imei", &ctrl.DeviceController{})
}
