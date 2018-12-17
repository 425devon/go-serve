package main

import (
	"log"
	"net/http"

	"github.com/425devon/api-server/routers"
)

func main() {
	router := routers.NewRouter(routers.AllRoutes())
	log.Fatal(http.ListenAndServe(":8080", router))
}

/*
➜  ~ curl -d '{"isdn":"1234","title":"testing","author":"Devon","pages":732}' -H "application/json" -X POST http://localhost:8080/books
{"meta":null,"data":{"isdn":"1234","title":"testing","author":"Devon","pages":732}}
➜  ~ curl -H "Accept:application/json" http://localhost:8080/books
{"meta":null,"data":[{"isdn":"1234","title":"testing","author":"Devon","pages":732}]}
➜  ~

*/
