package main

import "fmt"

type HTTPRequest struct {
	Method string
}

func ret(r HTTPRequest) {
	switch r.Method {
	case "GET":
		fmt.Println("GET request")
		fallthrough
	case "DELETE":
		fmt.Println("DELETE request")
		//fallthrough
	case "POST":
		fmt.Println("POST request")
		//fallthrough
	case "PUT":
		fmt.Println("PUT request")
	default:
		fmt.Println("Unhandled case")
	}
}
func main() {
	r := []HTTPRequest{}
	r = append(r, HTTPRequest{Method: "GET"})
	r = append(r, HTTPRequest{Method: "PUT"})
	r = append(r, HTTPRequest{Method: "HEAD"})

	for _, re := range r {
		ret(re)
	}

}
