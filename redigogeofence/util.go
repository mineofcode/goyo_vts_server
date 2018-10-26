package redigogeofence

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
	opts "goyo.in/gpstracker/const"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/models"
)

// type Req struct {
// 	Req interface{} `bson:"req" json:"req"`
// 	Res interface{} `bson:"res" json:"res"`
// }

func CallService(params datamodel.GeofenceDetect) {
	// rq := Req{}
	var hookin = strings.Index(params.Hook, ":") + 1

	var skey = params.Hook[hookin:len(params.Hook)]

	data, _ := models.GetGeoFenceSingle(bson.M{
		"imei":  params.Key,
		"fncnm": skey,
	})

	isCallurl := true //callable url assign
	// check time to send notification if time is available
	// fmt.Println(data.FenceTime)
	if data.FenceTime != "" { // if fence time is present then process time else anytime allow

		res := []datamodel.FenceTime{}                      // create varible
		err := json.Unmarshal([]byte(data.FenceTime), &res) //

		if err != nil {
			fmt.Println(err)
			return
		}

		for _, tm := range res { // loop for days loop

			loc, err := time.LoadLocation(opts.DefaultOpts.Config.TimeZone) // Get time zone from settings
			if err == nil {
				fmt.Println(err)
				start := time.Now().In(loc)                        //get curren time based on time zone
				weekd := strings.ToLower(start.Weekday().String()) // get weekday and make it lowr case "monday
				if strings.Contains(tm.Day, weekd[0:3]) {          //get first three character from week days "mon and chek its exists in strind mon,
					timess := strings.Split(tm.Time, ",") // split time from strinf 17:00,20:00
					frmtm := timess[0]
					fmt.Println(fmt.Sprintf("%04d-%02d-%02dT%v:00+05:30",
						start.Year(), int(start.Month()), int(start.Day()), frmtm))
					// first time
					t1, t1err := time.Parse(
						time.RFC3339,
						fmt.Sprintf("%04d-%02d-%02dT%v:00+05:30",
							start.Year(), int(start.Month()), int(start.Day()), frmtm)) //parse only start  time to date time
					if t1err != nil { //check error if error then skip this row
						continue
					}
					totm := timess[1]
					t2, t2err := time.Parse(
						time.RFC3339,
						fmt.Sprintf("%04d-%02d-%02dT%v:00+05:30",
							start.Year(), int(start.Month()), int(start.Day()), totm)) //parse only to  time to date time
					if t2err != nil { //check error if error then skip this row
						continue
					}
					// t1.Add(-15 * time.Minute)
					// t2.Add(15 * time.Minute)
					if start.After(t1) && start.Before(t2) { // current time is between two times then valid
						isCallurl = true // make api callable
					}
					break //if day found in string then skip next loop
				}
			}
		}
	}
	// rq.Req = data
	if !isCallurl {
		// InserLog(rq)
		// rq.Res = "Time Not Match"
		return
	}

	param := fmt.Sprintf("imei=%s&fncnm=%s&almtyp=%s&tm=%s&%s",
		data.Imei,
		data.FenceName,
		params.Detect,
		params.Time,
		data.Params,
	)

	//urls = "http://bulksms.mysmsmantra.com:8080/WebSMS/SMSAPI.jsp?username=CHETANAPUB&password=2054258156&sendername=chetna&mobileno=919819882904&message=vehiclereachxxx"
	t := &url.URL{Path: param}
	///urlenc, _ := UrlEncoded(urls)
	urlenc := t.String()

	urlenc = strings.Replace(urlenc, "./", "", -1)
	urls := fmt.Sprintf("%s?%s",
		data.CallBackURL,
		urlenc,
	)

	// fmt.Println(urls)
	req, err := http.NewRequest("GET", urls, nil)
	if err != nil {
		// rq.Res = err.Error()
		log.Fatal("NewRequest: ", err)
		// InserLog(rq)
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
		// rq.Res = err.Error()
		log.Fatal("Do: ", err)
		//	InserLog(rq)
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
		// rq.Res = err.Error()
		fmt.Println(err)
		//	InserLog(rq)
	} else {
		// rq.Res = body
		fmt.Println("geopoint", string(body))
	}

}

func InserLog(d interface{}) {
	// go models.CreateGeoFenceLog(d)
}

// UrlEncoded encodes a string like Javascript's encodeURIComponent()
func UrlEncoded(str string) (string, error) {
	u, err := url.Parse(str)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
