package services

import (
	"book_api/helper"
	"book_api/model/domain"
	"book_api/model/web"
	"book_api/repository"
	"context"
	"database/sql"
)

type BookServiceImpl struct {
	BookRepo repository.BookRepository
	DB *sql.DB
}

func (service *BookServiceImpl) Create(ctx context.Context, request web.BookCreateRequest) (_ web.BookResponse) {
	tx, err := service.DB.Begin()
	if err != nil{
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	book := domain.Book{
		Title: request.Title,
		Category: request.Category,
	}

	book = service.BookRepo.Create(ctx, tx, book)

	return helper.ToBookResponse(book)
	
}

func (service *BookServiceImpl) Update(ctx context.Context, request web.BookUpdateRequest) (_ web.BookResponse) {
	tx, err := service.DB.Begin()
	if err != nil{
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	//Mencari apakah ID ada
	book, err := service.BookRepo.FindById(ctx, tx, request.Id)
	if err != nil{
		panic(err)
	}

	book.Title = request.Title
	book.Category = request.Category

	book = service.BookRepo.Update(ctx, tx, book)

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) Delete(ctx context.Context, bookId int) {
	tx, err := service.DB.Begin()
	if err != nil{
		panic(err)
	}
	defer helper.CommitOrRollback(tx)
	
	//Mencari apakah ID ada
	book, err := service.BookRepo.FindById(ctx, tx, bookId)
	if err != nil{
		panic(err)
	}

	service.BookRepo.Delete(ctx, tx, book)
}

func (service *BookServiceImpl) FindById(ctx context.Context, bookId int) (_ web.BookResponse) {
	tx, err := service.DB.Begin()
	if err != nil{
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepo.FindById(ctx, tx, bookId)
	if err != nil{
		panic(err)
	}

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) FindAll(ctx context.Context) (_ []web.BookResponse) {
	tx, err := service.DB.Begin()
	if err != nil{
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	books := service.BookRepo.FindAll(ctx, tx)

	var bookResponses []web.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, helper.ToBookResponse(book))
	}

	return bookResponses
}

