package models

import (
	"fmt"

	patterns "goyo.in/gpstracker/patterns"
)

func CreateEvent(d interface{}, vhid string) {
	_sn := getDBSession().Copy()
	defer _sn.Close()

	ch := col(_sn, patterns.ColVHevts)
	err := ch.Insert(d)

	if err != nil {
		fmt.Println("Event Creation Error", err)
	}

}
