package routers

import (
	"github.com/425devon/api-server/loggers"
	"github.com/julienschmidt/httprouter"
)

// NewRouter Reads from the routes slice to translate the values to httprouter.Handle
func NewRouter(routes Routes) *httprouter.Router {

	router := httprouter.New()
	for _, route := range routes {
		var handle httprouter.Handle

		handle = route.HandlerFunc
		handle = loggers.Logger(handle)

		router.Handle(route.Method, route.Path, handle)
	}

	return router
}
