package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	// cmd := exec.Command("lst")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("SSSSSSSSSSSSSSSS")
		fmt.Printf("%s\n", stdoutStderr)
		fmt.Println("TTTTTTTTTTTTTTTT")

		fmt.Printf("%s\n", err)
		log.Fatal(err)
	}
	fmt.Println("SSSSSSSSSSSSSSSS")
	fmt.Printf("%s\n", stdoutStderr)
}
