package finally

import "fmt"

// DeferSample01 演示defer在返回前执行
func DeferSample01() (r int) {
	defer func() {
		r++
	}()
	return 0
}

// DeferSample02 实际上DeferSample01是这种实现顺序
func DeferSample02() (r int) {
	r = 0
	func() {
		r++
	}()
	return
}

// DeferSample03 演示defer执行顺序
func DeferSample03() {
	fmt.Println("defer exec order: ")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}