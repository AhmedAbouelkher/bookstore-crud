package controllers

import (
	"bookstore-crud/pkg/configs"
	"bookstore-crud/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


func InitBookMigrations() {
	db := configs.GetDB()
	db.AutoMigrate(&models.Book{})
}


func FetchAllBooksHandler(res http.ResponseWriter, req *http.Request) {
	db := configs.GetDB()
	books := []models.Book{}
	db.Find(&books)
	json.NewEncoder(res).Encode(books)
}

func FetchBookByIdHandler(res http.ResponseWriter, req *http.Request){
	db := configs.GetDB()
	bookId := extractIdParamFromRequest(req)
	var book models.Book
	db.First(&book, bookId)
	json.NewEncoder(res).Encode(book)
}

func CreateBookHandler(res http.ResponseWriter, req *http.Request) {
	
}


func UpdateBookByIdHandler(res http.ResponseWriter, req *http.Request) {
	
}

func DeleteBookByIdHandler(res http.ResponseWriter, req *http.Request) {
	
}

func extractIdParamFromRequest(req *http.Request) uint {
	params := mux.Vars(req)
	id, err := strconv.ParseInt(params["id"], 10, 8)
	
	if err != nil {
		return 0
	}

	return uint(id)
}