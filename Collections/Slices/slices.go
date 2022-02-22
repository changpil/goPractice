package main

import "fmt"

func main() {
	arr := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	slice := arr[:] //Pointer like slice
	arr[0] = 1
	fmt.Println(arr, slice)

	slice = append(slice, 5)
	slice = append(slice, 5)
	slice = append(slice, 5)
	// arr = append(arr, 5) No appending for array: Not Possible
	fmt.Println(arr, slice)

	slice1 := []int{0, 1, 2, 3, 4, 5}
	slice1 = append(slice1, 6, 7, 8, 9, 10)
	slice2 := slice1[1:]
	slice3 := slice1[:4]
	slice4 := slice1[2:6]
	fmt.Println(slice1, slice2, slice3, slice4)
}
