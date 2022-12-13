package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

// HelloServiceName is the name of HelloService
const HelloServiceName = "HelloService"

// HelloServiceInterface is a interface for HelloService
type HelloServiceInterface interface {
	Hello(request string, reply *string) error
}

// RegisterHelloService register the RPC service on svc
func RegisterHelloService(svc HelloServiceInterface) error {
	// 返回rpc.RegisterName(HelloServiceName, svc)的报错信息
	return rpc.RegisterName(HelloServiceName, svc)
}

func main() {
	// rpc.RegisterName("HelloService", new(HelloService))
	RegisterHelloService(new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}
