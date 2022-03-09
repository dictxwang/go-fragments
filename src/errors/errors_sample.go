package _errors

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func SampleMain()  {
	fmt.Println("\n[errors_sample]")

	_, err := os.OpenFile("F:/xx/xx.log", os.O_RDWR, 0)
	if err != nil {
		// errors包在众多公共包中被使用，几乎是行业标准了
		err := errors.Wrap(err, "wrap error")
		fmt.Println(err)
	}

	//switch err := errors.Cause(err).(type) {
	//case *MyError:
	//	// handle specifically
	//default:
	//	// unknown error
	//}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic recover: %s\n", r)
		}
	}()

	a := 0
	b := 10
	c := b / a
	fmt.Println("after error, ", c)
}
