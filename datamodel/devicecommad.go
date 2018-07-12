package datamodel

import (
	"time"
)

//DeviceCommands model
type DeviceCommands struct {
	Uniqid       uint      `bson:"uniqid" json:"uniqid"`
	Cmd          string    `bson:"cmd" json:"cmd"`
	Imei         string    `bson:"imei" json:"imei"`
	Extra        string    `bson:"extra" json:"extra"`
	UID          string    `bson:"uid" json:"uid"`
	Ucode        string    `bson:"ucode" json:"ucode"`
	Platform     string    `bson:"platform" json:"platform"`
	Src          string    `bson:"src" json:"src"`
	IP           string    `bson:"ip" json:"ip"`
	Time         time.Time `bson:"time" json:"time"`
	Resultraw    string    `bson:"resultraw" json:"resultraw"`
	Stringresult string    `bson:"result" json:"result"`
	DeviceID     string    `bson:"devid" json:"devid"`
}
