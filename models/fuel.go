package models

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/db"
	"goyo.in/gpstracker/utils"
)

// SaveFuelEntry check device exists and activated
func SaveFuelEntry(d map[string]interface{}) utils.Response {
	response := utils.Response{}
	_sn := getDBSession().Copy()
	defer _sn.Close()

	c := col(_sn, db.ColFuelEntry)

	if d["autoid"].(float64) == 0 {
		d["autoid"] = GetNextSequence(_sn, SEQFuelId)
		d["cron"] = time.Now()
	}

	d["upon"] = time.Now()

	var fuelentr datamodel.FuelEntry
	mapstructure.Decode(d, &fuelentr)

	validerr := GetValidator().Struct(fuelentr)

	if validerr != nil {
		response.Error = validerr.Error()
		response.Status = false
		response.StausCode = 0
		return response
	}

	info, err := c.Upsert(bson.M{"autoid": fuelentr.Autoid}, bson.M{"$set": d})

	msg := ""
	if info.Updated > 0 {
		msg = "Updated Successfully"
	} else if info.UpsertedId != nil {
		msg = "Inserted Successfully"
	} else if info.Matched > 0 {
		msg = "Not Change"
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

func GetFuelEntry(ob map[string]interface{}) utils.Response {
	response := utils.Response{}
	_sn := getDBSession().Copy()
	defer _sn.Close()

	c := col(_sn, db.ColFuelEntry)
	if ob["flag"] == "edit" {
		delete(ob, "flag")
		var result []datamodel.SelectEditFuelEntry
		c.Find(ob).All(&result)
		response.Status = true
		response.Data = result

	} else {
		delete(ob, "flag")
		var result []datamodel.SelectAllFuelEntry
		c.Find(ob).All(&result)
		response.Status = true
		response.Data = result

	}

	fmt.Println(ob)
	return response
}
