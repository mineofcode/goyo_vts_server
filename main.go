package main

import (
	"strconv"

	"goyo.in/gpstracker/const"
	"goyo.in/gpstracker/datashare"
	"goyo.in/gpstracker/tile"

	"goyo.in/gpstracker/models"
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

	// jobs for history creation
	go job.StartJob()

	//go redigogeofence.Start() client for tile18 server
	go tile.GRpcRun()
	//Start Rest API & Socket.io server
	go func() {
		restservice.RestfulAPIServiceInit("HTTP")
	}()

	models.Init()
	datashare.Init()

	en := network.TCPServer_new{Host: consts.TCPHost, Port: strconv.Itoa(consts.TCPPort), Timeout: 3000}
	en.StartServer()

}
