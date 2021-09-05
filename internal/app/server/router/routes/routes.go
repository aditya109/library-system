package routes

import (
	"context"
	"fmt"
	"net/http"

	"library-server/configs"
	"library-server/internal/app/pkg/db/mongodb"
	"library-server/internal/app/server/handlers"
	"library-server/pkg/logger"

	"go.mongodb.org/mongo-driver/mongo"
)

type Route struct {
	Path            string
	HandlerFunction func(http.ResponseWriter, *http.Request)
	MethodTypes     []string
}

type RoutingContext struct {
	config      configs.Config
	handlerName string
}

var err error
var collection *mongo.Collection
var hcontext context.Context
var standardLogger = logger.NewLogger()
var dbUri string

func (ctx RoutingContext) RouteToCorrespondingHandlerFunction(w http.ResponseWriter, r *http.Request) {
	if ctx.config.EnvironmentType == "PROD" {
	} else if ctx.config.EnvironmentType == "DEV" {
		dbUri = ctx.config.DevEnvs.MongodbUri
	}
	if collection, hcontext, err = mongodb.GetMongoDbCollection(ctx.config.DatabaseStrings.DatabaseName, ctx.config.DatabaseStrings.DatabaseCollections.BooksCollectionName, dbUri); err != nil {
		standardLogger.Issue(err.Error())
		return
	}
	hctx := handlers.HandlerContext{Collection: collection, Context: hcontext, W: w, R: r}
	switch ctx.handlerName {
	case "WelcomeHandler":
		handlers.WelcomeHandler(hctx)

	case "AddBookHandler":
		handlers.AddBookHandler(hctx)

	case "AddBooksHandler":
		handlers.AddBooksHandler(hctx)

	case "GetBooksHandler":
		handlers.GetBooksHandler(hctx)

	case "GetBookByBookIdHandler":
		hctx.FilterParam = "bookid"
		handlers.GetBookByBookIdHandler(hctx)

	case "GetBookByBookNameHandler":
		hctx.FilterParam = "bookname"
		handlers.GetBookByBookNameHandler(hctx)

	case "GetBookByIsbnHandler":
		hctx.FilterParam = "isbn"
		handlers.GetBookByIsbnHandler(hctx)

	case "GetBookByPriceHandler":
		hctx.FilterParam = "price"
		handlers.GetBookByPriceHandler(hctx)

	case "GetBookByBookAuthorNameHandler":
		hctx.FilterParam = "bookauthor"
		handlers.GetBookByBookAuthorNameHandler(hctx)
	}
}

func GetRoutes(prefix string, config configs.Config) []Route {
	routes := []Route{
		Route{
			Path:            "/",
			HandlerFunction: RoutingContext{config, "WelcomeHandler"}.RouteToCorrespondingHandlerFunction,
			MethodTypes:     []string{"GET"},
		},
		Route{
			Path:            fmt.Sprintf("%s/book", prefix),
			HandlerFunction: RoutingContext{config, "AddBookHandler"}.RouteToCorrespondingHandlerFunction,
			MethodTypes:     []string{"POST"},
		},
		Route{
			Path:            fmt.Sprintf("%s/books", prefix),
			HandlerFunction: RoutingContext{config, "GetBooksHandler"}.RouteToCorrespondingHandlerFunction,
			MethodTypes:     []string{"GET"},
		},
		Route{
			Path:            fmt.Sprintf("%s/bookByBookId", prefix),
			HandlerFunction: RoutingContext{config, "GetBookByBookIdHandler"}.RouteToCorrespondingHandlerFunction,
			MethodTypes:     []string{"GET"},
		},
		Route{
			Path:            fmt.Sprintf("%s/bookByBookName", prefix),
			HandlerFunction: RoutingContext{config, "GetBookByBookNameHandler"}.RouteToCorrespondingHandlerFunction,
			MethodTypes:     []string{"GET"},
		},
		Route{
			Path:            fmt.Sprintf("%s/bookByIsbn", prefix),
			HandlerFunction: RoutingContext{config, "GetBookByIsbnHandler"}.RouteToCorrespondingHandlerFunction,
			MethodTypes:     []string{"GET"},
		},
		Route{
			Path:            fmt.Sprintf("%s/bookByPrice", prefix),
			HandlerFunction: RoutingContext{config, "GetBookByPriceHandler"}.RouteToCorrespondingHandlerFunction,
			MethodTypes:     []string{"GET"},
		},
		Route{
			Path:            fmt.Sprintf("%s/bookByBookAuthorName", prefix),
			HandlerFunction: RoutingContext{config, "GetBookByBookAuthorNameHandler"}.RouteToCorrespondingHandlerFunction,
			MethodTypes:     []string{"GET"},
		},
		Route{
			Path:            fmt.Sprintf("%s/books", prefix),
			HandlerFunction: RoutingContext{config, "AddBooksHandler"}.RouteToCorrespondingHandlerFunction,
			MethodTypes:     []string{"POST"},
		},
	}
	return routes
}
