package controllers

import (
	"bookstore-crud/pkg/configs"
	"bookstore-crud/pkg/models"
	"fmt"
	"net/http"
)


func FetchAllAuthorsHandler(res http.ResponseWriter, req *http.Request) {
	db := configs.GetDB()
	var authors []models.Author
	db.Model(&models.Author{}).Find(&authors);

	sendResponse(res, authors)
}

func FetchAuthorByIdHandler(res http.ResponseWriter, req *http.Request) {	
	author, err := checkAuthorId(res, req)

	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		sendResponse(res, Payload{
			"message": "Failed to fetch author",
		})
		return
	}
	
	sendResponse(res, author)
}

func CreateAuthorHandler(res http.ResponseWriter, req *http.Request) {
	db := configs.GetDB()

	var author models.Author
	err := extractRequestBody(BodyExtraction{
		res: res,
		req: req,
		decodedValue: &author,
	})

	if err != nil {
		return
	}

	dbError := db.Create(&author).Error

	if dbError != nil {
		throwDBHttpError(res, req, dbError)
		return
	}

	sendResponse(res, author)
}

func UpdateAuthorByIdHandler(res http.ResponseWriter, req *http.Request) {
	db := configs.GetDB()
	authorId := extractIntegerParamFromRequest(req, "authorId")

	var update models.Author
	err := extractRequestBody(BodyExtraction{
		res: res,
		req: req,
		decodedValue: &update,
	})

	if err != nil {
		return
	}

	author := models.Author{
		ID: authorId,
	}
	
	dbError := db.Model(&author).Updates(map[string]interface{}{
		"name": update.Name,
		"age": update.Age,
	}).Error

	if dbError != nil {
		throwDBHttpError(res, req, dbError)
		return
	}

	sendResponse(res, author)
}

func DeleteAuthorByIdHandler(res http.ResponseWriter, req *http.Request) {
	db := configs.GetDB()
	authorId := extractIdParamFromRequest(req)
	
	if dbError := db.Delete(models.Author{}, authorId).Error; dbError != nil {
		throwDBHttpError(res, req, dbError)
		return
	}

	sendResponse(res, Payload{
		"message": fmt.Sprintf("Deleted author with id %d", authorId),
	})	
}

func checkAuthorId(res http.ResponseWriter, req *http.Request) (author models.Author, err error) {
	db := configs.GetDB()
	authorId := extractIntegerParamFromRequest(req, "authorId")
	
	var dbAuthor models.Author
	db.First(&dbAuthor, authorId)

	if dbAuthor.ID == 0 {
		return models.Author{}, fmt.Errorf("failed to fetch author with id %d", authorId)
	}

	return dbAuthor, nil
}