package exception

import "fmt"

func panicSample01() {
	fmt.Println("this is a panic example")
	panic("this is a panic")
	// 这种方式recover捕获异常会失效，程序会终止
	r := recover()
	fmt.Printf("panic recover: %s\n", r)
}

func panicSample02() {
	fmt.Println("this is a panic example")
	// 正解是recover捕获异常需要和defer配合使用，捕获成功，程序继续执行
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic recover: %s\n", r)
		}
	}()
	panic("this is a panic")
}

func SampleMain()  {

	fmt.Println("\n[exception_sample]")
	// 演示错误的recover使用方式，未捕获到异常程序退出
	//panicSample01()
	// 演示正确的recover使用方式
	panicSample02()
}