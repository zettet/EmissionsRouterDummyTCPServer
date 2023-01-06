package main

import (
	"fmt"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "3333"
	CONN_TYPE = "tcp"
)

// Implementation copied from: https://coderwall.com/p/wohavg/creating-a-simple-tcp-server-in-go
func main() {
	// Listen for incoming connections.
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	dummyMessage := []byte{
		65, 73, 82, 0, 0, 0, 6, 78, 50, 48, 57, 48, 52, 0, 0, 0,
		2, 0, 0, 0, 7, 71, 69, 110, 120, 45, 49, 66, 64, 67, 142, 214,
		235, 255, 29, 96, 192, 80, 212, 192, 142, 99, 1, 101, 64, 226, 3, 240,
		0, 0, 0, 0, 192, 74, 153, 153, 153, 153, 153, 154,
	}
	// Send a response back to person contacting us.
	conn.Write(dummyMessage)
	// Close the connection when you're done with it.
	conn.Close()
}
