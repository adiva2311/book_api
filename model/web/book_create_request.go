package web

type BookCreateRequest struct {
	Title    string `validate:"required" json:"title"`
	Category string `validate:"required" json:"category"`
}