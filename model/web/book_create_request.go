package web

type BookCreateRequest struct {
	Title    string `validate:"required"`
	Category string `validate:"required"`
}