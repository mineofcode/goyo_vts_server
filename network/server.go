package network

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"

	"goyo.in/gpstracker/protocal"
)

type TCPServer struct {
	Host     string
	Port     string
	Timeout  int
	Clients  map[string]net.Conn //keep ip address in map key
	listener net.Listener
	locker   sync.Mutex
}

func (this *TCPServer) Closed(client net.Conn) error {
	return nil
}
func (this *TCPServer) PingAll() {
	for _, conn := range this.Clients {
		conn.Write([]byte("p" + "\n"))
	}
	time.Sleep(1 * time.Minute)
	this.PingAll()
}
func (this *TCPServer) addClient(conn net.Conn) {
	this.locker.Lock()
	defer this.locker.Unlock()

	ip_address := conn.RemoteAddr().String()
	this.Clients[ip_address] = conn
	fmt.Println(len(this.Clients))
}
func (this *TCPServer) removeClient(conn net.Conn) {
	this.locker.Lock()
	defer this.locker.Unlock()

	ip_address := conn.RemoteAddr().String()
	delete(this.Clients, ip_address)
}
func (this *TCPServer) Open() error {
	//this.Clients = make(map[string]net.Conn)

	go func() {

		l, err := net.Listen("tcp", this.Host+":"+this.Port)

		if err != nil {
			fmt.Println("Error listening:", err.Error())
			os.Exit(1)
		}
		this.listener = l
		fmt.Println("Listening on " + this.Host + ":" + this.Port)
		for {
			// Listen for an incoming connection.
			conn, err := this.listener.Accept()
			if err != nil {
				fmt.Println("Error accepting: ", err.Error())
				continue
			}
			go this.handleRequest(conn)
			//this.addClient(conn)
		}
	}()
	//go this.PingAll()

	return nil
}

func (this *TCPServer) handleRequest(connection net.Conn) error {
	fmt.Println(connection.RemoteAddr().String())
	for {
		// make a buffer to hold incoming data
		buf := make([]byte, 1024)
		// Read the incoming connection into the buffer.
		n, err := connection.Read(buf)
		if err != nil {
			//if file/socket is closed remove the socket from list.
			if io.EOF == err {
				fmt.Printf("connection closed ip address:%s\n", connection.RemoteAddr().String())
				//this.removeClient(connection)
				protocalHandler.RemoveClient(connection)
				break
			} else {
				//fmt.Println(err)
				protocalHandler.RemoveClient(connection)
				continue
			}
		} else {
			if n > 0 {
				//data := string(buf[:n])
				if err != nil {
					fmt.Println("Error reading:", err.Error())
				}
				// code block to handle incoming data

				protocalHandler.ParseData(buf, n, connection)
				//connection.Write([]byte(data + "\n"))
			}
		}

	}
	return nil
}
func (this *TCPServer) Close() error {
	this.Clients = nil
	return this.listener.Close()
}
