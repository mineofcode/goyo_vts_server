package protocalHandler

import (
	"time"
)

type (
	// Vehicles represents the structure of our resource
	HertBt struct {
		Sertm  time.Time `bson:"sertm" json:"sertm"`
		Appvr  string    `bson:"appvr" json:"appvr"`
		Imei   string    `bson:"imei" json:"imei"`
		Btr    int       `bson:"btr" json:"btr"`
		Btrst  string    `bson:"btrst" json:"btrst"`
		Actvt  string    `bson:"actvt" json:"actvt"`
		Acttm  time.Time `bson:"acttm" json:"acttm"`
		Vhid   string    `bson:"vhid" json:"vhid"`
		Oe     int       `bson:"oe" json:"oe"`
		Gp     int       `bson:"gp" json:"gp"`
		Alm    string    `bson:"alm" json:"alm"`
		Chrg   int       `bson:"chrg" json:"chrg"`
		Acc    int       `bson:"acc" json:"acc"`
		Df     int       `bson:"df" json:"df"`
		Gsmsig int       `bson:"gsmsig" json:"gsmsig"`
		Flag   string    `bson:"flag" json:"flag"`
		Speed  int       `bson:"speed" json:"speed"`
	}
)
