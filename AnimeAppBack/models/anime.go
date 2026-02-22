package models

type Anime struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Episodes int    `json:"episodes"`
	Image    string `json:"image"`
}
