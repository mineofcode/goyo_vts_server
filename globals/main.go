package globals

import (
	"net"
	"sync"
	"time"
	"fmt"
)


type Global struct {
}

//Clients keep ip address in map key
var Clients = make(map[string]clientsMod)

var locker sync.Mutex

type clientsMod struct {
	Con       net.Conn
	Imei      string
	Lstm      time.Time
	Allwspd   int
	Geofences []interface{}
}

//Check client is dead or not
func (g Global) checkClientLiveStatus() {
	//fmt.Println("clled")
	locker.Lock()
	defer locker.Unlock()
	for k, v := range Clients {
		tm := time.Now().Sub(v.Lstm)
		//fmt.Println(tm)
		if tm > time.Minute*15 {

			delete(Clients, k)
			fmt.Println("Client Deleted ", k)
		}
	}
}



//Terminal connection persists
func (g Global) AddClient(conn net.Conn, client string, allwspd int) clientsMod {
	locker.Lock()
	defer locker.Unlock()
	ip_address := conn.RemoteAddr().String()
	fmt.Println("Client Added  ", ip_address)
	Clients[ip_address] = clientsMod{Con: conn, Imei: client, Lstm: time.Now(), Allwspd:allwspd }
	return Clients[ip_address]
}



func (g Global) RemoveClient(conn net.Conn) {
	locker.Lock()
	defer locker.Unlock()

	ip_address := conn.RemoteAddr().String()
	delete(Clients, ip_address)
}

func (g Global) GetClient(conn net.Conn) (client clientsMod) {
	ip_address := conn.RemoteAddr().String()
	_clientsMod := Clients[ip_address]
	return _clientsMod
}