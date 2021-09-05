package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"library-server/internal/app/server/models"
	"library-server/internal/app/server/repositories/bookrepository"
	"library-server/pkg/logger"
	"net/http"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		logger.RaiseAlert(w, log, err.Error(), http.StatusInternalServerError)
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

func GetBookHandler(w http.ResponseWriter, r *http.Request, ctx HandlerContext) {
	var collection *mongo.Collection
	var log *logger.StandardLogger
	var err error
	var book models.Book
	var filter bson.M
	collection = ctx.Collection
	log = logger.NewLogger()
	filterParam := ctx.FilterParam
	w.Header().Set("Content-Type", "application/json")
	switch filterParam { // filters
	case "bookid":
		bookId, err := strconv.Atoi(r.URL.Query().Get("bookid"))
		if err != nil {
			logger.RaiseAlert(w, log, fmt.Sprintf("error while casting bookid: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		filter = bson.M{filterParam: bookId}
	case "bookname":
		bookName := r.URL.Query().Get("bookname")
		filter = bson.M{
			filterParam: bson.M{
				"$regex": primitive.Regex{
					Pattern: bookName,
					Options: "i",
				},
			},
		}
	case "price":
		price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 64)
		if err != nil {
			logger.RaiseAlert(w, log, fmt.Sprintf("error while casting price: %s", err.Error()), http.StatusInternalServerError)
			return
		}
		filter = bson.M{filterParam: price}
	case "isbn":
		isbn := r.URL.Query().Get("isbn")
		filter = bson.M{
			filterParam: bson.M{
				"$regex": primitive.Regex{
					Pattern: isbn,
					Options: "i",
				},
			},
		}
	case "bookauthor":
		bookAuthor := r.URL.Query().Get("bookauthor")
		filter = bson.M{
			filterParam: bson.M{
				"$regex": primitive.Regex{
					Pattern: bookAuthor,
					Options: "i",
				},
			},
		}
	}
	if err = collection.FindOne(ctx.Context, filter).Decode(&book); err != nil {
		logger.RaiseAlert(w, log, fmt.Sprintf("error while fetching data from database: %s", err.Error()), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(book)
	log.DatabaseEvent(fmt.Sprintf("Fetch successful, #books ID: %d", book.BookId))
}

func AddBooksHandler(w http.ResponseWriter, r *http.Request, ctx HandlerContext) {
	collection := ctx.Collection
	log := logger.NewLogger()
	var booksAsSlice []models.Book
	err := json.NewDecoder(r.Body).Decode(&booksAsSlice)
	if err != nil {
		logger.RaiseAlert(w, log, err.Error(), http.StatusInternalServerError)
		return
	}
	var booksAsInterfaceSlice []interface{} = make([]interface{}, len(booksAsSlice))
	for i, d := range booksAsSlice {
		booksAsInterfaceSlice[i] = d
	}

	if _, err := collection.InsertMany(context.Background(), booksAsInterfaceSlice); err != nil {
		logger.RaiseAlert(w, log, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Book records sucessully injected : %d", len(booksAsSlice))
	log.DatabaseEvent(fmt.Sprintf("Insert successful, #Books: %d", len(booksAsSlice)))
}
