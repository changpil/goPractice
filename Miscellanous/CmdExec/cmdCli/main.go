package main

// import (
// 	"strings"
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"log"
// )


import (
	"flag"
	"strings"
	"bufio"
	"fmt"
	"os"
	"log"
)
// func main() {
// 	f, err := os.Open("myapp.log")
// 	if err != nil {
// 		log.Fatal(err) 
//     }
//     defer f.Close()
// 	r := bufio.NewReader(f)
// 	for {
// 		s, err := r.ReadString('\n')
// 		if err != nil {
// 			break
// 		}
// 		if strings.Contains(s, "ERROR") {
// 			fmt.Println(s) 
// 		}
// 	}
// }
func main() {
	path := flag.String("path", "myapp.log", "The path to the log that should be analyzed")
	level := flag.String("level", "ERROR", "Log level to search for. Options are DEBUG, INFO, ERROR, and CRITICAL")

	flag.Parse()

	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err) 
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		s, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(s, *level) {
			fmt.Println(s) 
		}
	}
}
