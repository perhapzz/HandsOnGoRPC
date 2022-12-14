// server.go
package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// server.go

// HelloService is a RPC service for helloWorld
type HelloService struct{}

// Hello say hello to request
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	// 用将给客户端访问的名字和HelloService实例注册 RPC 服务
	rpc.RegisterName("HelloService", new(HelloService))

	// TCP 服务
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		// 开一个协程处理连接上的接口
		go jsonrpc.ServeConn(conn)
	}
}
