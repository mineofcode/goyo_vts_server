package models

import (
	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/db"
	"goyo.in/gpstracker/utils"
)

type (
	STRUCTMom struct {
		Value  string      `json:"value"`
		Valued interface{} `json:"valued"`
		Active bool        `json:"active"`
	}
)

func GetMoM(search map[string]interface{}) utils.Response {
	response := utils.Response{}
	_sn := getDBSession().Copy()
	defer _sn.Close()

	c := col(_sn, db.ColMoM)
	var dResult1 []STRUCTMom
	err := c.Find(search).All(&dResult1)
	if err != nil {

		response.Error = err.Error()
		return response
	}
	response.Status = true
	response.Data = dResult1
	return response
}

func SaveMoM(d map[string]interface{}) utils.Response {
	response := utils.Response{}
	_sn := getDBSession().Copy()
	defer _sn.Close()

	c := col(_sn, db.ColMoM)

	info, err := c.Upsert(bson.M{"group": d["group"], "key": d["key"]}, bson.M{"$set": d})

	msg := ""
	if info.Updated > 0 {
		msg = "Updated Successfully"
	} else if info.UpsertedId != nil {
		msg = "Inserted Successfully"
	}
	if err != nil {
		response.Status = false
		response.Error = err.Error()
	} else {
		response.Status = true
		response.Message = msg
	}
	return response
}
