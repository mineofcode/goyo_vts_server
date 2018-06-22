package datamodel

import (
	"time"
)

//DeviceCommands model
type DeviceMaster struct {
	Imei       string    `bson:"imei" 	json:"imei" validate:"required,numeric,len=10"`
	SimNo      string    `bson:"sim" 	json:"sim"`
	DeviceType string    `bson:"devtyp" json:"devtyp"`
	Date       time.Time `bson:"date" 	json:"date"`
	CreateOn   time.Time `bson:"cron" 	json:"-"`
	UpdateOn   time.Time `bson:"upon" 	json:"-"`
	CreatedBy  string    `bson:"crby" 	json:"-" validate:"required"`
	UpdatedBy  string    `bson:"upby"   json:"-"`
}
