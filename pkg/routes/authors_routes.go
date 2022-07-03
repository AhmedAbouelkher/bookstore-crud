package routes

import (
	"bookstore-crud/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterAuthorsRoutes(r *mux.Router) {
	subRouter := r.PathPrefix("/authors").Subrouter()

	subRouter.HandleFunc("/", controllers.FetchAllAuthorsHandler).Methods(http.MethodGet)
	subRouter.HandleFunc("/", controllers.CreateAuthorHandler).Methods(http.MethodPost)
	subRouter.HandleFunc("/{authorId}", controllers.FetchAuthorByIdHandler).Methods(http.MethodGet)
	subRouter.HandleFunc("/{authorId}", controllers.UpdateAuthorByIdHandler).Methods(http.MethodPut)
	subRouter.HandleFunc("/{id}", controllers.DeleteAuthorByIdHandler).Methods(http.MethodDelete)
	
	bookRouter := subRouter.PathPrefix("/{authorId}/books").Subrouter()
	bookRouter.HandleFunc("/", controllers.FetchAuthorBooksByIdHandler).Methods(http.MethodGet)
	bookRouter.HandleFunc("/", controllers.CreateAuthorBookByIdHandler).Methods(http.MethodPost)
	bookRouter.HandleFunc("/{bookId}", controllers.UpdateAuthorBookByIdHandler).Methods(http.MethodPut)
}