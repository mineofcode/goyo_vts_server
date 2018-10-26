package models

import (
	"fmt"

	"goyo.in/gpstracker/db"
)

func CreateGeoFenceLog(d interface{}) {
	_sn := getDBSession().Copy()
	defer _sn.Close()
	c := col(_sn, db.ColLogFence)
	if dberr := c.Insert(d); dberr != nil {
		fmt.Println(dberr)
	}
}
