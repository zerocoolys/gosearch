package request

import (
	"fmt"
)

type Request struct {
	indices []string
	types   []string
}

func Say() {
	fmt.Println("Say Hello")
}
