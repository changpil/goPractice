package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	s := "/users/344/"
	re := regexp.MustCompile(`^/users/(\d+)/?`)
	match := re.FindStringSubmatch(s)
	fmt.Println(match)
	n, err := strconv.Atoi(match[1])
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Println(s, n)

}
