package shared

import (
	"net"
	"sync"
	"time"

	"github.com/googollee/go-socket.io"
)

var Locker sync.Mutex
var Clients = make(map[string]ClientsMod)

type ClientsMod struct {
	Con       net.Conn
	Imei      string
	Lstm      time.Time
	Allwspd   int
	Geofences []interface{}
}

var Socket *socketio.Server
