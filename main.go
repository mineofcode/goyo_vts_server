package main

import (
	"fmt"
	"strconv"

	"goyo.in/gpstracker/const"

	"goyo.in/gpstracker/network"
	// import "goyo.in/gpstracker/network"
	// import "goyo.in/gpstracker/crc16"

	restservice "goyo.in/gpstracker/conf"

	"goyo.in/gpstracker/job"
)

// import "goyo.in/gpstracker/network"
// import "goyo.in/gpstracker/crc16"

func main() {

	//	dataprocess.GetData()
	startAll()
}

func startAll() {
	//Start TCP server
	en := network.TCPServer{Host: consts.TCPHost, Port: strconv.Itoa(consts.TCPPort), Timeout: 3000}
	err := en.Open()

	if err != nil {
		fmt.Println("Error TCP: ", err.Error())
	}

	job.StartJob()

	//Start Rest API & Socket.io server
	restservice.RestfulAPIServiceInit("HTTP")

}
