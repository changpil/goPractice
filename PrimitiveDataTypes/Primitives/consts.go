package main

import "fmt"

func main() {

	const pi = 3.1425
	fmt.Println(pi)
	// pi = 1.2  Eror no reassigning to a const value


	const c = 3  // Const is omplicitly typed const 
	fmt.Println(c + 3) // c typed with int
	fmt.Println(c + 3.3) // c typed with string

	const cc int = 3
	fmt.Println(cc + 3) // c typed with int
	fmt.Println(float32(cc) + 3.3) // c typed with string

}
