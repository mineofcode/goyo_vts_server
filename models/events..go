package models

import (
	"fmt"

	"goyo.in/gpstracker/db"
)

func CreateEvent(d interface{}, vhid string) {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	ch := col(_sn, db.ColVHevts)
	err := ch.Insert(d)

	if err != nil {
		fmt.Println("Event Creation Error", err)
	}

}
