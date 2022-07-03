package controllers

import (
	"bookstore-crud/pkg/configs"
	"bookstore-crud/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// General Books

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

// Author Books

func FetchAuthorBooksByIdHandler(res http.ResponseWriter, req *http.Request) {
	db := configs.GetDB()
	author, err := checkAuthorId(res, req)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		sendResponse(res, Payload{
			"message": "Failed to fetch author",
		})
		return
	}

	
	books := []models.Book{}
	
	dbError := db.Where("author_id = ?", author.ID).Find(&books).Error
	if dbError != nil {
		throwDBHttpError(res, req, dbError)
		return
	}
	
	fmt.Println(author, len(books))

	json.NewEncoder(res).Encode(books)
}

func CreateAuthorBookByIdHandler(res http.ResponseWriter, req *http.Request) {
	db := configs.GetDB()
	
	// Validate Author id
	author, err := checkAuthorId(res, req)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		sendResponse(res, Payload{
			"message": fmt.Sprintf("Failed to fetch author with id %d", author.ID),
		})
		return
	}

	// Create new book and attach Author id
	var book models.Book

	if extractRequestBody(BodyExtraction{
		res: res,
		req: req,
		decodedValue: &book,
	}) != nil {
		return
	}

	book.AuthorID = int(author.ID)
	
	if err := db.Create(&book).Error; err != nil {
		throwDBHttpError(res, req, err)
		return
	}

	sendResponse(res, book)
}

func UpdateAuthorBookByIdHandler(res http.ResponseWriter, req *http.Request) {
	db := configs.GetDB()
	
	// Validate Author id
	author, err := checkAuthorId(res, req)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		sendResponse(res, Payload{
			"message": fmt.Sprintf("Failed to fetch author with id %d", author.ID),
		})
		return
	}

	// Decode updated data
	var update models.Book

	if extractRequestBody(BodyExtraction{
		res: res,
		req: req,
		decodedValue: &update,
	}) != nil {
		return
	}

	bookId := extractIntegerParamFromRequest(req, "bookId")

	book := models.Book{
		ID: bookId,
		AuthorID: int(author.ID),
	}

	// dbError := db.Where("id = ? AND author_id = ?", bookId, author.ID).First(&book).Error
	dbError := db.Model(&book).Updates(map[string]interface{}{
			"title": update.Title,
			"price": update.Price,
		},
	).Error

	if dbError != nil {
		throwDBHttpError(res, req, dbError)
		return
	}

	sendResponse(res, Payload{
		"message": "Updated",
		"book": book,
	})
}