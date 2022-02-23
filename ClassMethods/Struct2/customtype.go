package main

import (
	"fmt"
	"media/media"
)

func main() {

	myFav := media.NewMovie("Top Gun", media.PG, 32.8)

	fmt.Printf("My favorite movie is %s\n", myFav.GetTitle())
	fmt.Printf("It was rated %s\n", myFav.GetRating())
	fmt.Printf("It made %f in the box office\n", myFav.GetBoxOffice())

	myFav.SetTitle("Dumb and Dumber")

	fmt.Printf("My favorite movie now is %s\n", myFav.GetTitle())

}
