package main

import "fmt"

func main() {
	// var fname *string
	// fmt.Println(fname)  (nil)

	// fname = "aurther"  cannot use "aurther" (type string) as type *string in assignment
	// *fname = "aurther" Go does not allow to assine value to uninitialized Pointers

	var fname *string = new(string)
	*fname = "Arthur"
	fmt.Println(fname)
	fmt.Println(*fname)

	ptr := fname
	fmt.Println(ptr, *ptr)

	firstName := "Chang"
	ptr = &firstName
	fmt.Println(ptr, *ptr)

	firstName = "Changpil"
	fmt.Println(ptr, *ptr)
}
