package media

import "strings"

type Movie struct {
	title     string
	rating    Rating
	boxOffice float32
}

type Rating string

const (
	R    = "R (Restricted)"
	G    = "G (General audiences)"
	PG   = "PG (Parental Guidance)"
	PG13 = "PG-13 (Parental Caution)"
	NC17 = "NC-17 (No children under 17)"
)

func NewMovie(title string, rating Rating, boxOffice float32) Movie {
	return Movie{
		title:     title,
		rating:    rating,
		boxOffice: boxOffice,
	}
}

func (m Movie) GetTitle() string {
	return strings.ToTitle(m.title)
}

func (m Movie) GetRating() string {
	return string(m.rating)
}

func (m Movie) GetBoxOffice() float32 {
	return m.boxOffice
}

func (m Movie) SetTitle(newTitle string) {
	m.title = newTitle
}

func (m Movie) SetRating(newRating Rating) {
	m.rating = newRating
}

func (m Movie) SetBoxOffice(newBoxOffice float32) {
	m.boxOffice = newBoxOffice
}
