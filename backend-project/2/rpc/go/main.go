package main

import (
	"fmt"
	"log"
	"net"
)

// Request represents a JSON-RPC request
type Request struct {
	Method     string        `json:"method"`
	Params     []interface{} `json:"params"`
	ParamTypes []string      `json:"paramTypes"`
	ID         int           `json:"id"`
}

// Response represents a JSON-RPC request
type Response struct {
	Results    string `json:"result"`
	ResultType string `json:"resultType"`
	ID         int    `json:"id"`
	Error      string `json:"error,omitempty"`
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println("Error reading from connection:", err)
		return
	}
	fmt.Println("Received:", string(buffer[:n]))
	_, err = conn.Write([]byte("Message received"))
	if err != nil {
		log.Println("Error writing to connection:", err)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()
	log.Println("Server is listening on port 12345...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConnection(conn)
	}

}
