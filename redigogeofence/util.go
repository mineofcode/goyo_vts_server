package redigogeofence

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

		params := fmt.Sprintf("imei=%s&fncnm=%s&almtyp=%s&tm=%s&%s",
			data.Imei,
			data.FenceName,
			params.Detect,
			params.Time,
			data.Params,
		)

		//urls = "http://bulksms.mysmsmantra.com:8080/WebSMS/SMSAPI.jsp?username=CHETANAPUB&password=2054258156&sendername=chetna&mobileno=919819882904&message=vehiclereachxxx"
		t := &url.URL{Path: params}
		///urlenc, _ := UrlEncoded(urls)
		urlenc := t.String()

		urlenc = strings.Replace(urlenc, "./", "", -1)
		urls := fmt.Sprintf("%s?%s",
			data.CallBackURL,
			urlenc,
		)

		fmt.Println(urls)
		req, err := http.NewRequest("GET", urls, nil)
		if err != nil {
			log.Fatal("NewRequest: ", err)
			return
		}

		// For control over HTTP client headers,
		// redirect policy, and other settings,
		// create a Client
		// A Client is an HTTP client
		client := &http.Client{}

		// Send the request via a client
		// Do sends an HTTP request and
		// returns an HTTP response
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Do: ", err)
			return
		}

		// Callers should close resp.Body
		// when done reading from it
		// Defer the closing of the body
		defer resp.Body.Close()

		// //var param = ""
		// //for _, _ob := range obj {
		// data := obj[0]
		// url = data.url + "?" + data.param
		//param = _ob.param

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(body)
		}

	}

}

// UrlEncoded encodes a string like Javascript's encodeURIComponent()
func UrlEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
