// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/googollee/go-socket.io"
// 	"github.com/rs/cors"
// )

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write([]byte("{\"hello\": \"world\"}"))
// 	})

// 	// cors.Default() setup the middleware with default options being
// 	// all origins accepted with simple methods (GET, POST). See
// 	// documentation below for more options.
// 	server, err := socketio.NewServer(nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	server.On("connection", func(so socketio.Socket) {
// 		so.Emit("chat", "hello message")
// 		log.Println("on connection")
// 		so.On("chat", func(msg string) {
// 			log.Println("recieved message", msg)
// 		})
// 		// Socket.io acknowledgement example
// 		// The return type may vary depending on whether you will return
// 		// For this example it is "string" type
// 		so.On("chat message with ack", func(msg string) string {
// 			return msg
// 		})
// 		so.On("disconnection", func() {
// 			log.Println("disconnected from chat")
// 		})
// 	})
// 	server.On("error", func(so socketio.Socket, err error) {
// 		log.Println("error:", err)
// 	})

// 	mux.Handle("/socket.io/", server)
// 	mux.Handle("/assets", http.FileServer(http.Dir("./assets")))

// 	// provide default cors to the mux
// 	handler := cors.Default().Handler(mux)

// 	c := cors.New(cors.Options{
// 		AllowedOrigins:   []string{"*"},
// 		AllowCredentials: true,
// 	})

// 	// decorate existing handler with cors functionality set in c
// 	handler = c.Handler(handler)

// 	log.Println("Serving at localhost:5000...")
// 	log.Fatal(http.ListenAndServe(":6980", handler))

// }
package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/googollee/go-socket.io"
// 	"github.com/rs/cors"
// )

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write([]byte("{\"hello\": \"world\"}"))
// 	})

// 	// cors.Default() setup the middleware with default options being
// 	// all origins accepted with simple methods (GET, POST). See
// 	// documentation below for more options.
// 	server, err := socketio.NewServer(nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	server.On("connection", func(so socketio.Socket) {
// 		so.Emit("chat", "hello message")
// 		log.Println("on connection")
// 		so.On("chat", func(msg string) {
// 			log.Println("recieved message", msg)
// 		})
// 		// Socket.io acknowledgement example
// 		// The return type may vary depending on whether you will return
// 		// For this example it is "string" type
// 		so.On("chat message with ack", func(msg string) string {
// 			return msg
// 		})
// 		so.On("disconnection", func() {
// 			log.Println("disconnected from chat")
// 		})
// 	})
// 	server.On("error", func(so socketio.Socket, err error) {
// 		log.Println("error:", err)
// 	})

// 	mux.Handle("/socket.io/", server)
// 	mux.Handle("/assets", http.FileServer(http.Dir("./assets")))

// 	// provide default cors to the mux
// 	handler := cors.Default().Handler(mux)

// 	c := cors.New(cors.Options{
// 		AllowedOrigins:   []string{"*"},
// 		AllowCredentials: true,
// 	})

// 	// decorate existing handler with cors functionality set in c
// 	handler = c.Handler(handler)

// 	log.Println("Serving at localhost:5000...")
// 	log.Fatal(http.ListenAndServe(":6980", handler))

// }
