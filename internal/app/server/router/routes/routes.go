package routes

import (
	"fmt"
	"net/http"

	"library-server/configs"
	"library-server/internal/app/pkg/db/mongodb"
	"library-server/internal/app/server/handlers"
	"library-server/pkg/logger"
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

func (ctx RoutingContext) RouteToCorrespondingHandlerFunction(w http.ResponseWriter, r *http.Request) {
	var dbUri string
	standardLogger := logger.NewLogger()
	if ctx.config.EnvironmentType == "PROD" {

	} else if ctx.config.EnvironmentType == "DEV" {
		dbUri = ctx.config.DevEnvs.MongodbUri
	}
	switch ctx.handlerName {
	case "WelcomeHandler":
		handlers.WelcomeHandler(w, r)

	case "AddBookHandler":
		collection, ctx, err := mongodb.GetMongoDbCollection(ctx.config.DatabaseStrings.DatabaseName,
			ctx.config.DatabaseStrings.DatabaseCollections.BooksCollectionName,
			dbUri)
		if err != nil {
			standardLogger.Issue(err.Error())
			return
		}
		hctx := handlers.HandlerContext{
			Collection: collection,
			Context:    ctx,
		}
		handlers.AddBookHandler(w, r, hctx)

	case "GetBooksHandler":
		collection, ctx, err := mongodb.GetMongoDbCollection(ctx.config.DatabaseStrings.DatabaseName,
			ctx.config.DatabaseStrings.DatabaseCollections.BooksCollectionName,
			dbUri)
		if err != nil {
			standardLogger.Issue(err.Error())
			return
		}
		hctx := handlers.HandlerContext{
			Collection: collection,
			Context:    ctx,
		}
		handlers.GetBooksHandler(w, r, hctx)

	case "GetBookByBookIdHandler":
		collection, ctx, err := mongodb.GetMongoDbCollection(ctx.config.DatabaseStrings.DatabaseName,
			ctx.config.DatabaseStrings.DatabaseCollections.BooksCollectionName,
			dbUri)
		if err != nil {
			standardLogger.Issue(err.Error())
			return
		}
		hctx := handlers.HandlerContext{
			Collection:  collection,
			Context:     ctx,
			FilterParam: "bookid",
		}
		handlers.GetBookByBookIdHandler(w, r, hctx)

	case "GetBookByBookNameHandler":
		collection, ctx, err := mongodb.GetMongoDbCollection(ctx.config.DatabaseStrings.DatabaseName,
			ctx.config.DatabaseStrings.DatabaseCollections.BooksCollectionName,
			dbUri)
		if err != nil {
			standardLogger.Issue(err.Error())
			return
		}
		hctx := handlers.HandlerContext{
			Collection:  collection,
			Context:     ctx,
			FilterParam: "bookname",
		}
		handlers.GetBookByBookNameHandler(w, r, hctx)

	case "GetBookByIsbnHandler":
		collection, ctx, err := mongodb.GetMongoDbCollection(ctx.config.DatabaseStrings.DatabaseName,
			ctx.config.DatabaseStrings.DatabaseCollections.BooksCollectionName,
			dbUri)
		if err != nil {
			standardLogger.Issue(err.Error())
			return
		}
		hctx := handlers.HandlerContext{
			Collection:  collection,
			Context:     ctx,
			FilterParam: "isbn",
		}
		handlers.GetBookByIsbnHandler(w, r, hctx)

	case "GetBookByPriceHandler":
		collection, ctx, err := mongodb.GetMongoDbCollection(ctx.config.DatabaseStrings.DatabaseName,
			ctx.config.DatabaseStrings.DatabaseCollections.BooksCollectionName,
			dbUri)
		if err != nil {
			standardLogger.Issue(err.Error())
			return
		}
		hctx := handlers.HandlerContext{
			Collection:  collection,
			Context:     ctx,
			FilterParam: "price",
		}
		handlers.GetBookByPriceHandler(w, r, hctx)

	case "GetBookByBookAuthorNameHandler":
		collection, ctx, err := mongodb.GetMongoDbCollection(ctx.config.DatabaseStrings.DatabaseName,
			ctx.config.DatabaseStrings.DatabaseCollections.BooksCollectionName,
			dbUri)
		if err != nil {
			standardLogger.Issue(err.Error())
			return
		}
		hctx := handlers.HandlerContext{
			Collection:  collection,
			Context:     ctx,
			FilterParam: "bookauthor",
		}
		handlers.GetBookByBookAuthorNameHandler(w, r, hctx)

	case "AddBooksHandler":
		collection, ctx, err := mongodb.GetMongoDbCollection(ctx.config.DatabaseStrings.DatabaseName,
			ctx.config.DatabaseStrings.DatabaseCollections.BooksCollectionName,
			dbUri)
		if err != nil {
			standardLogger.Issue(err.Error())
			return
		}
		hctx := handlers.HandlerContext{
			Collection: collection,
			Context:    ctx,
		}
		handlers.AddBooksHandler(w, r, hctx)
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
