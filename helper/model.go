package helper

import (
	"book_api/model/domain"
	"book_api/model/web"
)

func ToBookResponse(book domain.Book) web.BookResponse{
	return web.BookResponse{
		Id: book.Id,
		Title: book.Title,
		Category: book.Category,
	}
}