package main

import (
	calc "HansOnGoRPC/3_Clac/calc_tcp"
	"log"
	"net"
	"net/rpc"
)

/* RPC 服务实现 */

// CalcService 是计算器 RPC 服务的实现
type CalcService struct{}

// CalcTwoNumber 对两个数进行加减乘除运算
func (c *CalcService) CalcTwoNumber(request calc.Calc, reply *float64) error {
	// 因为在同一个 package 下，可以直接读取 calc.go 里的函数
	oper, err := CreateOperation(request.Operator)
	if err != nil {
		return err
	}
	*reply = oper(request.Number1, request.Number2)
	return nil
}

// GetOperators 获取所有支持的运算
func (c *CalcService) GetOperators(request struct{}, reply *[]string) error {
	opers := make([]string, 0, len(Operators))
	for key := range Operators {
		opers = append(opers, key)
	}
	*reply = opers
	return nil
}

/* 运行 RPC 服务 */

func main() {
	calc.RegisterCalcService(new(CalcService))
	// rpc.HandleHTTP()
	// http.ListenAndServe(":8080", nil)
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
