package routers

import (
	"github.com/425devon/api-server/handlers"
	"github.com/julienschmidt/httprouter"
)

/*
Route Define all the routes here.
A new Route entry passed to the routes slice will be automatically
translated to a handler with the NewRouter() function
*/
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc httprouter.Handle
}

//Routes []Route
type Routes []Route

//AllRoutes returns []Routes
func AllRoutes() Routes {
	routes := Routes{
		Route{"Index", "GET", "/", handlers.Index},
		Route{"BookIndex", "GET", "/books", handlers.BookIndex},
		Route{"Bookshow", "GET", "/books/:isdn", handlers.BookShow},
		Route{"Bookshow", "POST", "/books", handlers.BookCreate},
	}
	return routes
}
