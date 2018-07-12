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
	Con      net.Conn
	Imei     string
	Lstm     time.Time
	Allwspd  int
	VtsID    int
	Speed    int
	Loc      []float64
	Acc      int
	AC       int
	PClients []string
}

var Socket *socketio.Server
