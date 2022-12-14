package main

import (
	calc "HansOnGoRPC/3_Clac/calc_tcp"
	"log"
	"net/rpc"
)

/* 定义客户端实现 */

// CalcClient is a client for CalcService
type CalcClient struct {
	*rpc.Client
}

var _ calc.ServiceInterface = (*CalcClient)(nil)

// DialCalcService dial CalcService
func DialCalcService(network, address string) (*CalcClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &CalcClient{Client: c}, nil
}

// CalcTwoNumber 对两个数进行运算
func (c *CalcClient) CalcTwoNumber(request calc.Calc, reply *float64) error {
	return c.Client.Call(calc.ServiceName+".CalcTwoNumber", request, reply)
}

// GetOperators 获取所有支持的运算
func (c *CalcClient) GetOperators(request struct{}, reply *[]string) error {
	return c.Client.Call(calc.ServiceName+".GetOperators", request, reply)
}

/* 使用客户端调用 RPC 服务 */

func main() {
	client, err := DialCalcService("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("Err Dial Client:", err)
	}

	// Test GetOperators
	var opers []string
	err = client.GetOperators(struct{}{}, &opers)
	if err != nil {
		log.Println(err)
	}
	log.Println(opers)

	// Test CalcTwoNumber
	testAdd := calc.Calc{
		Number1:  2.0,
		Number2:  3.14,
		Operator: "+",
	}
	var result float64
	client.CalcTwoNumber(testAdd, &result)
	log.Println(result)
}
