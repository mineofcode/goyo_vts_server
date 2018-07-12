package datamodel

//DeviceCommands model
import "time"

type FuelEntry struct {
	Autoid        int       `bson:"autoid" json:"autoid" validate:"required"`
	UID           int       `bson:"uid" json:"uid" validate:"required"`
	Imei          string    `bson:"imei" json:"imei" validate:"required,numeric,len=15"`
	VtsID         int       `bson:"vtsid" json:"vtsid" validate:"required,numeric"`
	VhName        string    `bson:"vhnm" json:"vhnm"`
	VhNo          string    `bson:"regno" json:"regno"`
	UOM           string    `bson:"uom" json:"uom"`
	ODO           int       `bson:"odo" json:"odo"`
	FuelType      string    `bson:"fueltyp" json:"fueltyp"`
	Date          time.Time `bson:"date" json:"date"`
	PricePerLiter float64   `bson:"prcperltr" json:"prcperltr"`
	Liter         float64   `bson:"liter" json:"liter"`
	LatLon        []float64 `bson:"latlon" json:"latlon"`
	Amount        float64   `bson:"amount" json:"amount"`
	Attachement   string    `bson:"attch" json:"attch"`
	Remark        string    `bson:"remark" json:"remark"`
	CreateOn      time.Time `bson:"cron" json:"-"`
	UpdateOn      time.Time `bson:"upon" json:"-"`
	CreatedBy     string    `bson:"crby" json:"crby"`
	UpdatedBy     string    `bson:"upby" json:"-"`
}

//
type SelectAllFuelEntry struct {
	Autoid   int     `bson:"autoid" json:"autoid"`
	VtsID    int     `bson:"vtsid" json:"vtsid"`
	VhName   string  `bson:"vhnm" json:"vhnm"`
	VhNo     string  `bson:"regno" json:"regno"`
	UOM      string  `bson:"uom" json:"uom"`
	FuelType string  `bson:"fueltyp" json:"fueltyp"`
	Date     string  `bson:"date" json:"date"`
	Liter    float64 `bson:"liter" json:"liter"`
	Amount   float64 `bson:"amount" json:"amount"`
}

//
type SelectEditFuelEntry struct {
	Autoid        int       `bson:"autoid" json:"autoid"`
	UID           int       `bson:"uid" json:"uid"`
	ODO           int       `bson:"odo" json:"odo"`
	VhNo          string    `bson:"regno" json:"regno"`
	Imei          string    `bson:"imei" json:"imei"`
	VtsID         int       `bson:"vtsid" json:"vtsid"`
	VhName        string    `bson:"vhnm" json:"vhnm"`
	UOM           string    `bson:"uom" json:"uom"`
	FuelType      string    `bson:"fueltyp" json:"fueltyp"`
	Date          string    `bson:"date" json:"date"`
	PricePerLiter float64   `bson:"prcperltr" json:"prcperltr"`
	Liter         float64   `bson:"liter" json:"liter"`
	LatLon        []float64 `bson:"latlon" json:"latlon"`
	Amount        float64   `bson:"amount" json:"amount"`
	Attachement   string    `bson:"attch" json:"attch"`
	Remark        string    `bson:"remark" json:"remark"`
}
