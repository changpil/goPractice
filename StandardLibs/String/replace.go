package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "This is my string. There are many strings like it, but this one is mine"
	sampleString = strings.Replace(sampleString, "string", "compliler", -1)

	fmt.Println(sampleString)

	sampleString = "This is my string. There are many strings like it, but this one is mine"
	sampleString = strings.Replace(sampleString, "string", "compliler", 1)

	fmt.Println(sampleString)
}
