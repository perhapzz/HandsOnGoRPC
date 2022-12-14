// server.go
package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
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
	RegisterHelloService(new(HelloService))

	// HTTP 服务
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)
}
