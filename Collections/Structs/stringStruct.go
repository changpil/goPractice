package main

import (
	"fmt"
)

type AA struct {
	A string
	B *string
}

func main() {
	test := new(AA)
	test.A = "aaaaa"
    b := "bbbb"
	test.B = &b

	fmt.Println(test.A)
	fmt.Println(test.B)

}
