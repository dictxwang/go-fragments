package sync

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func doPrint(wg *sync.WaitGroup, msg string, t int64) {
	time.Sleep(time.Duration(t) * time.Second)
	fmt.Println(msg)
	wg.Done()
}

func UseWaitGroup() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go doPrint(wg, "use WaitGroup: " + strconv.Itoa(i), rand.Int63n(5))
	}
	wg.Wait()
	fmt.Println("main processing continue.")
}