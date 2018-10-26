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

type Response struct {
	Status    bool        `json:"status"`
	Message   string      `json:"msg"`
	Error     string      `json:"error"`
	StausCode int         `json:"statuscode"`
	Data      interface{} `json:"data"`
	Extra     interface{} `json:"extra"`
}
