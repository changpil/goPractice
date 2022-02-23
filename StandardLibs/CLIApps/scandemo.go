package main

import (
	"fmt"
)

func main() {

	// var name string
	name := ""
	fmt.Println("What is your name?")
	test, _ := fmt.Scanf("%s", &name)
	// test, _ := fmt.Scanf("%q", &name) //"You need put quoted string like "chang pil"

	switch test {
	case 0:
		fmt.Println("You must enter a name!")
	case 1:
		fmt.Printf("Hello %s! \n", name)
	}
}
