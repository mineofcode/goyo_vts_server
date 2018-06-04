package socketios

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"goyo.in/gpstracker/datamodel"
	"goyo.in/gpstracker/protocal"

	"github.com/astaxie/beego"
	"github.com/googollee/go-socket.io"
	"github.com/rs/cors"
	"gopkg.in/mgo.v2/bson"
	"goyo.in/gpstracker/models"
	"goyo.in/gpstracker/shared"
)

// var socket *socketio.Server

func init() {
	start()
}

//GetSocketIO get server
// func GetSocketIO() *socketio.Server {
// 	return socket
// }

// type (
// 	// ParamsTripdata represents the structure of our resource
// 	DeviceCommands struct {
// 		cmd      string `json:"cmd"`
// 		imei     string `json:"imei"`
// 		extra    string `json:"extra"`
// 		uid      string `json:"uid"`
// 		ucode    string `json:"ucode"`
// 		platform string `json:"platform"`
// 		ip       string `json:"ip"`
// 	}
// )

func start() {

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.On("connection", func(so socketio.Socket) {
		//log.Println("on connection")
		so.Emit("msgd", bson.M{"evt": "regreq"})

		so.On("reg_v", func(msg string) {
			// client.
			vhids := strings.Split(msg, ",")
			//fmt.Println(vhids)
			for i := 0; i < len(vhids); i++ {
				so.Join(vhids[i])
			}
			so.Emit("msgd", bson.M{"evt": "registered"})
		})

		so.On("cmd", func(msg string) string {
			// client.
			var message datamodel.DeviceCommands
			raw := json.RawMessage(msg)
			err := json.Unmarshal(raw, &message)
			if err != nil {
				fmt.Println("There was an error:", err)
				return "There was an error:" + err.Error()
			}
			return sendCommnd(message)
		})

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	shared.Socket = server

	mux := http.NewServeMux()
	mux.Handle("/socket.io/", server)

	crs := cors.AllowAll()
	handler := crs.Handler(mux)
	beego.Handler("/socket.io/", handler)
	//http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5658...")
	//http.ListenAndServe(":5658", handler)
}

func sendCommnd(msg datamodel.DeviceCommands) string {
	ip := models.GetVehicleIP(msg.Imei)
	_ClientsMod := shared.Clients[ip]
	if _ClientsMod.Con != nil {
		models.StoreCommandLog(msg)
		protocalHandler.SendCommand(_ClientsMod.Con, msg)
		return "Command send successfully"
	}
	return "Error while sending command"
}
