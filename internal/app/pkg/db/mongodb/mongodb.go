package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"library-server/pkg/logger"
)

type DatabaseContext struct {
	DbName string
	CollectionName string
	Uri string
}
var standardLogger *logger.StandardLogger

//GetMongoDbClient get connection of mongodb
func GetMongoDbClient(uri string) (*mongo.Client, context.Context, error) {

	// client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		standardLogger.Issue(err.Error())
		return nil, nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		standardLogger.Issue(err.Error())
		return nil, nil, err
	}
	return client, ctx, nil
}

func GetMongoDbCollection(dctx DatabaseContext) (*mongo.Collection, context.Context, error) {
	// initializing standard logger
	standardLogger = logger.NewLogger()
	client, ctx, err := GetMongoDbClient(dctx.Uri)
	if err != nil {
		return nil, nil, err
	}
	collection := client.Database(dctx.DbName).Collection(dctx.CollectionName)
	return collection, ctx, nil
}
