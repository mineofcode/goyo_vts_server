package models

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/db"
)

//GetNextSequence generate sequence number
func GetNextSequence(_sn *mgo.Session, key string) int {
	return createSequence(key, 1, 1, _sn)
}

//GetNextSequenceWitS With start val
func GetNextSequenceWitS(_sn *mgo.Session, key string, startval int) int {
	return createSequence(key, startval, 1, _sn)
}

//GetNextSequenceWitSAndI Start val and Increamentval
func GetNextSequenceWitSAndI(_sn *mgo.Session, key string, startval int, increament int) int {
	return createSequence(key, startval, increament, _sn)
}

func createSequence(key string, startval int, incval int, _sn *mgo.Session) int {
	change := mgo.Change{}
	var restart = false
setagain:
	if restart {
		change = mgo.Change{
			Update:    bson.M{"$set": bson.M{"value": startval}},
			Upsert:    true,
			ReturnNew: true,
		}
	} else {

		change = mgo.Change{
			Update:    bson.M{"$inc": bson.M{"value": incval}},
			Upsert:    true,
			ReturnNew: true,
		}
	}

	c := col(_sn, db.ColSequence)
	var doc = datamodel.Counters{}
	_, _ = c.Find(bson.M{"_id": key}).Apply(change, &doc)

	if doc.Seq < startval {
		restart = true
		goto setagain
	}

	return doc.Seq
}

const (
	SEQVehicleID = "vehicleid"
	SEQUserID    = "uid"
	SEQSessionID = "session"
	SEQFuelId    = "fuelid"
)
