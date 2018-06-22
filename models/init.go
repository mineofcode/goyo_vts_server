package models

import (
	validator "gopkg.in/go-playground/validator.v9"
	mgo "gopkg.in/mgo.v2"
	"goyo.in/gpstracker/db"
)

var validate *validator.Validate

//Init Database
func Init() {

	//initial validator

	validate = validator.New()

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

	//indexing for devinv
	ColSimEnv := col(_sn, db.ColSimEnv)
	index = mgo.Index{
		Key:        []string{"mobno"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = ColSimEnv.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

}

func GetValidator() *validator.Validate {
	return validate
}
