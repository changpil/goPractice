package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	outString := "This is a string"
	stringCollection := strings.Split(outString, " ")
	for _, v := range stringCollection {
		fmt.Println(v)
	}

	outString = "This is a string!|Tjis is another one| I like this"
	stringCollection = strings.Split(outString, "|")
	for _, v := range stringCollection {
		fmt.Println(v)
	}

	outString = "This is a string!|Tjis is another one| I like this"
	stringCollection = strings.SplitAfter(outString, "|")
	for _, v := range stringCollection {
		fmt.Println(v)
	}

	outString = "This is a string!|Tjis is another one| I like this"
	stringCollection = strings.SplitAfterN(outString, "|", 2)
	for _, v := range stringCollection {
		fmt.Println(v)
	}

	file, _ := os.Open("customerlist.csv")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		fmt.Println("--- New record ---")
		for _, v := range items {
			fmt.Println(v)
		}
		fmt.Println()
	}
}
