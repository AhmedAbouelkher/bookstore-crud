package controllers

import (
	"bookstore-crud/pkg/configs"
	"bookstore-crud/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/manveru/faker"
)


func InitBookMigrations() {
	db := configs.GetDB()
	db.AutoMigrate(&models.Book{})
}


func FetchAllBooksHandler(res http.ResponseWriter, req *http.Request) {
	db := configs.GetDB()
	books := []models.Book{}
	result := db.Find(&books)
	
	if result.Error != nil {
		throwDBHttpError(res, req, result.Error)
		return
	}
	json.NewEncoder(res).Encode(books)
}

func FetchBookByIdHandler(res http.ResponseWriter, req *http.Request) {
	db := configs.GetDB()
	bookId := extractIdParamFromRequest(req)

	var book models.Book
	result := db.First(&book, bookId)

	if result.Error != nil {
		throwDBHttpError(res, req, result.Error)
		return
	}

	if book.Title == "" {
		res.WriteHeader(http.StatusNotFound)
		json.NewEncoder(res).Encode(map[string]string{
			"message": "Book was not found",
		})
		return
	}
	json.NewEncoder(res).Encode(book)
}

func CreateBookHandler(res http.ResponseWriter, req *http.Request) {
	var book models.Book
	err := json.NewDecoder(req.Body).Decode(&book)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]any{
			"status": http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	// Using faker
	fake, err := faker.New("en")
	if err != nil {
		panic(err)
	}

	book.Title = fake.Name()
	book.Price = fake.Rand.Float32()

	db := configs.GetDB()
	result := db.Create(&book)

	if result.Error != nil {
		throwDBHttpError(res, req, result.Error)
		return
	}

	json.NewEncoder(res).Encode(book)
}

func UpdateBookByIdHandler(res http.ResponseWriter, req *http.Request) {
	
	bookId := extractIdParamFromRequest(req)

	// Decode updated data
	var bookUpdate models.Book
	err := json.NewDecoder(req.Body).Decode(&bookUpdate)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]any{
			"status": http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	db := configs.GetDB()

	// Fetching original row from database
	var book models.Book
	result := db.First(&book, bookId)
	
	if result.Error != nil {
		throwDBHttpError(res, req, result.Error)
		return
	}
	
	// Updating table row fields
	book.Title = bookUpdate.Title
	book.Price = bookUpdate.Price
	
	// Saving updated row from database
	updateResult := db.Save(&book)
	
	if updateResult.Error != nil {
		throwDBHttpError(res, req, result.Error)
		return
	}

	json.NewEncoder(res).Encode(book)
}

func DeleteBookByIdHandler(res http.ResponseWriter, req *http.Request) {
	db := configs.GetDB()
	var book models.Book

	bookId := extractIdParamFromRequest(req)

	result := db.Delete(&book, bookId)

	if result.Error != nil {
		throwDBHttpError(res, req, result.Error)
		return
	}

	json.NewEncoder(res).Encode(map[string]any{
		"message": "Book was deleted",
	})
}

func extractIdParamFromRequest(req *http.Request) uint {
	params := mux.Vars(req)
	id, err := strconv.ParseInt(params["id"], 10, 8)
	
	if err != nil {
		return 0
	}

	return uint(id)
}

func throwDBHttpError(res http.ResponseWriter, req *http.Request, err error) {
	res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(map[string]any{
			"status": http.StatusBadRequest,
			"error": err.Error(),
		})
}