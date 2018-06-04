package models

import (
	"fmt"
	"time"

	"goyo.in/gpstracker/datamodel"

	patterns "goyo.in/gpstracker/patterns"
)

func StoreCommandLog(msg datamodel.DeviceCommands) {
	_sn := getDBSession().Copy()
	defer _sn.Close()
	msg.Time = time.Now()
	//fmt.Println(vhid, d)
	c := col(_sn, patterns.ColCommandsLog)
	if dberr := c.Insert(msg); dberr != nil {
		fmt.Println(dberr)
		if dberr.Error() == "not found" {

		}
	}

}
