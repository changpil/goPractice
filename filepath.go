package main

import (
    "fmt"
    "path/filepath"
    "os"
)

func main() {

        f, _ := filepath.Abs("~/.helm")
	fmt.Println(f)
	fmt.Println(os.Getenv("HOME"))

}

