package main

import "fmt"

// Array, slice,  map
func main() {
	slice := []int{1, 2, 3}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, n := range slice {
		fmt.Println(i, n)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}

	for k, v := range wellKnownPorts {
		fmt.Println(k, v)
	}
}
