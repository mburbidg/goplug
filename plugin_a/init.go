package main

import (
	"github.com/mburbidg/goplug/ops"
)

var OperationList ops.OpList

func init() {
	OperationList = ops.OpList{}
	OperationList["+"] = add
	OperationList["-"] = subtract
}

func add() string {
	return "add operation"
}

func subtract() string {
	return "subtract operation"
}
