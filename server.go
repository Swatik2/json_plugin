package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	// "encoding/json"
	"encoding/base64"
	"strings"
)
// "encoding/base64"
type Data struct {
	Plugin_type           string
	Local_resource_Source string
}

// type JSONResponse struct {
// 	Result string
// }

type RPCMethods struct{}

func (m *RPCMethods) GetData(args interface{}, reply *Data) error {
	data := Data{
		Plugin_type:           "json",
		Local_resource_Source: "/home/vboxuser/Documents/source_dir",
	}

	*reply = data
	return nil
}

func (r *RPCMethods) ProcessJSON(jsonData string, res *string) error {
	// Process the JSON data on the server
	decodedBytes, err := base64.StdEncoding.DecodeString(jsonData)
	if err != nil {
		fmt.Println("Error decoding Base64 string:", err)
	}

	// var data map[string]interface{}

	// err = json.Unmarshal(decodedBytes, &data)
	// if err != nil {
	// 	fmt.Println("Error unmarshaling JSON:", err)
	// }

	// count, ok := data["count"].(float64)
	// if !ok {
	// 	fmt.Println("Invalid count value")
	// }
	// fmt.Println(decodedBytes)
	// fieldValue := openLineageForm.name
	// fmt.Println("Field value:", fieldValue)

	jsonStr := string(decodedBytes)
	trimmedStr := strings.TrimSpace(jsonStr)
	fmt.Println(len(trimmedStr))
	if len(trimmedStr) < 10 {
		*res = "No JSON to process"
	} else{
		fmt.Println("Decoded JSON string:", jsonStr)
		*res = "Processed JSON successfully"
	}
	// log.Printf("Received JSON data on server: %s\n", jsonData)

	

	return nil
}

func main() {
	rpcMethods := new(RPCMethods)
	rpc.Register(rpcMethods)

	// Start the RPC server
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("Error starting RPC server:", err)
	}

	fmt.Println("RPC server started on port 1234")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("Error accepting connection:", err)
		}
		go jsonrpc.ServeConn(conn)
	}
}