package network

import (
	"fmt"
	"io"
	"log"
	"net"
)

// Client holds info about connection
type Client struct {
	Con    net.Conn
	Server *server
}

// TCP server
type server struct {
	address                  string // Address to open connection: localhost:9999
	onNewClientCallback      func(c *Client)
	onClientConnectionClosed func(c *Client, err error)
	onNewMessage             func(c *Client, message []byte)
}

// Read client data from channel
func (c *Client) listen() {
	//reader := bufio.NewReader(c.Con)
	for {

		buf := make([]byte, 1024)
		// Read the incoming connection into the buffer.
		n, err := c.Con.Read(buf)
		if err != nil {
			//if file/socket is closed remove the socket from list.
			if io.EOF == err {
				fmt.Printf("connection closed ip address:%s\n", c.Con.RemoteAddr().String())
				//this.removeClient(connection)
				c.Con.Close()
				c.Server.onClientConnectionClosed(c, err)
				break
			} else {
				//fmt.Println(err)
				//protocalHandler.RemoveClient(connection)
				continue
			}
		} else {
			if n > 0 {
				//data := string(buf[:n])
				if err != nil {
					fmt.Println("Error reading:", err.Error())
				}
				// code block to handle incoming data

				//protocalHandler.ParseData(buf, n, connection)
				//connection.Write([]byte(data + "\n"))
				c.Server.onNewMessage(c, buf)
			}
		}

		// message, err := reader.ReadBytes('\n')
		// if err != nil {
		// 	c.Con.Close()
		// 	c.Server.onClientConnectionClosed(c, err)
		// 	return
		// }
		// c.Server.onNewMessage(c, message)
	}
}

// Send text message to client
func (c *Client) Send(message string) error {
	_, err := c.Con.Write([]byte(message))
	return err
}

// Send bytes to client
func (c *Client) SendBytes(b []byte) error {
	_, err := c.Con.Write(b)
	return err
}

func (c *Client) Conn() net.Conn {
	return c.Con
}

func (c *Client) Close() error {
	return c.Con.Close()
}

// Called right after server starts listening new client
func (s *server) OnNewClient(callback func(c *Client)) {
	s.onNewClientCallback = callback
}

// Called right after connection closed
func (s *server) OnClientConnectionClosed(callback func(c *Client, err error)) {
	s.onClientConnectionClosed = callback
}

// Called when Client receives new message
func (s *server) OnNewMessage(callback func(c *Client, message []byte)) {
	s.onNewMessage = callback
}

// Start network server
func (s *server) Listen() {
	listener, err := net.Listen("tcp", s.address)
	if err != nil {
		log.Fatal("Error starting TCP server.")
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		client := &Client{
			Con:    conn,
			Server: s,
		}
		go client.listen()
		s.onNewClientCallback(client)
	}
}

// Creates new tcp server instance
func NewTCP(address string) *server {
	log.Println("Creating server with address", address)
	server := &server{
		address: address,
	}

	server.OnNewClient(func(c *Client) {})
	server.OnNewMessage(func(c *Client, message []byte) {})
	server.OnClientConnectionClosed(func(c *Client, err error) {})

	return server
}

//*socekt server commands*/
