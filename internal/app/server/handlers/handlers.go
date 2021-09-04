package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"library-server/internal/app/server/models"
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

func AddBookHandler(w http.ResponseWriter, r *http.Request, ctx HandlerContext) {
	collection := ctx.Collection
	log := logger.NewLogger()
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Issue(err.Error())
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Issue(err.Error())
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	res, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	response, err := json.Marshal(res)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
	fmt.Fprintf(w, "Book record sucessully injected : %s", response)
	fmt.Println(book.BookId)
	log.DatabaseEvent(fmt.Sprintf("Insert successful, BookID: %d", book.BookId))
}

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Welcome handler is working fine")
}

func GetBooksHandler(w http.ResponseWriter, r *http.Request, ctx HandlerContext) {
	collection := ctx.Collection
	log := logger.NewLogger()

	w.Header().Set("Content-Type", "application/json")

	var books []models.Book

	// bson.M{}, we passes empty filter, so we can get all the data
	cursor, err := collection.Find(ctx.Context, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.DatabaseEvent(fmt.Sprintf("error while fetching data from database: %s", err.Error()))
		fmt.Fprintf(w, "error while fetching data from database: %s", err.Error())
		return
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var book models.Book
		err := cursor.Decode(&book)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.DatabaseEvent(fmt.Sprintf("error while reading stream data from database: %s", err.Error()))
			fmt.Fprintf(w, "error while reading stream data from database: %s", err.Error())
			return
		}
		books = append(books, book)
	}
	if err := cursor.Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.DatabaseEvent(fmt.Sprintf("error while parsing cursor: %s", err.Error()))
		fmt.Fprintf(w, "error while parsing cursor: %s", err.Error())
		return
	}
	json.NewEncoder(w).Encode(books)

	log.DatabaseEvent(fmt.Sprintf("Fetch successful, #books: %d", len(books)))
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
	// filters

	switch filterParam {
	case "bookid":
		bookId, err := strconv.Atoi(r.URL.Query().Get("bookid"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.DatabaseEvent(fmt.Sprintf("error while casting bookid: %s", err.Error()))
			fmt.Fprintf(w, "error while casting bookid: %s", err.Error())
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
			w.WriteHeader(http.StatusInternalServerError)
			log.DatabaseEvent(fmt.Sprintf("error while casting price: %s", err.Error()))
			fmt.Fprintf(w, "error while casting price: %s", err.Error())
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
		w.WriteHeader(http.StatusInternalServerError)
		log.DatabaseEvent(fmt.Sprintf("error while fetching data from database: %s", err.Error()))
		fmt.Fprintf(w, "error while fetching data from database: %s", err.Error())
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
		w.WriteHeader(http.StatusInternalServerError)
		log.Issue(err.Error())
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Issue(err.Error())
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
	var booksAsInterfaceSlice []interface{} = make([]interface{}, len(booksAsSlice))
	for i, d := range booksAsSlice {
		booksAsInterfaceSlice[i] = d
	}

	if _, err := collection.InsertMany(context.Background(), booksAsInterfaceSlice); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
	fmt.Fprintf(w, "Book records sucessully injected : %d", len(booksAsSlice))
	log.DatabaseEvent(fmt.Sprintf("Insert successful, #Books: %d", len(booksAsSlice)))
}
