package main

import "fmt"
// func make([]type, length, capacity) []type
func main() {
    // Creating another array of size 6
    // and return the reference of the slice
    // Using make() function
    var sliceB = make([]int, 6)
    fmt.Printf("SliceB = %v, \nlength = %d, \ncapacity = %d\n",
        sliceB, len(sliceB), cap(sliceB))


	var sliceC = make(map[string]int)
	sliceC["a"] = 23
	sliceC["b"] = 24
	sliceC["c"] = 25
	sliceC["d"] = 26
	sliceC["e"] = 27
	fmt.Println(sliceC)
}