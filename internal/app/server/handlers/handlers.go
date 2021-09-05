package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"library-server/internal/app/server/models"
	"library-server/internal/app/server/repositories/bookrepository"
	"library-server/pkg/logger"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type HandlerContext struct {
	Collection  *mongo.Collection
	Context     context.Context
	FilterParam string
	W           http.ResponseWriter
	R           *http.Request
}

var (
	response []byte
	err      error
)
var log = logger.NewLogger()

// AddBookHandler inserts one book into the table `books`
func AddBookHandler(ctx HandlerContext) {
	var book models.Book
	if err := json.NewDecoder(ctx.R.Body).Decode(&book); err != nil {
		logger.RaiseAlert(logger.LoggerContext{W: ctx.W, Message: err.Error(),Status: http.StatusInternalServerError})
		return
	}
	if response, err = bookrepository.InsertOneBook(bookrepository.BookRepositoryContext{Collection: ctx.Collection, W: ctx.W, Context: nil, Book: book}); err != nil {
		return
	}
	fmt.Fprintf(ctx.W, "Book record sucessully injected : %s", response)
}

// WelcomeHandler
func WelcomeHandler(ctx HandlerContext) {
	ctx.W.WriteHeader(http.StatusOK)
	fmt.Fprintf(ctx.W, "Welcome handler is working fine")
}

// GetBooksHandler gets all the books and returns paginated response
func GetBooksHandler(ctx HandlerContext) {
	ctx.W.Header().Set("Content-Type", "application/json")
	var books []models.Book
	if books, err = bookrepository.GetAllBooks(bookrepository.BookRepositoryContext{Collection: ctx.Collection, W: ctx.W, Context: ctx.Context}); err != nil {
		return
	}
	json.NewEncoder(ctx.W).Encode(books)
}

// GetBookByBookIdHandler gets first book based on bookid
func GetBookByBookIdHandler(ctx HandlerContext) {
	var book models.Book
	ctx.W.Header().Set("Content-Type", "application/json")
	if book, err = bookrepository.GetBookByBookId(bookrepository.BookRepositoryContext{Collection: ctx.Collection, W: ctx.W, Context: ctx.Context, Param: ctx.FilterParam, ParamValue: ctx.R.URL.Query().Get("bookid")}); err != nil {
		return
	}
	json.NewEncoder(ctx.W).Encode(book)
}

// GetBookByBookNameHandler gets first book based on bookname
func GetBookByBookNameHandler(ctx HandlerContext) {
	var book models.Book
	ctx.W.Header().Set("Content-Type", "application/json")
	if book, err = bookrepository.GetBookByBookName(bookrepository.BookRepositoryContext{Collection: ctx.Collection, W: ctx.W, Context: ctx.Context, Param: ctx.FilterParam, ParamValue: ctx.R.URL.Query().Get("bookname")}); err != nil {
		return
	}
	json.NewEncoder(ctx.W).Encode(book)
}

// GetBookByBookAuthorHandler gets first book based on book author name
func GetBookByBookAuthorNameHandler(ctx HandlerContext) {
	var book models.Book
	ctx.W.Header().Set("Content-Type", "application/json")
	if book, err = bookrepository.GetBookByBookAuthorName(bookrepository.BookRepositoryContext{Collection: ctx.Collection, W: ctx.W, Context: ctx.Context, Param: ctx.FilterParam, ParamValue: ctx.R.URL.Query().Get("bookauthor")}); err != nil {
		return
	}
	json.NewEncoder(ctx.W).Encode(book)
}

// GetBookByIsbnHandler gets first book based on isbn
func GetBookByIsbnHandler(ctx HandlerContext) {
	var book models.Book
	ctx.W.Header().Set("Content-Type", "application/json")
	if book, err = bookrepository.GetBookByIsbn(bookrepository.BookRepositoryContext{Collection: ctx.Collection, W: ctx.W, Context: ctx.Context,Param:  ctx.FilterParam, ParamValue: ctx.R.URL.Query().Get("isbn")}); err != nil {
		return
	}
	json.NewEncoder(ctx.W).Encode(book)
}

// GetBookByPriceHandler get first book based on price
func GetBookByPriceHandler(ctx HandlerContext) {
	var book models.Book
	ctx.W.Header().Set("Content-Type", "application/json")
	if book, err = bookrepository.GetBookByPrice(bookrepository.BookRepositoryContext{Collection: ctx.Collection, W: ctx.W, Context: ctx.Context, Param: ctx.FilterParam, ParamValue: ctx.R.URL.Query().Get("price")}); err != nil {
		return
	}
	json.NewEncoder(ctx.W).Encode(book)
}

func AddBooksHandler(ctx HandlerContext) {
	var booksAsSlice []models.Book
	if err = json.NewDecoder(ctx.R.Body).Decode(&booksAsSlice); err != nil {
		logger.RaiseAlert(logger.LoggerContext{W: ctx.W, Message: err.Error(),Status: http.StatusInternalServerError})
		return
	}
	var booksAsInterfaceSlice []interface{} = make([]interface{}, len(booksAsSlice))
	for i, d := range booksAsSlice {
		booksAsInterfaceSlice[i] = d
	}
	if _, err = bookrepository.InsertMultipleBook(bookrepository.BookRepositoryContext{Collection: ctx.Collection, W: ctx.W, Context: ctx.Context, Books: booksAsInterfaceSlice}); err != nil {
		return
	}
	fmt.Fprintf(ctx.W, "Book records sucessully injected : %d", len(booksAsSlice))
}
