package datamodel

import (
	"gopkg.in/mgo.v2/bson"
)

type GeoFenceModelAr struct {
	Data []GeoFenceModel `bson:"data" json:"data"` // imei number of device

}

type GeoFenceModel struct {
	Imei        string    `bson:"imei" json:"imei"`     // imei number of device
	FenceName   string    `bson:"fncnm" json:"fncnm"`   // fence unique name
	FenceType   string    `bson:"fnctyp" json:"fnctyp"` // fence type  circle|polygon|arc etc...
	AlamTyp     string    `bson:"almtyp" json:"almtyp"` // alarm type	 in|out|inout
	Active      bool      `bson:"isact" json:"isact"`   // active true|false
	Points      []float64 `bson:"points" json:"points"` // points based on fence type
	Radius      int       `bson:"radius" json:"radius"` // radious for circle type fence
	CallBackURL string    `bson:"url" json:"url"`       // call back url if alarm occured
	Params      string    `bson:"param" json:"param"`   // callback url params
	TimeType    string    `bson:"tmtyp" json:"tmtyp"`   // TimeType  daily|datetime|date|time
	FenceTime   string    `bson:"time" json:"time"`     // Time to Fence

}

type Meta struct {
	Time []FenceTime `json:"time"` // TimeType  daily|datetime|date|time
}

type FenceTime struct {
	Batchid int    `json:"batchid"` // TimeType  daily|datetime|date|time
	Day     string `json:"day"`     // Time to Fence
	Time    string `json:"time"`
}

type GetGeoFenceParams struct {
	Imei      string `bson:"imei" json:"imei"`   // imei number of device
	FenceName string `bson:"fncnm" json:"fncnm"` // fence unique name
}

type GeoFenceResponse struct {
	Imei      string `json:"imei"`  // imei number of device
	FenceName string `json:"fncnm"` // fence unique name
	Status    bool   `json:"status"`
	Msg       string `json:"msg"`
}

type GeoFenceResponseWrap struct {
	Message  string             `json:"msg"`  // imei number of device
	GeoFResp []GeoFenceResponse `json:"resp"` // fence unique name
}

type GeofenceDetect struct {
	Key    string `json:"key"`    // imei number of device
	Hook   string `json:"hook"`   // Hook
	Detect string `json:"detect"` // detect
	Time   string `json:"time"`
	Meta   bson.M `json:"meta"`
}
