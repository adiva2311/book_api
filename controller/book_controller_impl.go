package controller

import (
	"book_api/model/web"
	"book_api/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type BookControllerImpl struct {
	BookService services.BookService
}

func NewBookController(bookService services.BookService) BookController {
	return &BookControllerImpl{
		BookService: bookService,
	}
}

func (controller *BookControllerImpl) Create(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	bookCreateRequest := web.BookCreateRequest{}
	err := decoder.Decode(&bookCreateRequest)
	if err!=nil {
		panic(err)
	}

	bookResponse := controller.BookService.Create(request.Context(), bookCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "Success Create Data",
		Data: bookResponse,
	}

	writer.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	if err != nil{
		panic(err)
	}
}

func (controller *BookControllerImpl) Update(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	decoder := json.NewDecoder(request.Body)
	bookUpdateRequest := web.BookUpdateRequest{}
	err := decoder.Decode(&bookUpdateRequest)
	if err!=nil {
		panic(err)
	}

	bookId := param.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	if err != nil{
		panic(err)
	}
	bookUpdateRequest.Id = id 

	bookResponse := controller.BookService.Update(request.Context(), bookUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "Success Update Data",
		Data: bookResponse,
	}

	writer.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	if err != nil{
		panic(err)
	}
}

func (controller *BookControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	bookId := param.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	if err != nil{
		panic(err)
	}

	controller.BookService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "Success Delete Data",
	}

	writer.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	if err != nil{
		panic(err)
	}
}

func (controller *BookControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	bookId := param.ByName("bookId")
	id, err := strconv.Atoi(bookId)
	if err != nil{
		panic(err)
	}

	bookResponse, err := controller.BookService.FindById(request.Context(), id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusNotFound)
		return
	}
	webResponse := web.WebResponse{
		Code: 200,
		Status: "Success Get Data by Id",
		Data: bookResponse,
	}

	writer.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writer)
	err = encoder.Encode(webResponse)
	if err != nil{
		panic(err)
	}
}

func (controller *BookControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, param httprouter.Params) {
	bookResponses := controller.BookService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "Success Get All Data",
		Data: bookResponses,
	}

	writer.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writer)
	err := encoder.Encode(webResponse)
	if err != nil{
		panic(err)
	}
}
