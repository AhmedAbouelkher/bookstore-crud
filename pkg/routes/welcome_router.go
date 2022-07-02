package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)


func RegisterWelcomeRoute(r *mux.Router) {
	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		json.NewEncoder(res).Encode(map[string]string{
			"message": "Welcome to Bookstore API",
		})

	}).Methods(http.MethodGet)
}