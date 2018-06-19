package db

import (
	"time"

	mgo "gopkg.in/mgo.v2"
	"goyo.in/gpstracker/const"
)

const (
	dbip   = consts.MGODbip
	dbport = consts.MGODbport
	//Dbname Mongodb database name
	Dbname = "goyosch"
	//Dbname = "19j"

	urldb        = dbip + ":" + dbport
	authUserName = ""
	authPassword = ""
)

var mongoDBDialInfo = &mgo.DialInfo{
	Addrs:    []string{urldb},
	Timeout:  2 * time.Minute,
	Database: Dbname,
}
