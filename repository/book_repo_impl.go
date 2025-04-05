package repository

import (
	"book_api/model/domain"
	"context"
	"database/sql"
	"errors"
)

type BookRepositoryImpl struct {
	
}

func (repo *BookRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, book domain.Book) (_ domain.Book) {
	query := "INSERT INTO book (title, category) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, query, book.Title, book.Category)
	if err != nil{
		panic(err)
	}

	id, err := result.LastInsertId()
	if err != nil{
		panic(err)
	}

	book.Id = int(id)
	return book
}

func (repo *BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book domain.Book) (_ domain.Book) {
	query := "UPDATE book SET title = ?, category = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, book.Title, book.Category, book.Id)
	if err != nil{
		panic(err)
	}

	return book
}

func (repo *BookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, book domain.Book) {
	query := "DELETE FROM book WHERE id = ?"
	_, err := tx.ExecContext(ctx, query, book.Id)
	if err != nil{
		panic(err)
	}
}

func (repo *BookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookId int) (domain.Book, error) {
	query := "SELECT FROM book WHERE id = ?"
	rows, err := tx.QueryContext(ctx, query, bookId)
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	book := domain.Book{}
	if rows.Next(){
		err := rows.Scan(&book.Id, &book.Title, &book.Category)
		if err != nil{
			panic(err)
		}

		return book, nil
	} else {
		return book, errors.New("book not found")
	}
}

func (repo *BookRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) (_ []domain.Book) {
	query := "SELECT id, title, category FROM book"
	rows, err := tx.QueryContext(ctx, query)
	if err != nil{
		panic(err)
	}	
	defer rows.Close()

	var books []domain.Book
	if rows.Next(){
		book := domain.Book{}
		err := rows.Scan(&book.Id, &book.Title, &book.Category)
		if err != nil{
			panic(err)
		}
		books = append(books, book)
	}
	return books
}

