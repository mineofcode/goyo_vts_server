package models

import "gopkg.in/mgo.v2/bson"


import "time"

type (  
    // User represents the structure of our resource
    User struct {
		Id         	bson.ObjectId   `bson:"_id,omitempty" json:"-"`
		Sertm  		time.Time       `bson:"sertm" json:"-"`
		Actvt	 	string          `bson:"actvt" json:"actvt"`
	}
)