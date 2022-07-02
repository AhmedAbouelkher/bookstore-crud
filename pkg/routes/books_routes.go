package routes

import (
	"bookstore-crud/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)


func RegisterBooksRoutes(r *mux.Router) {
	router := r.PathPrefix("/books").Subrouter()

	router.HandleFunc("/", 	   controllers.FetchAllBooksHandler).Methods(http.MethodGet)
	router.HandleFunc("/", 	   controllers.CreateBookHandler).Methods(http.MethodPost)
	router.HandleFunc("/{id}", controllers.FetchBookByIdHandler).Methods(http.MethodGet)
	router.HandleFunc("/{id}", controllers.UpdateBookByIdHandler).Methods(http.MethodPut)
	router.HandleFunc("/{id}", controllers.DeleteBookByIdHandler).Methods(http.MethodDelete)
}