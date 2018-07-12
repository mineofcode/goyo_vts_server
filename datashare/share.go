package datashare

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"goyo.in/gpstracker/shared"
)

//main function

type sgpsdata struct {
	imei        string
	lat         string
	lon         string
	speed       string
	acc         int
	timestamp   string
	timestamp_o time.Time
	ac          int
}

var gpsdata = sgpsdata{}

//SendToClient function
func SendToClient(vhclient shared.ClientsMod) {
	if vhclient.PClients == nil && len(vhclient.PClients) == 0 {
		return
	}

	if vhclient.Loc == nil {
		return
	}

	loc := vhclient.Loc
	gpsdata = sgpsdata{
		lat:         FloatToString(loc[1]),
		lon:         FloatToString(loc[0]),
		timestamp:   IntToString(vhclient.Lstm.Unix()),
		timestamp_o: vhclient.Lstm,
		acc:         vhclient.Acc,
		ac:          vhclient.AC,
		imei:        vhclient.Imei,
		speed:       strconv.Itoa(vhclient.Speed),
	}

	for d := range vhclient.PClients {
		var key = vhclient.PClients[d]
		go funcs[key](gpsdata)
	}

}

type fn func(sgpsdata)

var funcs map[string]fn

func Init() {
	funcs = map[string]fn{
		"rb": sendToRedBus,
		"tg": sendToTrackingo,
	}
}

///////////////////////////////////////////////RED BUS//////////////////////////////////////////////////
var (
	RDBusURL = "http://track1.yourbus.in/processGPSV2.php"
	RDKey    = "BhUm1tr0V8LsGhp0s"
)

func sendToRedBus(vh sgpsdata) {
	//http://track1.yourbus.in/processGPSV2.php?acc_key=BhUm1tr0V8LsGhp0s&gps_id=123456789012345&
	//llt1=18.6643966666667,73.7663966666667,1408366983,0,0,0,258.89,Pune District Maharashtra
	var acc = 0
	if vh.acc == 1 {
		acc = 0
	} else {
		acc = 1
	}

	var buffer bytes.Buffer
	buffer.WriteString(vh.lat) //lat
	buffer.WriteString(",")
	buffer.WriteString(vh.lon) //lon
	buffer.WriteString(",")
	buffer.WriteString(vh.timestamp) //timestamp unix
	buffer.WriteString(",")
	buffer.WriteString(vh.speed) //speed
	buffer.WriteString(",")
	buffer.WriteString(strconv.Itoa(acc)) //Ignition_status acc
	buffer.WriteString(",")
	buffer.WriteString(strconv.Itoa(vh.ac)) //ac staus
	buffer.WriteString(",")
	buffer.WriteString("0") //Orientation
	buffer.WriteString(",")
	buffer.WriteString("") //Address

	params := url.Values{
		"acc_key": {RDKey},
		"gps_id":  {vh.imei},
		"llt1":    {buffer.String()},
	}
	furl := fmt.Sprintf("%s?%s", RDBusURL,
		params.Encode())
	timeout := time.Duration(10 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	req, _ := http.NewRequest("GET", furl, nil)

	resp, _ := client.Do(req)

	if resp.StatusCode == 200 {
		fmt.Println(vh.imei, "rb - 0k")
	} else {
		fmt.Println(vh.imei, resp)
	}

}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func IntToString(input_num int64) string {
	// to convert a float number to a string
	return strconv.FormatInt(input_num, 10)
}

///////////////////////////////////////////////TrackinGo//////////////////////////////////////////////////
var (
	TgAccountKey = "4JENHJHU"
	TgVenderKey  = "DDOI7C5Z"
	TgURL        = "http://feed.trackingo.in/api/gf/feed/" + TgAccountKey + "/" + TgVenderKey
)

func sendToTrackingo(vh sgpsdata) {
	//http://track1.yourbus.in/processGPSV2.php?acc_key=BhUm1tr0V8LsGhp0s&gps_id=123456789012345&
	//llt1=18.6643966666667,73.7663966666667,1408366983,0,0,0,258.89,Pune District Maharashtra
	var acc = "IGN_OFF"
	if vh.acc == 1 {
		acc = "IGN_ON"
	} else {
		acc = "IGN_OFF"
	}

	formatedTime := vh.timestamp_o.Format("2006-01-02 15:04:05")
	fmt.Println(formatedTime)
	params := url.Values{
		"lat":        {vh.lat},
		"lng":        {vh.lon},
		"speed":      {vh.speed},
		"sn":         {vh.imei},
		"state":      {acc},
		"event_type": {"PERIODIC"},
		"datetime":   {formatedTime},
		"odometer":   {"0.0"},
	}

	timeout := time.Duration(10 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}
	req, _ := http.NewRequest("POST", TgURL, strings.NewReader(params.Encode()))

	resp, _ := client.Do(req)

	if resp.StatusCode == 200 {
		fmt.Println(vh.imei, "tg - 0k")
	} else {
		fmt.Println(vh.imei, resp)
	}
}
