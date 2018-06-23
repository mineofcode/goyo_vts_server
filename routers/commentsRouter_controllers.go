package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:CreateGeoFenceController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:CreateGeoFenceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:DeletGeoFenceController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:DeletGeoFenceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:DeviceMasterController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:DeviceMasterController"],
		beego.ControllerComments{
			Method: "GetDevices",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:DeviceMasterController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:DeviceMasterController"],
		beego.ControllerComments{
			Method: "AddDevices",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:GetGeoFenceController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:GetGeoFenceController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:LoginController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:LoginController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:LoginController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:username/:passwd`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:LoginSessionController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:LoginSessionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:LoginSessionController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:LoginSessionController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:username/:passwd`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:LogoutController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:LogoutController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:LogoutController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:LogoutController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:username/:passwd`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:OverSpeedCountController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:OverSpeedCountController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:RegisterController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:RegisterController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:RegisterController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:RegisterController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:username/:passwd`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:ReportController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:ReportController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:SIMCardController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:SIMCardController"],
		beego.ControllerComments{
			Method: "GetSIM",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:SIMCardController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:SIMCardController"],
		beego.ControllerComments{
			Method: "CheckSIM",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:SIMCardController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:SIMCardController"],
		beego.ControllerComments{
			Method: "AddSIM",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:TripDataController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:TripDataController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:TripDataController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:TripDataController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/:username/:passwd`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:TripHistoryDataController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:TripHistoryDataController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:TripHistoryDataMakerController"] = append(beego.GlobalControllerRouter["goyo.in/gpstracker/controllers:TripHistoryDataMakerController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
