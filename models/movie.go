package models

type Movie struct {
	Id int `json:"id"`
	Title string `json:"title"`
	ReleaseDate string `json:"release_date"`
}
