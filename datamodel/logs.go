package datamodel

type Logs struct {
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
