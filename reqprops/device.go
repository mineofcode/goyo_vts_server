package reqprops

type DeviceActivationProp struct {
	IMEI     string `json:"imei"`
	UID      string `json:uid`
	Vhname   string `json:vhname`
	RegNo    string `json:regno`
	VhType   string `json:vhtyp`
	AlwSpeed string `json:alwspeed`
	D1Func   string `json:d1`
}
