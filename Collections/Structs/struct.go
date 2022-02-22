package main

import "fmt"

func main() {
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	fmt.Println(u)
	u.ID = 1234
	u.FirstName = "Chang"
	u.LastName = "Lee"
	fmt.Println(u)

	u2 := user{ID: 12, FirstName: "Chang", LastName: "Lee"}
	fmt.Println(u2)

	u3 := user{ID: 12,
		FirstName: "Chang",
		LastName:  "Lee", // Comma needed: error  unexpected newline
	}
	fmt.Println(u3)
}
