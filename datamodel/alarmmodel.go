package datamodel

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Vehicles represents the structure of our resource
	Events struct {
		ID    bson.ObjectId `bson:"_id,omitempty" json:"-"`
		Sertm time.Time     `bson:"sertm" json:"sertm"`
		Vhid  string        `bson:"vhid" json:"vhid"`
		Evt   string        `bson:"evt" json:"evt"`
		Val   int           `bson:"val" json:"val"`
		Evts  interface{}   `bson:"evts" json:"evts"`
	}
)
