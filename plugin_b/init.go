package main

import (
	"github.com/mburbidg/goplug/ops"
)

var OperationList ops.OpList

func init() {
	OperationList = ops.OpList{}
	OperationList["*"] = multiply
	OperationList["/"] = divide
}

func multiply() string {
	return "multiply operation"
}

func divide() string {
	return "divide operation"
}
