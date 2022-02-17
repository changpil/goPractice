package main

import (
	"fmt"
)

type AA struct {
    A  string
    B  *string
}

func main() {
   test := new(AA)
   test.A = "aaaaa"
   test.B = &"bbbb"
	
   fmt.Println(test.A)
   fmt.Println(test.B)

}
