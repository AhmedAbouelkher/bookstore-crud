package main

import (
	"bookstore-crud/pkg/configs"
	"bookstore-crud/pkg/controllers"
	"bookstore-crud/pkg/middlewares"
	"bookstore-crud/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	configs.ConnectToDatabase()
	controllers.InitBookMigrations();
	defer configs.CloseDB()

	router := mux.NewRouter()
	routes.RegisterWelcomeRoute(router)
	routes.RegisterBooksRoutes(router)
	router.Use(middlewares.ContentTypeApplicationJsonMiddleware)

	srv := &http.Server{
		Handler: router,
		Addr: "127.0.0.1:8005",
	}
	log.Println("Starting the server...")
	log.Fatal(srv.ListenAndServe())
}