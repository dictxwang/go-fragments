package runtime

import (
	"fmt"
	"runtime"
)

func SampleMain()  {

	fmt.Println("\n[runtime_sample]")
	fmt.Printf("numcpu=%d\n", runtime.NumCPU())
	fmt.Printf("goroot=%s\n", runtime.GOROOT())
	fmt.Printf("goos=%s\n", runtime.GOOS)
	fmt.Printf("numgoroutine=%d\n", runtime.NumGoroutine())
}
