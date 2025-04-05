package domain

type Book struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
}