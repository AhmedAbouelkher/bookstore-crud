package routes

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Payload map[string]any

func RegisterWelcomeRoute(r *mux.Router) {
	r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		json.NewEncoder(res).Encode(Payload{
			"message": "Welcome to Bookstore API",
			"status": 200,
		})

	}).Methods(http.MethodGet)
}