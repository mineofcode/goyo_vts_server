package main

import (
	"fmt"
	"time"
)

func main1() {

	location, _ := time.LoadLocation("Asia/Kolkata")
	// this should give you time in location
	t := time.Now().In(location)

	var GMTIME = t.AddDate(0, 0, -1)
	fmt.Println(GMTIME)

	var utctime = time.Now().UTC().AddDate(0, 0, -1)
	fmt.Println(utctime)
	fmt.Println(GMTIME.Equal(utctime))

}
