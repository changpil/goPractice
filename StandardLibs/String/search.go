package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	sampleString := "I really enjoy the Go language. It is one of my favorite"
	searchTerm := "Go"
	result := false
	// result := strings.Contains(sampleString, searchTerm)
	// fmt.Printf("The sample text include %s: %t\n", searchTerm, result)

	if len(os.Args) > 1 {
		searchTerm = os.Args[1]
		result = strings.Contains(sampleString, searchTerm)
	}
	if result {
		fmt.Printf("The sample text include %s: %t\n", searchTerm, result)
	} else {
		fmt.Printf("The sample text does not include %s: %t\n", searchTerm, result)
	}

	if len(os.Args) > 1 {
		searchTerm = os.Args[1]
		result = strings.HasPrefix(sampleString, searchTerm)
	}
	if result {
		fmt.Printf("The sample text starts with %s: %t\n", searchTerm, result)
	} else {
		fmt.Printf("The sample text does not start with %s: %t\n", searchTerm, result)
	}

	if len(os.Args) > 1 {
		searchTerm = os.Args[1]
		result = strings.HasSuffix(sampleString, searchTerm)
	}
	if result {
		fmt.Printf("The sample text ends with %s: %t\n", searchTerm, result)
	} else {
		fmt.Printf("The sample text does not ends with %s: %t\n", searchTerm, result)
	}

}
