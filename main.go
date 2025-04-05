package main

import (
	"book_api/app"
	"book_api/controller"
	"book_api/repository"
	"book_api/services"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDb()
	validate := validator.New()
	
	bookRepository := repository.NewBookControllerImpl()
	bookService := services.NewBookService(bookRepository, db, validate)
	bookController := controller.NewBookController(bookService)

	router := httprouter.New()

	router.GET("/api/books", bookController.FindAll)
	router.GET("/api/books/:bookId", bookController.FindById)
	router.POST("/api/books", bookController.Create)
	router.PUT("/api/books/:bookId", bookController.Update)
	router.DELETE("/api/books/:bookId", bookController.Delete)

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}