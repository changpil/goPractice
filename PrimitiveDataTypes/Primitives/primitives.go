package main

import "fmt"

func main() {
	var i int
	i = 52
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	firstName := "Arthuer" // Implicit variable
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Println(r, im)

}
