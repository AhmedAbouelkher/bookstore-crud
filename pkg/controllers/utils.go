package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Wraps map[string]any
type Payload map[string]any

func extractIdParamFromRequest(req *http.Request) uint {
	return extractIntegerParamFromRequest(req, "id")
}

func extractIntegerParamFromRequest(req *http.Request, param string) uint {
	params := mux.Vars(req)
	p, err := strconv.ParseInt(params[param], 10, 8)
	
	if err != nil {
		return 0
	}

	return uint(p)
}

type BodyExtraction struct {
	res http.ResponseWriter
	req *http.Request
	decodedValue any
}

func extractRequestBody(extraction BodyExtraction) error {
	err := json.NewDecoder(extraction.req.Body).Decode(extraction.decodedValue)

	if err == nil {
		return nil
	}

	extraction.res.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(extraction.res).Encode(map[string]any{
		"status": http.StatusBadRequest,
		"error": err.Error(),
	})

	return err
}


func throwDBHttpError(res http.ResponseWriter, req *http.Request, err error) {
	res.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(res).Encode(map[string]any{
		"status": http.StatusBadRequest,
		"error": err.Error(),
	})
}

func sendResponse(w http.ResponseWriter, v any) {
	err := json.NewEncoder(w).Encode(v)
	if err == nil {
		return
	}
	
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]any{
		"status": 500,
		"error": err.Error(),
	})
}