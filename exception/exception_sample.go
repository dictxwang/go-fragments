package exception

import "fmt"

func PanicSample01() {
	fmt.Println("this is a panic example")
	panic("this is a panic")
	// 这种方式recover捕获异常会失效，程序会终止
	r := recover()
	fmt.Printf("panic recover: %s", r)
}

func PanicSample02() {
	fmt.Println("this is a panic example")
	// 正解是recover捕获异常需要和defer配合使用，捕获成功，程序继续执行
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic recover: %s", r)
		}
	}()
	panic("this is a panic")
}
