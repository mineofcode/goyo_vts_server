package patterns

import (
	"time"

	"gopkg.in/mgo.v2"
	"goyo.in/gpstracker/const"
)

const (
	dbip   = consts.MGODbip
	dbport = consts.MGODbport
	//Dbname Mongodb database name
	Dbname = "goyosch"
	//Dbname = "19j"

	//ColVhcls 'vhcls' Collection
	ColVhcls = "vhcls"
	//ColVhtrps 'vhtrps' Collection
	ColVhtrps = "vhtrps"
	//ColHistory 'vhdyhst' Collection
	ColHistory = "vhdyhst"
	//ColGeofence 'geofnc' Collection
	ColGeofence = "geofnc"

	urldb        = dbip + ":" + dbport
	authUserName = ""
	authPassword = ""
)

var mongoDBDialInfo = &mgo.DialInfo{
	Addrs:    []string{urldb},
	Timeout:  5 * time.Second,
	Database: Dbname,
}
