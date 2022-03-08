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

func useWaitGroup() {
	wg := &sync.WaitGroup{}
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go doPrint(wg, "use WaitGroup: " + strconv.Itoa(i), rand.Int63n(5))
	}
	wg.Wait()
	fmt.Println("main processing continue.")
}

func useOnce() {
	once := &sync.Once{}
	for i := 1; i < 4; i++ {
		go func() {
			// 函数仅执行一次
			once.Do(func() {
				fmt.Println("once do executing...")
			})
		}()
	}
}

func useMap() {
	m := &sync.Map{}
	m.Store(1, "one")
	m.Store(2, "two")

	value, contains := m.Load(1)
	if contains {
		fmt.Printf("load: %d => %s\n", 1, value)
	}

	value, loaded := m.LoadOrStore(3, "three")
	// 如果之前map中没有3这个key，loaded将返回false
	if !loaded {
		fmt.Printf("loadOrStore: %d => %s\n", 3, value)
	}

	m.Delete(3)

	m.Range(func(key, value interface{}) bool {
		fmt.Printf("range: %d => %s\n", key.(int), value.(string))
		return true
	})
}

func syncRead(locker *sync.RWMutex, num *int) int {
	locker.RLock()
	time.Sleep(1 * time.Second)
	val := *num
	time.Sleep(1 * time.Second)
	locker.RUnlock()
	return val
}

func syncWrite(locker *sync.RWMutex, num *int, value int) int {
	locker.Lock()
	*num = value
	time.Sleep(1 * time.Second)
	locker.Unlock()
	return *num
}

func useLocker() {
	num := 5
	// 读写互斥锁
	locker := &sync.RWMutex{}
	go func() {
		val := syncWrite(locker, &num, 10)
		fmt.Printf("lock write: value = %d\n", val)
	}()
	go func() {
		val := syncRead(locker, &num)
		fmt.Printf("lock read: value = %d\n", val)
	}()
	go func() {
		val := syncRead(locker, &num)
		fmt.Printf("lock read: value = %d\n", val)
	}()
	time.Sleep(5 * time.Second)
}

func useCond() {
	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)
	cond.L.Lock()

	go func() {
		time.Sleep(2 * time.Second)
		cond.Signal()
		fmt.Println("cond: signal")
	}()

	cond.Wait()
	fmt.Println("cond: after wait...")
}

func SampleMain()  {

	fmt.Println("\n[sync_sample]")
	useOnce()
	useWaitGroup()
	useMap()
	useLocker()
	useCond()
}