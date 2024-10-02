package id

import (
	"fmt"
	"github.com/google/uuid"
)

func UuidMain() {

	ustr := uuid.NewString()
	fmt.Printf("uuid string: %s\n", ustr)
}
