package utils

import (
	"gopkg.in/mgo.v2/bson"
)

func CreateWrap(status string, data interface{}) bson.M {

	return bson.M{
		"status": status,
		"data":   data,
	}

}
