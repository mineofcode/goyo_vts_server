package fcm

import (
	"fmt"
	"strconv"
	"time"

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

	duration := time.Since(vhdet.LstSpdtm)

	if duration.Minutes() <= 2 && (vhdet.AllowSpd-speed) < 20 {
		return
	}

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
	sendNotification(topic, data, NP, 60000)
}

//SendSpeedAlertTotopic send fcm message to topic
func SendEventAlertTotopic(id uint32, replytext string) {
	topic := fmt.Sprintf("/topics/evt_%d", id)
	// //
	// vhdet := models.GetVehiclesData(imei)

	var NP fcm.NotificationPayload
	NP.Title = fmt.Sprintf("Command Reply")
	NP.Body = fmt.Sprintf(replytext)
	NP.Sound = "default"

	data := map[string]string{
		"topic": "evt",
		"title": NP.Title,
		"body":  NP.Body,
	}
	sendNotification(topic, data, NP, 5000)
}

// final send notificaiton
func sendNotification(topic string, data map[string]string, NP fcm.NotificationPayload, timetolive int) {

	if fcmclnt == nil {
		fcmclnt = fcm.NewFcmClient(serverKey)
	}
	fcmclnt.SetTimeToLive(timetolive)
	fcmclnt.NewFcmTopicMsg(topic, data)
	fcmclnt.SetNotificationPayload(&NP)

	_, err := fcmclnt.Send()
	if err == nil {
		//status.PrintResults()
	} else {
		fmt.Println(err)
	}
	//fmt.Printf(topic)
}
