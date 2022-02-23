package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "This is Chang"
	s2 := "This is chang"

	if s1 == s2 {
		fmt.Println("Two strings are same")
	} else {
		fmt.Println("Two strings are not same")
	}

	if strings.Compare(s1, s2) == 0 {
		fmt.Println("Two strings are same")
	} else {
		fmt.Println("Two strings are not same")
	}

	if CompareCaseIns(s1, s2) {
		fmt.Println("Two strings are same")
	} else {
		fmt.Println("Two strings are not ame")
	}
}

func CompareCaseIns(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	if strings.ToLower(a) == strings.ToLower(b) {
		return true
	}
	return false
}
