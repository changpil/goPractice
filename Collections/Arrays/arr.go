package main

import "fmt"

func main() {
	var arr [3]int //[size]datatype
	arr[0] = 1
	arr[1] = 2
	arr[2] = 5

	fmt.Println(arr)

	arr1 := [4]int{2, 3, 4}
	fmt.Println(arr1)

	// Not working
	// arr2 := []int{}
	// arr2 := append(arr2, 4)

	// Slice
	s := make([]int, 0)
	s = append(s, 1)
	s = append(s, 2)
	s = append(s, 3)

	fmt.Println(s)

	v := arr1[:]
	v = append(v, 5)
	v = append(v, 9)
	fmt.Println(v)

}
