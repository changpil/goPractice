package main

import (
	"flag"
	"fmt"
)

func main() {

	arch := flag.String("arch", "x86", "CPU Architecture")
	os := flag.String("os", "Mac", "OS Type")
	flag.Parse()

	switch *arch {

	case "x86":
		{
			fmt.Println("running in 32 bit mode")
		}
	case "AMD64":
		{
			fmt.Println("Running in 64 bit mode")
		}
	case "IA64":
		{
			fmt.Println("Remember IA64?")
		}
	}

	switch *os {

	case "Mac":
		{
			fmt.Println("running in Mac")
		}
	case "Windows":
		{
			fmt.Println("Running in Wondows")
		}
	case "Linux":
		{
			fmt.Println("Running in Linux")
		}
	}

	fmt.Println("Ran our process!")
}
