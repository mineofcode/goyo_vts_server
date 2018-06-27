package datamodel

import (
	"time"
)

type (
	// MasterOfMaster represents the structure of our resource
	MasterOfMaster struct {
		ID       int         `bson:"_id" json:"_id"`
		Key      string      `bson:"key" json:"key"`
		Value    string      `bson:"value" json:"value"`
		ValueD   interface{} `bson:"valued" json:"valued"`
		Group    string      `bson:"group" json:"group"`
		IsActive bool        `bson:"active" json:"active"`
		Cron     time.Time   `bson:"cron" json:"cron"`
		Crby     string      `bson:"crby" json:"crby"`
		Upon     time.Time   `bson:"upon" json:"upon"`
		Upby     string      `bson:"upby" json:"upby"`
	}
)
