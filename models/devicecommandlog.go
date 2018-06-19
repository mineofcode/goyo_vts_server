package models

import (
	"fmt"
	"time"

	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/db"
)

func StoreCommandLog(msg datamodel.DeviceCommands) {
	_sn := getDBSession().Copy()
	defer _sn.Close()
	msg.Time = time.Now()
	//fmt.Println(vhid, d)
	c := col(_sn, db.ColCommandsLog)
	if dberr := c.Insert(msg); dberr != nil {
		fmt.Println(dberr)
		if dberr.Error() == "not found" {

		}
	}

}
