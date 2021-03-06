package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/425devon/api-server/models"
	"github.com/julienschmidt/httprouter"
)

// Index display message for '/' route
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

//BookCreate for the books Create action
// POST /books
func BookCreate(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	book := &models.Book{}
	if err := populateModelFromHandler(w, r, params, book); err != nil {
		writeErrorResponse(w, http.StatusUnprocessableEntity, "Unprocessible Entity")
		return
	}
	models.Bookstore[book.ISDN] = book
	writeOKResponse(w, book)
}

// BookIndex Handler for the books index action
// GET /books
func BookIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	books := []*models.Book{}
	for _, book := range models.Bookstore {
		books = append(books, book)
	}
	writeOKResponse(w, books)
}

// BookShow Handler for the books Show action
// GET /books/:isdn
func BookShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	isdn := params.ByName("isdn")
	book, ok := models.Bookstore[isdn]
	if !ok {
		// No book with the isdn in the url has been found
		writeErrorResponse(w, http.StatusNotFound, "Record Not Found")
		return
	}
	writeOKResponse(w, book)
}

// Writes the response as a standard JSON response with StatusOK
func writeOKResponse(w http.ResponseWriter, m interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(&models.JsonResponse{Data: m}); err != nil {
		writeErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
	}
}

// Writes the error response as a Standard API JSON response with a response code
func writeErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	json.
		NewEncoder(w).
		Encode(&models.JsonErrorResponse{Error: &models.ApiError{Status: errorCode, Title: errorMsg}})
}

//Populates a model from the params in the Handler
func populateModelFromHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params, model interface{}) error {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		return err
	}
	if err := r.Body.Close(); err != nil {
		return err
	}
	if err := json.Unmarshal(body, model); err != nil {
		return err
	}
	return nil
}
