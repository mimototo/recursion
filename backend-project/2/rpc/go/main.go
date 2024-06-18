package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
)

// Request represents a JSON-RPC request
type Request struct {
	Method     string        `json:"method"`
	Params     []interface{} `json:"params"`
	ParamTypes []string      `json:"param_types"`
	ID         int           `json:"id"`
}

// Response represents a JSON-RPC request
type Response struct {
	Results    string `json:"result"`
	ResultType string `json:"result_type"`
	ID         int    `json:"id"`
	Error      string `json:"error,omitempty"`
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	var req Request
	if err := decoder.Decode(&req); err != nil {
		sendErrorResponse(conn, req.ID, err.Error())
		return
	}

	var result interface{}
	var resultType string

	// Dispatch the request to the appropriate function
	switch req.Method {
	case "floor":
		if len(req.Params) == 1 && req.ParamTypes[0] == "double" {
			param, ok := req.Params[0].(float64)
			if !ok {
				sendErrorResponse(conn, req.ID, "Invalid parameter type")
				return
			}
			result = int(param)
			resultType = "int"
		} else {
			sendErrorResponse(conn, req.ID, "Invalid parameters")
			return
		}
	default:
		sendErrorResponse(conn, req.ID, "Unknown method")
		return
	}

	// Convert the result to a string
	var resultStr string
	switch v := result.(type) {
	case int:
		resultStr = strconv.Itoa(v)
	case string:
		resultStr = v
	default:
		resultStr = fmt.Sprintf("%v", v)
	}

	// Send the response
	response := Response{Results: resultStr, ResultType: resultType, ID: req.ID}
	encoder := json.NewEncoder(conn)
	if err := encoder.Encode(&response); err != nil {
		log.Println("Error encoding response:", err)
	}
}

func sendErrorResponse(conn net.Conn, id int, errMsg string) {
	response := Response{ID: id, Error: errMsg}
	encoder := json.NewEncoder(conn)
	if err := encoder.Encode(&response); err != nil {
		log.Println("Error encoding response:", err)
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
