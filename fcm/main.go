package fcm

import (
	"fmt"
	"strconv"

	fcm "github.com/NaySoftware/go-fcm"
	opts "goyo.in/gpstracker/const"
	"goyo.in/gpstracker/models"
)

var serverKey = opts.DefaultOpts.Fcm.ServerKey
var fcmclnt *fcm.FcmClient

//SendSpeedAlertTotopic send fcm message to topic
func SendSpeedAlertTotopic(imei string, speed int) {
	topic := fmt.Sprintf("/topics/speed_%s", imei)
	//
	vhdet := models.GetVehiclesData(imei)

	var NP fcm.NotificationPayload
	NP.Title = fmt.Sprintf("Speed Alert : %s", vhdet.VhNm)
	NP.Body = fmt.Sprintf("Allowd speed %d Current speed %d", vhdet.AllowSpd, speed)
	NP.Sound = "default"

	data := map[string]string{
		"vhid":  imei,
		"speed": strconv.Itoa(speed),
		"topic": "speed",
		"title": NP.Title,
		"body":  NP.Body,
	}

	if fcmclnt == nil {
		fcmclnt = fcm.NewFcmClient(serverKey)
	}
	fcmclnt.SetTimeToLive(60000)
	fcmclnt.NewFcmTopicMsg(topic, data)
	fcmclnt.SetNotificationPayload(&NP)

	status, err := fcmclnt.Send()
	if err == nil {
		status.PrintResults()
	} else {
		fmt.Println(err)
	}
	fmt.Printf(topic)
}
