package main

import (
	"fmt"
	"media/media"
)

func main() {
	fmt.Println("My Favorite Movie")
	m2 := media.Movie2{}
	m2.Title = "Top Gun"
	m2.Rating = media.R
	m2.BoxOffice = 43.2

	fmt.Printf("My favorite movie is %s\n", m2.Title)
	fmt.Printf("It was rated %s\n", m2.Rating)
	fmt.Printf("It made %f in the box office\n", m2.BoxOffice)

	// Not Working
	// m := media.Movie{
	// 	title:     "Top Gun",
	// 	rating:    media.R,
	// 	boxOffice: 43.2,
	// }


	var myFav media.Catalogable = &media.Movie{}

	myFav.NewMovie("Top Gun", media.PG, 32.8)

	fmt.Printf("My favorite movie is %s\n", myFav.GetTitle())
	fmt.Printf("It was rated %s\n", myFav.GetRating())
	fmt.Printf("It made %f in the box office\n", myFav.GetBoxOffice())

	myFav.SetTitle("Dumb and Dumber")

	fmt.Printf("My favorite movie now is %s\n", myFav.GetTitle())

}
