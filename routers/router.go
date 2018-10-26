package routers

import (
	"fmt"

	"github.com/astaxie/beego"

	"goyo.in/gpstracker/controllers"
)

func init() {

	fmt.Println("Adding Router")
	// namespaces

	//

	ns := beego.NewNamespace("/goyoapi",
		beego.NSNamespace("/tripapi/user",
			beego.NSRouter("/login", &controllers.LoginController{}, "post:Login"),
			beego.NSRouter("/logout", &controllers.LogoutController{}, "post:Logout"),
			beego.NSRouter("/register", &controllers.RegisterController{}, "post:Register"),
			beego.NSRouter("/verifyotp", &controllers.RegisterController{}, "post:VerifyOtp"),
			beego.NSRouter("/resendotp", &controllers.RegisterController{}, "post:ResendOtp"),
			beego.NSRouter("/loginsession", &controllers.LoginSessionController{}, "post:LoginSession"),
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

		beego.NSNamespace("/tripapi/syncGeoFence",
			beego.NSInclude(
				&controllers.SyncGeoFenceController{},
			),
		),

		beego.NSNamespace("/tripapi/vehicle",
			beego.NSRouter("", &controllers.VehicleController{}, "post:Post"),
			beego.NSRouter("/getVehicleByUID", &controllers.VehicleController{}, "post:GetVehicleByUID"),
			beego.NSRouter("/getVehicleDetails", &controllers.VehicleController{}, "post:GetVehicleDetails"),
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
		beego.NSNamespace("/tripapi/mom",
			beego.NSRouter("/save", &controllers.MoMController{}, "post:Save"),
			beego.NSRouter("/get", &controllers.MoMController{}, "post:Get"),
		),
		beego.NSNamespace("/tripapi/fuel",
			beego.NSRouter("/add", &controllers.FuelController{}, "post:Add"),
			beego.NSRouter("/get", &controllers.FuelController{}, "post:Get"),
			beego.NSRouter("/get/edit", &controllers.FuelController{}, "post:GetEdit"),
		),
		beego.NSNamespace("/tripapi/locations",
			beego.NSRouter("/getnearbyvehicles", &controllers.LocationsController{}, "post:GetNearByVehicles"),
		),
	)

	beego.AddNamespace(ns)

	//OverSpeed
	//tripapi/deleteGeoFence
	//
	/*Web site*/
	//tripapi/deleteGeoFence

}
