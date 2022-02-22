package main

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	u1 := User{
		ID:        1,
		FirstName: "Chang",
		LastName:  "Lee",
	}
	u2 := User{
		ID:        2,
		FirstName: "Changpil",
		LastName:  "Lee",
	}

	if u1.ID == u2.ID {
		fmt.Println("Same user")
	}

	if u1.ID != u2.ID {
		fmt.Println("Not Same user")
	}

	if u1.ID == u2.ID {
		fmt.Println("Not Same user")
	} else {
		fmt.Println("Not Same user")
	}

	if u1 == u2 {
		fmt.Println("Same user")
	} else if u1.LastName == u2.LastName {
		fmt.Println("Same family")
	} else {
		fmt.Println("differnt users")
	}

	u3 := User{
		ID:        1,
		FirstName: "Chang",
		LastName:  "Lee",
	}
	//Struct Value comparison
	if u1 == u3 {
		fmt.Println("Same user")
	} else if u1.LastName == u3.LastName {
		fmt.Println("Same family")
	} else {
		fmt.Println("differnt users")
	}
}
