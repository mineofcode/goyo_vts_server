package redigogeofence

import (
	"fmt"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/models"
)

//var m map[string][]Sel_params

// func CallService(points []Sel_params) {
// 	m = make(map[string][]Sel_params)
// 	for _, var1 := range points {
// 		m[var1.url+""+var1.param] = append(m[var1.url+""+var1.param], var1)

// 	}
// 	fmt.Println(len(m))

// 	for _, obj := range m {
// 		var url = ""
// 		//var param = ""
// 		//for _, _ob := range obj {
// 		data := obj[0]
// 		url = data.url + "?" + data.param
// 		//param = _ob.param
// 		//fmt.Println(url)
// 		//}
// 		//123456789012345 schoolt circle inout [] http://callme21232131 test
// 		resp, err := http.Get(url)

// 		if err != nil {
// 			//fmt.Println(err)
// 			continue
// 		}
// 		fmt.Println(url)
// 		// Callers should close resp.Body
// 		// when done reading from it
// 		// Defer the closing of the body
// 		defer resp.Body.Close()

// 	}
// }

func CallService(params datamodel.GeofenceDetect) {

	var hook = strings.Split(params.Hook, ":")

	var skey = hook[1]

	result, _ := models.GetGeoFence(bson.M{
		"imei":  params.Key,
		"fncnm": skey,
	})
	if len(result) > 0 {
		data := result[0]
		url := fmt.Sprintf("%s?imei=%s&fncnm=%s&almtyp=%s&tm=%s&%s",
			data.CallBackURL,
			data.Imei,
			data.FenceName,
			params.Detect,
			params.Time,
			data.Params,
		)
		// //var param = ""
		// //for _, _ob := range obj {
		// data := obj[0]
		// url = data.url + "?" + data.param
		//param = _ob.param
		fmt.Println(url)
		//}
		//123456789012345 schoolt circle inout [] http://callme21232131 test
		resp, err := http.Get(url)

		if err != nil {
			fmt.Println(err)
		}
		// Callers should close resp.Body
		// when done reading from it
		// Defer the closing of the body
		defer resp.Body.Close()

	}

}
