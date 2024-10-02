package finally

import "fmt"

// DeferSample01 演示defer在返回前执行
func deferSample01() (r int) {
	defer func() {
		r++
	}()
	return 10 // 最终返回11(10+1)
}

// DeferSample02 实际上DeferSample01是这种实现顺序
func deferSample02() (r int) {
	r = 10
	func() {
		r++
	}()
	return
}

// DeferSample03 演示defer执行顺序
func deferSample03() {
	fmt.Println("defer exec order: ")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
}

func SampleMain() {

	fmt.Println("\n[finally_sample]")
	// 演示defer在返回前执行
	fmt.Printf("DeferSample01 return: %d\n", deferSample01()) // 11
	fmt.Printf("DeferSample02 return: %d\n", deferSample02()) // 11
	// 演示defer执行顺序（先进后出的顺序）
	deferSample03()
}
