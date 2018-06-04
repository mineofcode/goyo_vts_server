package bridges

import (
	"goyo.in/gpstracker/socketios"
)

// func SendCommandToNetwork(msg datamodel.DeviceCommands) {
// 	ip := models.GetVehicleIP(msg.imei)
// 	con := network.TCPServer_new.Clients[ip]
// 	if con != nil {
// 		protocalHandler.SendCommand(con, msg.cmd)
// 	}
// }

func sendCommand() {

}

func SendLive() {
	socket := socketios.GetSocketIO()
	socket.BroadcastTo(vhid, "msgd", data)
}
