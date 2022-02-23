package main

import "fmt"

func main() {
	// String is a  slice of bytes
	outString := "\x47\x6f\x20\x69\x73\x20\x41\x41\x77\x65\x73\x6f\x6d\x65\x21"
	fmt.Println(outString)
	for _, c := range outString {
		fmt.Printf("%x ", c)
		fmt.Println()
	}
	for _, c := range outString {
		fmt.Printf("%q", c)
		fmt.Println()
	}

	outString = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(outString)

	for i, _ := range outString {
		fmt.Printf("%q\n", outString[i])
	}

	for _, v := range outString {
		fmt.Printf("%q\n", v)
	}
}
