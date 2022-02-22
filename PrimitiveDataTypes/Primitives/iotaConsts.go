package main

import "fmt"

const pi = 3.1415

// const (
// 	first  = 1
// 	second = "second"
// )

const (
	first  = iota
	second = iota
)

const (
	third  = iota + 3
	fourth = 2 << iota
)

const ( //iota reset after new const block
	fifth = iota
	six = fifth + 1
)

func main() {
	fmt.Println(pi)
	fmt.Println(first, second)
	fmt.Println(third, fourth)
	fmt.Println(third, fourth, fifth, six)
}
