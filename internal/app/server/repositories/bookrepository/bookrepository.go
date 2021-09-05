package bookrepository

import (
	"context"
	"encoding/json"
	"fmt"
	"library-server/internal/app/server/models"
	"library-server/pkg/logger"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	response []byte
	err      error
	cursor   *mongo.Cursor
	books    []models.Book
	book     models.Book
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
		logger.RaiseAlert(ctx.W, log, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	if response, err = json.Marshal(res); err != nil {
		logger.RaiseAlert(ctx.W, log, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	log.DatabaseEvent(fmt.Sprintf("Insert successful, BookID: %d", book.BookId))
	return response, nil
}

func GetAllBooks(ctx BookRepositoryContext) ([]models.Book, error) {
	if cursor, err = ctx.Collection.Find(ctx.Context, bson.M{}); err != nil { // bson.M{}, we passes empty filter, so we can get all the data
		logger.RaiseAlert(ctx.W, log, fmt.Sprintf("error while fetching data from database: %s", err.Error()), http.StatusInternalServerError)
		return nil, err
	}
	defer cursor.Close(context.TODO())
	if books, err = convertBooksDbResponseToModelBook(ctx); err != nil {
		return nil, err
	}
	log.DatabaseEvent(fmt.Sprintf("Fetch successful, #books: %d", len(books)))
	return books, nil
}

/**
=========HELPER FUNCTIONS============
*/

func convertBooksDbResponseToModelBook(ctx BookRepositoryContext) ([]models.Book, error) {
	for cursor.Next(context.TODO()) {
		var book models.Book
		if err = cursor.Decode(&book); err != nil {
			logger.RaiseAlert(ctx.W, log, fmt.Sprintf("error while reading stream data from database: %s", err.Error()), http.StatusInternalServerError)
			return nil, err
		}
		books = append(books, book)
	}
	if err := cursor.Err(); err != nil {
		logger.RaiseAlert(ctx.W, log, fmt.Sprintf("error while parsing cursor: %s", err.Error()), http.StatusInternalServerError)
		return nil, err
	}
	return books, err
}
