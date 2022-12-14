// 简单计算器实现

package main

import (
	"errors"
)

/* 抽象的计算函数类型 */

// Operation 是计算的抽象
// 定义长 func(Number1, Number2 float64) float64 这样的可以用 Operation 来代替，并且是导出的
type Operation func(Number1, Number2 float64) float64

/* 加减乘除的具体 Operation 实现 */

// Add 是加法的 Operation 实现
func Add(Number1, Number2 float64) float64 {
	return Number1 + Number2
}

// Sub 是减法的 Operation 实现
func Sub(Number1, Number2 float64) float64 {
	return Number1 - Number2
}

// Mul 是乘法的 Operation 实现
func Mul(Number1, Number2 float64) float64 {
	return Number1 * Number2
}

// Div 是除法的 Operation 实现
func Div(Number1, Number2 float64) float64 {
	return Number1 / Number2
}

/* 工厂 */

// Operators 注册所有支持的运算
var Operators = map[string]Operation{
	"+": Add,
	"-": Sub,
	"*": Mul,
	"/": Div,
}

// CreateOperation 通过 string 表示的 operator 获取适合的 Operation 函数
func CreateOperation(operator string) (Operation, error) {
	oper, ok := Operators[operator]
	if ok {
		return oper, nil
	}
	return oper, errors.New("illegal Operator")
}
