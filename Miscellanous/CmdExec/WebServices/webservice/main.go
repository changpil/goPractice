package main

import "net/http"
import "log"

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		w.Write([]byte("Hello World"))
// 	})

// 	err := http.ListenAndServe(":3000", nil)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		names := r.URL.Query()["name"]
		var name string
		if len(names) == 1 {
			name = names[0]
		}
		w.Write([]byte("Hello, " + name))
	})

	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}