package bookrepository

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

var (
	response []byte
	err      error
	cursor   *mongo.Cursor
	books    []models.Book
	book     models.Book
	filter   bson.M
)

type BookRepositoryContext struct {
	Collection *mongo.Collection
	W          http.ResponseWriter
	Context    context.Context
}

var log = logger.NewLogger()

func InsertOneBook(ctx BookRepositoryContext, book models.Book) ([]byte, error) {
	res, err := ctx.Collection.InsertOne(context.Background(), book)
	if err != nil {
		logger.RaiseAlert(ctx.W, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	if response, err = json.Marshal(res); err != nil {
		logger.RaiseAlert(ctx.W, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	log.DatabaseEvent(fmt.Sprintf("Insert successful, BookID: %d", book.BookId))
	return response, nil
}

func GetAllBooks(ctx BookRepositoryContext) ([]models.Book, error) {
	if cursor, err = ctx.Collection.Find(ctx.Context, bson.M{}); err != nil { // bson.M{}, we passes empty filter, so we can get all the data
		logger.RaiseAlert(ctx.W, fmt.Sprintf("error while fetching data from database: %s", err.Error()), http.StatusInternalServerError)
		return nil, err
	}
	defer cursor.Close(context.TODO())
	if books, err = convertBooksDbResponseToModelBook(ctx); err != nil {
		return nil, err
	}
	log.DatabaseEvent(fmt.Sprintf("Fetch successful, #books: %d", len(books)))
	return books, nil
}

func GetBookByBookId(ctx BookRepositoryContext, param string, id string) (models.Book, error) {
	bookId, err := strconv.Atoi(id)
	if err != nil {
		logger.RaiseAlert(ctx.W, fmt.Sprintf("error while casting bookid: %s", err.Error()), http.StatusInternalServerError)
		return models.Book{}, err
	}
	filter = bson.M{param: bookId}
	if book, err := FindBookByFilter(ctx, filter); err != nil {
		return models.Book{}, err
	} else {
		return book, nil
	}
}

func GetBookByBookName(ctx BookRepositoryContext, param string, bookname string) (models.Book, error) {
	filter = bson.M{param: bson.M{"$regex": primitive.Regex{Pattern: bookname, Options: "i"}}}
	if book, err := FindBookByFilter(ctx, filter); err != nil {
		return models.Book{}, err
	} else {
		return book, nil
	}
}

func GetBookByBookAuthorName(ctx BookRepositoryContext, param string, bookauthorname string) (models.Book, error) {
	filter = bson.M{param: bson.M{"$regex": primitive.Regex{Pattern: bookauthorname, Options: "i"}}}
	if book, err := FindBookByFilter(ctx, filter); err != nil {
		return models.Book{}, err
	} else {
		return book, nil
	}
}

func GetBookByIsbn(ctx BookRepositoryContext, param string, isbn string) (models.Book, error) {
	filter = bson.M{param: bson.M{"$regex": primitive.Regex{Pattern: isbn, Options: "i"}}}
	if book, err := FindBookByFilter(ctx, filter); err != nil {
		return models.Book{}, err
	} else {
		return book, nil
	}
}

func GetBookByPrice(ctx BookRepositoryContext, param string, amt string) (models.Book, error) {
	price, err := strconv.ParseFloat(amt, 64)
	if err != nil {
		logger.RaiseAlert(ctx.W, fmt.Sprintf("error while casting price: %s", err.Error()), http.StatusInternalServerError)
		return models.Book{}, err
	}
	filter = bson.M{param: price}
	if book, err := FindBookByFilter(ctx, filter); err != nil {
		return models.Book{}, err
	} else {
		return book, nil
	}
}

func InsertMultipleBook(ctx BookRepositoryContext, books []interface{}) ([]byte, error) {
	if _, err := ctx.Collection.InsertMany(context.Background(), books); err != nil {
		logger.RaiseAlert(ctx.W, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	return nil, nil
}

func FindBookByFilter(ctx BookRepositoryContext, filter bson.M) (models.Book, error) {
	if err = ctx.Collection.FindOne(ctx.Context, filter).Decode(&book); err != nil {
		logger.RaiseAlert(ctx.W, fmt.Sprintf("error while fetching data from database: %s", err.Error()), http.StatusInternalServerError)
		return models.Book{}, err
	}
	log.DatabaseEvent(fmt.Sprintf("Fetch successful, #books ID: %d", book.BookId))
	return book, nil
}

/**
=========HELPER FUNCTIONS============
*/

func convertBooksDbResponseToModelBook(ctx BookRepositoryContext) ([]models.Book, error) {
	for cursor.Next(context.TODO()) {
		var book models.Book
		if err = cursor.Decode(&book); err != nil {
			logger.RaiseAlert(ctx.W, fmt.Sprintf("error while reading stream data from database: %s", err.Error()), http.StatusInternalServerError)
			return nil, err
		}
		books = append(books, book)
	}
	if err := cursor.Err(); err != nil {
		logger.RaiseAlert(ctx.W, fmt.Sprintf("error while parsing cursor: %s", err.Error()), http.StatusInternalServerError)
		return nil, err
	}
	return books, err
}
