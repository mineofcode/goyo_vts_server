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
	//ColGeofence 'evt' Collection
	ColVHevts = "vhevts"
	//ColGeofence 'evt' Collection
	ColCommandsLog = "cmdlogs"

	urldb        = dbip + ":" + dbport
	authUserName = ""
	authPassword = ""
)

var mongoDBDialInfo = &mgo.DialInfo{
	Addrs:    []string{urldb},
	Timeout:  2 * time.Minute,
	Database: Dbname,
}
