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
	hctx := handlers.HandlerContext{Collection: collection, Context: hcontext}
	switch ctx.handlerName {
	case "WelcomeHandler":
		handlers.WelcomeHandler(w, r)

	case "AddBookHandler":
		handlers.AddBookHandler(w, r, hctx)

	case "AddBooksHandler":
		handlers.AddBooksHandler(w, r, hctx)

	case "GetBooksHandler":
		handlers.GetBooksHandler(w, r, hctx)

	case "GetBookByBookIdHandler":
		hctx.FilterParam = "bookid"
		handlers.GetBookByBookIdHandler(w, r, hctx)

	case "GetBookByBookNameHandler":
		hctx.FilterParam = "bookname"
		fmt.Println("Came here 1 !")
		handlers.GetBookByBookNameHandler(w, r, hctx)

	case "GetBookByIsbnHandler":
		hctx.FilterParam = "isbn"
		handlers.GetBookByIsbnHandler(w, r, hctx)

	case "GetBookByPriceHandler":
		hctx.FilterParam = "price"
		handlers.GetBookByPriceHandler(w, r, hctx)

	case "GetBookByBookAuthorNameHandler":
		hctx.FilterParam = "bookauthor"
		handlers.GetBookByBookAuthorNameHandler(w, r, hctx)


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
