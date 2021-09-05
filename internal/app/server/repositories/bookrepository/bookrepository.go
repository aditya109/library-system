package bookrepository

import (
	"context"
	"encoding/json"
	"fmt"
	"library-server/internal/app/server/models"
	"library-server/pkg/logger"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)
var res []byte
var err error 

type BookRepositoryContext struct {
	Book       models.Book
	Collection *mongo.Collection
	W          http.ResponseWriter
	Log        *logger.StandardLogger
}

func InsertOneBook(ctx BookRepositoryContext) ([]byte, error) {
	res, err := ctx.Collection.InsertOne(context.Background(), ctx.Book)
	if err != nil {
		logger.RaiseAlert(ctx.W, ctx.Log, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	response, err := json.Marshal(res)
	if err != nil {
		logger.RaiseAlert(ctx.W, ctx.Log, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	ctx.Log.DatabaseEvent(fmt.Sprintf("Insert successful, BookID: %d", ctx.Book.BookId))
	return response, nil
}
