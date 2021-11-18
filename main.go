package main

import (
	"fmt"
	"go-fragments/exception"
	"go-fragments/finally"
)

func main() {
	// 演示错误的recover使用方式
	//exception.PanicSample01()

	// 演示正确的recover使用方式
	exception.PanicSample02()

	// 演示defer在返回前执行
	fmt.Printf("DeferSample01 return: %d\n", finally.DeferSample01())
	fmt.Printf("DeferSample02 return: %d\n", finally.DeferSample02())
	// 演示defer执行顺序（先进后出的顺序）
	finally.DeferSample03()
}
