package services

import (
	"book_api/helper"
	"book_api/model/domain"
	"book_api/model/web"
	"book_api/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type BookServiceImpl struct {
	BookRepo repository.BookRepository
	DB *sql.DB
	Validate *validator.Validate
}

func NewBookService(bookRepo repository.BookRepository, db *sql.DB, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		BookRepo: bookRepo,
		DB: db,
		Validate: validate,
	}
}

func (service *BookServiceImpl) Create(ctx context.Context, request web.BookCreateRequest) (_ web.BookResponse) {
	err := service.Validate.Struct(request)
	if err != nil{
		panic(err)
	}

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
	err := service.Validate.Struct(request)
	if err != nil{
		panic(err)
	}

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

func (service *BookServiceImpl) FindById(ctx context.Context, bookId int) (web.BookResponse, error) {
	tx, err := service.DB.Begin()
	if err != nil{
		panic(err)
	}
	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepo.FindById(ctx, tx, bookId)
	if err != nil{
		panic(err)
	}

	return helper.ToBookResponse(book), nil
}

func (service *BookServiceImpl) FindAll(ctx context.Context) []web.BookResponse {
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

