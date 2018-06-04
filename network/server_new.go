package network

import (
	"fmt"
	"net"
	"time"

	"github.com/jasonlvhit/gocron"
	"goyo.in/gpstracker/protocal"
	"goyo.in/gpstracker/shared"
)

type TCPServer_new struct {
	Host     string
	Port     string
	Timeout  int
	Clients  map[string]net.Conn //keep ip address in map key
	listener net.Listener
	serv     *server
}

func startNonWorkinClinetLookUp() {
	go func() {
		fmt.Println("Client Garbage Started")
		s := gocron.NewScheduler()
		s.Every(10).Minutes().Do(checkClientLiveStatus)
		<-s.Start()
	}()
}

//Check client is dead or not
func checkClientLiveStatus() {
	for k, v := range shared.Clients {
		tm := time.Now().Sub(v.Lstm)
		fmt.Println(v)
		if tm > time.Minute*20 {
			shared.Locker.Lock()
			defer shared.Locker.Unlock()
			err := v.Con.Close()
			if err != nil {
				fmt.Println("Client Close Error ", err)
			}
			protocalHandler.RemoveClient(v.Con)
			fmt.Println("Client Deleted ", k)
		}
	}
}

// func (this *TCPServer_new) addClient(conn net.Conn) {
// 	this.locker.Lock()
// 	defer this.locker.Unlock()

// 	ip_address := conn.RemoteAddr().String()
// 	this.Clients[ip_address] = conn
// 	fmt.Println(len(this.Clients))
// }

// func (this *TCPServer_new) removeClient(conn net.Conn) {
// 	this.locker.Lock()
// 	defer this.locker.Unlock()

// 	ip_address := conn.RemoteAddr().String()
// 	delete(this.Clients, ip_address)
// }

func (this *TCPServer_new) StartServer() {
	//startNonWorkinClinetLookUp()

	this.serv = NewTCP(this.Host + ":" + this.Port)

	this.serv.OnNewClient(func(c *Client) {
		// new client connected
		// lets send some message
		//c.Send("Hello")
		//protocalHandler.addClient(c.Con,)
	})
	this.serv.OnNewMessage(func(c *Client, message []byte) {
		// new message received
		protocalHandler.ParseData(message, 0, c.Con)

	})
	this.serv.OnClientConnectionClosed(func(c *Client, err error) {
		// connection with client lost
		protocalHandler.RemoveClient(c.Con)
	})

	this.serv.Listen()

}
