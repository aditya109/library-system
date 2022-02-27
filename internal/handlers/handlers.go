package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	resp "github.com/aditya109/library-system/internal/responses"
	svc "github.com/aditya109/library-system/internal/services"
	logger "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

// WelcomeHandler returns welcome message for home URL
func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	responseStatusCode := 200
	var response resp.WelcomeResponseWrapper = "welcome to server ⚡ !"
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(responseStatusCode)
	w.Write([]byte(response))
	logger.Info(fmt.Sprintf("STATUS: %d === / route was hit", responseStatusCode))
}

// GetItemsHandler returns a static list of items, no query params required
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	var responseStatusCode int
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	items, err := svc.GetAllItems()
	if err != nil {
		responseStatusCode = http.StatusInternalServerError
		w.WriteHeader(responseStatusCode)
		logger.Error(err)
		w.Write([]byte(fmt.Sprintf("error occurred while getting items, %s", err.Error())))
	} else {
		responseStatusCode = http.StatusOK
		w.WriteHeader(responseStatusCode)
		json.NewEncoder(w).Encode(items)
	}
	logger.Info(fmt.Sprintf("STATUS: %d === /items route was hit", responseStatusCode))
}

// GetItemWithIDHandler returns item with specified id
func GetItemWithIDHandler(w http.ResponseWriter, r *http.Request) {
	var responseStatusCode int
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		responseStatusCode = http.StatusInternalServerError
		w.WriteHeader(responseStatusCode)
		logger.Error(err)
		w.Write([]byte(fmt.Sprintf("error occurred while type-casting id, %s", err.Error())))
	}
	item, err := svc.GetItemByID(int64(id))
	if err != nil {
		responseStatusCode = http.StatusInternalServerError
		logger.Error(err)
		w.WriteHeader(responseStatusCode)
		w.Write([]byte(fmt.Sprintf("error occurred while getting items, %s", err.Error())))
	} else {
		responseStatusCode = http.StatusOK
		w.WriteHeader(responseStatusCode)
		json.NewEncoder(w).Encode(item)
	}
	logger.Info(fmt.Sprintf("STATUS: %d === /item/{id} route was hit, ", responseStatusCode))
}

// GetWithQueryParamsHandler returns query-based filtered list of items
func GetWithQueryParamsHandler(w http.ResponseWriter, r *http.Request) {
	var responseStatusCode int
	idParam, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		responseStatusCode = http.StatusInternalServerError
		w.WriteHeader(responseStatusCode)
		logger.Error(err)
		w.Write([]byte(fmt.Sprintf("error occurred while type-casting id, %s", err.Error())))
	}
	var nameParam string = r.URL.Query().Get("name")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	items, err := svc.GetItemsByIDAndName(idParam, nameParam)
	if err != nil {
		responseStatusCode = http.StatusInternalServerError
		logger.Error(err)
		w.WriteHeader(responseStatusCode)
		w.Write([]byte(fmt.Sprintf("error occurred while getting items, %s", err.Error())))
	} else {
		responseStatusCode = http.StatusOK
		w.WriteHeader(responseStatusCode)
		json.NewEncoder(w).Encode(items)
	}
	logger.Info(fmt.Sprintf("STATUS: %d === /item/{id} route was hit, ", responseStatusCode))
}
