package main

import "fmt"

func main() {
	m := map[string]int{"chang": 14, "Albert": 67, "Celi": 34}

	fmt.Println(m)
	fmt.Println(m["chang"])

	delete(m, "Albert")
	fmt.Println(m)

	var m1 map[string]int
	fmt.Println(m1)
	// m1["sss"] = 23 panic: assignment to entry in nil map

	m2 := make(map[string]int)
	m2["sss"] = 23
	fmt.Println(m2)

	m3 := map[string]int{}
	m3["sss"] = 23
	fmt.Println(m3)

	// hits := make(map[string]map[string]int)
	// hits["foo"] =

}
