package web

type BookUpdateRequest struct {
	Id       int    `validate:"required" json:"id"`
	Title    string `validate:"required" json:"title"`
	Category string `validate:"required" json:"category"`
}