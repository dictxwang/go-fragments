package channel

import (
	"fmt"
	"time"
)

func doChanMethod1(ch chan bool) {
	fmt.Println("ChanMethod1 executing...")
	time.Sleep(time.Second)
	ch <- true
}

func chanMethod1() {
	ch := make(chan bool)
	go doChanMethod1(ch)
	<-ch
}

func doFindEven(max int, findChannel chan int) {
	for i := 1; i <= max; i++ {
		if i % 2 == 0 {
			findChannel <- i
		}
	}
	close(findChannel)
}

func doSquare(findChannel chan int, printChannel chan int) {
	for v := range findChannel {
		printChannel <- v * v
	}
	close(printChannel)
}

func doPrint(printChannel chan int) {
	for i := range printChannel {
		fmt.Println(i)
	}
}

func findEvenNumberAndSquare(max int) {
	findChannel := make(chan int)
	printChannel := make(chan int)
	go doFindEven(max, findChannel)
	go doSquare(findChannel, printChannel)
	// doPrint需要回到主协程执行
	doPrint(printChannel)
}

func SampleMain() {

	fmt.Println("\n[channel_sample]")
	chanMethod1()
	findEvenNumberAndSquare(10)
}