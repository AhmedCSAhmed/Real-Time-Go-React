package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

/*
	creating a simple websocket server

	this endpoint will do 3 things,
	1st it'll check the origin of the incoming HTTP request and return True every time to open up our endpoint to every potential client
	2nd it'll upgrade the incoming HTTP request to a websocket connection


*/
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	/*
	CheckOrigin returns true if the request Origin header is acceptable. 
	If CheckOrigin is nil, 
	then a safe default is used: return false if the Origin request header is present 
	and the origin host is not equal to request Host header.
	*/
	CheckOrigin:  func(r *http.Request) bool {return true}, // doing no ckecking here just letting anyhting go through
}

// our reader will listen for new messaged being sent to our websocket

func reader(conn *websocket.Conn) {
    defer conn.Close() // Ensure the connection is properly closed
    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println("Read error:", err)
            return
        }
        fmt.Println("Received:", string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println("Write error:", err)
            return
        }
    }
}




// Will create our websocket endpoint 
// the reader is a function that will listen for new messages sent to our websocket endpoint
func serveWS(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Host) // For server requests, Host specifies the host on which the URL is sought
	// Will upgrade this connection to a websocket connection 
	ws, err := upgrader.Upgrade(w, r, nil) // Upgrade upgrades the HTTP server connection to the WebSocket protocol.
	if err != nil{
		log.Println(err) // If there is an error, log it
		return
	}
	fmt.Println("Client Connected") // If the connection is successful, print this message
	// will listen indefinitely for new messages
	reader(ws) // Call the reader function and pass in the websocket connection
}

// setting up our routes
// what handleFunc does is it takes a path and a function as arguments and maps the path to the function called handler in this case

func setupRoutes(){
	/*
	what handleFunc does is it takes a path and a function as arguments and maps the path to the function called handler in this case
	// the path is "/" and the function is handler and the handler is a function that takes a response writer and a request as arguments 
	and returns nothing all it does is write "simple server" to the response writer

	the response writer is the object that the server uses to write the response to the client

	*/
	http.HandleFunc("/ws", serveWS) // Will map the /ws path to the serveWS function

}
func main() {
	fmt.Println("Chat App v0.01")
	setupRoutes() // Call the setupRoutes function
	http.ListenAndServe(":8080", nil) // ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux.
}




