package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

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
