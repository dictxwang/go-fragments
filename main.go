package main

import (
	"go-fragments/exception"
)

func main() {
	// 演示错误的recover使用方式
	//exception.PanicSample01()

	// 演示正确的recover使用方式
	exception.PanicSample02()
}
