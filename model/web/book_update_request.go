package web

type BookUpdateRequest struct {
	Id       int    `validate:"required"`
	Title    string `validate:"required"`
	Category string `validate:"required"`
}