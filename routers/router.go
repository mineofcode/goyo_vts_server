package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"

	"goyo.in/gpstracker/controllers"
	ctrl "goyo.in/gpstracker/webapp/controllers"
)

func init() {

	// namespaces
	var namespaces []string = []string{"goyoapi", "another"}

	//
	beego.NSNamespace("/",
		beego.NSRouter("/socket.io", &controllers.SocketController{}, "get:Teset"),
	)
	beego.AddNamespace(beego.NewNamespace("/"+namespaces[0],
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

		//tripapi/deleteGeoFence

		beego.NSNamespace("/tripapi/vehicle",
			beego.NSInclude(
				&controllers.VehicleController{},
			),
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

		// beego.NSNamespace("/tripapi/locations",
		// 	beego.NSInclude(
		// 		&controllers.VehicleController{},
		// 	),

	))

	//OverSpeed
	//tripapi/deleteGeoFence

	/*Web site*/
	//tripapi/deleteGeoFence
	beego.Router("/hello-world", &ctrl.MainController{})
	beego.Router("view/device:imei", &ctrl.DeviceController{})

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
}
