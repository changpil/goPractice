package main

import (
	"fmt"
	"regexp"
)

func main() {
	sampleString1 := "I really enjoy the Go language. It is one of my favorite"
	sampleString2 := "This is my song. There are many strings like it, but this one is mine"
	// r, _ := regexp.Compile(`s([a-z]+)g`)
	r, _ := regexp.Compile(`\w([a-z]+)g\b`)

	result := r.MatchString(sampleString1)
	if result {
		fmt.Printf("%v matches with %s", r, sampleString1)
	}
	fmt.Println()
	result = r.MatchString(sampleString2)
	if result {
		fmt.Printf("%v matches with %s", r, sampleString2)
	}
	fmt.Println()

	fmt.Println(r.FindAllString(sampleString2, -1))
	fmt.Println(r.FindStringIndex(sampleString2))
	newText := r.ReplaceAllString(sampleString2, "apple")
	fmt.Println(newText)
}
