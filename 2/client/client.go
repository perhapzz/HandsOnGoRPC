package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// HelloServiceName is the name of HelloService
const HelloServiceName = "HelloService"

// HelloServiceClient is a client for HelloService
type HelloServiceClient struct {
	*rpc.Client
}

// HelloServiceInterface is a interface for HelloService
type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

var _ HelloServiceInterface = (*HelloServiceClient)(nil)

// DialHelloService dial HelloService
func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

// Hello calls HelloService.Hello
func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(HelloServiceName+".Hello", request, reply)
}

func main() {
	client, err := DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("world", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
