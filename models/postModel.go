package models

type Post struct {
	CommonField
	Title string `json:"title"`
	Body  string `json:"body"`
}
