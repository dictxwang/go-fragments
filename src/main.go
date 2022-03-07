package main

import (
	"fmt"
	"go-fragments/src/channel"
	"go-fragments/src/context"
	"go-fragments/src/exception"
	"go-fragments/src/finally"
	"go-fragments/src/method"
	"go-fragments/src/sync"
)

// 获取编译参数 ldflags
var ENV string
var VERSION string

func main() {
	fmt.Printf("ENV=%s\n", ENV)
	fmt.Printf("VERSION=%s\n", VERSION)

	// 演示错误的recover使用方式
	//exception.PanicSample01()

	// 演示正确的recover使用方式
	exception.PanicSample02()

	// 演示defer在返回前执行
	fmt.Printf("DeferSample01 return: %d\n", finally.DeferSample01())
	fmt.Printf("DeferSample02 return: %d\n", finally.DeferSample02())
	// 演示defer执行顺序（先进后出的顺序）
	finally.DeferSample03()

	person := method.Person{}
	person.ChangeName("liudehua")
	fmt.Println(person.GetName())

	person.ChangeNameWithPointer("liudehua")
	fmt.Println(person.GetName())

	channel.ChanMethod1()
	channel.FindEvenNumberAndSquare(10)

	sync.UseWaitGroup()

	context.ContextHandle1()
}
