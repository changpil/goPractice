package main

import (
	"fmt"
	"strings"
)

func main() {
	sampleString := "      I really enjoy the     "
	fmt.Printf("%q\n", sampleString)
	trimedString := strings.TrimSpace(sampleString)
	fmt.Printf("%q\n", trimedString)

	trimedString = strings.TrimLeft(sampleString, " ")
	fmt.Printf("%q\n", trimedString)

	trimedString = strings.TrimRight(sampleString, " ")
	fmt.Printf("%q\n", trimedString)

	url := "https://www.changpil.com"
	domainName := strings.TrimPrefix(url, "https://")
	fmt.Printf("%q\n", domainName)

	filepath := "/home/chang/docs/sample.txt"
	path := strings.TrimSuffix(filepath, ".txt")
	fmt.Printf("%q\n", path)
}
