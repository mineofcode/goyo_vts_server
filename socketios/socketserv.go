package socketios

import (
	"log"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	"github.com/googollee/go-socket.io"
	"github.com/rs/cors"
	"gopkg.in/mgo.v2/bson"
)

var socket *socketio.Server

func init() {
	start()
}

//GetSocketIO get server
func GetSocketIO() *socketio.Server {
	return socket
}

// type (
// 	// ParamsTripdata represents the structure of our resource
// 	ParamsTripdata struct {
// 		Vhids []string `json:"vhids"`
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

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})
	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	socket = server

	mux := http.NewServeMux()
	mux.Handle("/socket.io/", server)

	crs := cors.AllowAll()
	handler := crs.Handler(mux)
	beego.Handler("/socket.io/", handler)
	//http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:5658...")
	//http.ListenAndServe(":5658", handler)
}
