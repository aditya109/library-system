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
}

var (
	response []byte
	err      error
)
var log = logger.NewLogger()

// AddBookHandler inserts one book into the table `books`
func AddBookHandler(w http.ResponseWriter, r *http.Request, ctx HandlerContext) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		logger.RaiseAlert(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if response, err = bookrepository.InsertOneBook(bookrepository.BookRepositoryContext{Collection: ctx.Collection, W: w, Context: nil}, book); err != nil {
		return
	}
	fmt.Fprintf(w, "Book record sucessully injected : %s", response)
}

// WelcomeHandler
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome handler is working fine")
}

// GetBooksHandler gets all the books and returns paginated response
func GetBooksHandler(w http.ResponseWriter, r *http.Request, ctx HandlerContext) {
	w.Header().Set("Content-Type", "application/json")
	var books []models.Book
	if books, err = bookrepository.GetAllBooks(bookrepository.BookRepositoryContext{Collection: ctx.Collection, W: w, Context: ctx.Context}); err != nil {
		return
	}
	json.NewEncoder(w).Encode(books)
}

// GetBookByBookIdHandler gets first book based on bookid
func GetBookByBookIdHandler(w http.ResponseWriter, r *http.Request, ctx HandlerContext) {
	var book models.Book
	w.Header().Set("Content-Type", "application/json")
	if book, err = bookrepository.GetBookByBookId(bookrepository.BookRepositoryContext{Collection: ctx.Collection, W: w, Context: ctx.Context}, ctx.FilterParam, r.URL.Query().Get("bookid")); err != nil {
		return
	}
	json.NewEncoder(w).Encode(book)
}

// GetBookByBookNameHandler gets first book based on bookname
func GetBookByBookNameHandler(w http.ResponseWriter, r *http.Request, ctx HandlerContext) {
	var book models.Book
	w.Header().Set("Content-Type", "application/json")
	if book, err = bookrepository.GetBookByBookName(bookrepository.BookRepositoryContext{Collection: ctx.Collection, W: w, Context: ctx.Context}, ctx.FilterParam, r.URL.Query().Get("bookname")); err != nil {
		return
	}
	json.NewEncoder(w).Encode(book)
	log.DatabaseEvent(fmt.Sprintf("Fetch successful, #books ID: %d", book.BookId))
}

// GetBookByBookAuthorHandler gets first book based on book author name
func GetBookByBookAuthorNameHandler(w http.ResponseWriter, r *http.Request, ctx HandlerContext) {
	var book models.Book
	w.Header().Set("Content-Type", "application/json")
	if book, err = bookrepository.GetBookByBookAuthorName(bookrepository.BookRepositoryContext{Collection: ctx.Collection, W: w, Context: ctx.Context}, ctx.FilterParam, r.URL.Query().Get("bookauthor")); err != nil {
		return
	}
	json.NewEncoder(w).Encode(book)
	log.DatabaseEvent(fmt.Sprintf("Fetch successful, #books ID: %d", book.BookId))
}

// GetBookByIsbnHandler gets first book based on isbn
func GetBookByIsbnHandler(w http.ResponseWriter, r *http.Request, ctx HandlerContext) {
	var book models.Book
	w.Header().Set("Content-Type", "application/json")
	if book, err = bookrepository.GetBookByIsbn(bookrepository.BookRepositoryContext{Collection: ctx.Collection, W: w, Context: ctx.Context}, ctx.FilterParam, r.URL.Query().Get("isbn")); err != nil {
		return
	}
	json.NewEncoder(w).Encode(book)
	log.DatabaseEvent(fmt.Sprintf("Fetch successful, #books ID: %d", book.BookId))
}

// GetBookByPriceHandler get first book based on price
func GetBookByPriceHandler(w http.ResponseWriter, r *http.Request, ctx HandlerContext) {
	var book models.Book
	w.Header().Set("Content-Type", "application/json")
	if book, err = bookrepository.GetBookByPrice(bookrepository.BookRepositoryContext{Collection: ctx.Collection, W: w, Context: ctx.Context}, ctx.FilterParam, r.URL.Query().Get("price")); err != nil {
		return
	}
	json.NewEncoder(w).Encode(book)
	log.DatabaseEvent(fmt.Sprintf("Fetch successful, #books ID: %d", book.BookId))
}

func AddBooksHandler(w http.ResponseWriter, r *http.Request, ctx HandlerContext) {
	var booksAsSlice []models.Book
	if err = json.NewDecoder(r.Body).Decode(&booksAsSlice); err != nil {
		logger.RaiseAlert(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var booksAsInterfaceSlice []interface{} = make([]interface{}, len(booksAsSlice))
	for i, d := range booksAsSlice {
		booksAsInterfaceSlice[i] = d
	}
	if _ ,err = bookrepository.InsertMultipleBook(bookrepository.BookRepositoryContext{Collection: ctx.Collection, W: w, Context: ctx.Context}, booksAsInterfaceSlice); err != nil {
		return 
	}
	fmt.Fprintf(w, "Book records sucessully injected : %d", len(booksAsSlice))
}
