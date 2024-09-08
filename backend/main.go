package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
    // Connect to the WebSocket server
    conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
    if err != nil {
        log.Fatal("Error connecting to WebSocket:", err)
    }
    defer func() {
        if err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, "")); err != nil {
            log.Println("Write close error:", err)
        }
        conn.Close()
    }()

    reader := bufio.NewReader(os.Stdin)

    // Loop to continuously send and receive messages
    for {
        fmt.Print("Enter your message (type 'exit' to quit): ")
        message, _ := reader.ReadString('\n')

        // Exit the loop and close the connection if the user types 'exit'
        if message == "exit\n" {
            fmt.Println("Closing connection...")
            break
        }

        // Send the message to the server
        err = conn.WriteMessage(websocket.TextMessage, []byte(message))
        if err != nil {
            log.Println("Write error:", err)
            return
        }

        // Read the response from the server
        _, response, err := conn.ReadMessage()
        if err != nil {
            log.Println("Read error:", err)
            return
        }
        log.Printf("Received from server: %s", response)
    }
}
 