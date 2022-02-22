package main

import "fmt"

func main() {
	i := 1
	for i < 5 {
		fmt.Println(i)
		i++
	}

	for {
		if i == 10 {
			break
		}
		if i%2 == 0 {
			fmt.Println(i, "continue ...")
			i++
			continue
		}
		fmt.Println(i)
		i++
	}

	for ; i < 20; i++ {
		fmt.Println(i)
	}

	for j := 0; j < 20; j++ {
		fmt.Print(j, " ")
	}
	// fmt.Println(j)    Out of scope
	fmt.Println()

	j := 0
	for j := 0; j < 10; j++ {
		j := 9
		fmt.Println(j)
	}
	fmt.Println(j)

}
