package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
