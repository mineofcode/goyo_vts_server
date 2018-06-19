package models

import (
	mgo "gopkg.in/mgo.v2"
	"goyo.in/gpstracker/db"
)

//Init Database
func Init() {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	// User Collection
	ColUsers := col(_sn, db.ColUsers)

	index := mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := ColUsers.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	//indexing for devinv
	ColDeviceEnv := col(_sn, db.ColDeviceEnv)
	index = mgo.Index{
		Key:        []string{"imei"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err = ColDeviceEnv.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

}
