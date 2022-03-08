package context

import (
	"context"
	"fmt"
	"time"
)

func doContextHandle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func contextHandle() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// 当执行时间小于上面的超时时间，方法正常执行
	// 否则收到信息：handle context deadline exceeded
	go doContextHandle(ctx, 500 * time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main:", ctx.Err())

	}
}

func contextCancel() {
	ctx, cancel := context.WithCancel(context.Background())
	// 同时传递value
	ctx = context.WithValue(ctx, "v001", "123")

	go func() {
		for {
			val := ctx.Value("v001").(string)
			fmt.Printf("contextCancel: not cancelled, value=%s\n", val)
			select {
			case <-ctx.Done():
				fmt.Println("contextCancel: has done...")
				return
			default:
				fmt.Println("contextCancel: not done")
			}
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func SampleMain()  {

	fmt.Println("\n[context_sample]")
	contextHandle()
	contextCancel()
}
