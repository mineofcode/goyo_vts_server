package datamodel

import (
	"time"
)

//DeviceCommands model
type SIMMaster struct {
	MobNo       string    `bson:"mobno" json:"mobno" validate:"required,numeric,len=10"`
	SimId       string    `bson:"simid" json:"simid"`
	SimOperator string    `bson:"oper" 	json:"oper"  validate:"required"`
	CreateOn    time.Time `bson:"cron" 	json:"-"`
	UpdateOn    time.Time `bson:"upon" 	json:"-"`
	CreatedBy   string    `bson:"crby" 	json:"-"  validate:"required"`
	UpdatedBy   string    `bson:"upby"  json:"-"`
}
