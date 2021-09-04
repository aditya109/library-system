package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"library-server/pkg/logger"
)

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
	// defer client.Disconnect(ctx)
	return client, ctx, nil
}

// func CloseClientDB(client *mongo.Client) {
// 	if client == nil {
// 		return
// 	}

// 	err := client.Disconnect(context.TODO())
// 	if err != nil {
// 		standardLogger.Issue(err.Error())
// 	}

// 	standardLogger.ServerEvent("Connection to MongoDB closed.")
// }

func GetMongoDbCollection(dbName string, collectionName string, uri string) (*mongo.Collection, context.Context, error) {
	// initializing standard logger
	standardLogger = logger.NewLogger()
	client, ctx, err := GetMongoDbClient(uri)

	if err != nil {
		return nil, nil, err
	}

	collection := client.Database(dbName).Collection(collectionName)

	return collection, ctx, nil
}
