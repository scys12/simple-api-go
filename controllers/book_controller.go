package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/globalsign/mgo/bson"
	"github.com/gorilla/mux"
	"github.com/scys12/simple-api-go/models"
)

var (
	dao = models.Books{}
)

func responseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var books []models.Books
	books, err := dao.FindAllBooks()
	if err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, books)

}

func FindBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	result, err := dao.FindBookById(id)
	if err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, result)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var book models.Books
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		responseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	book.ID = bson.NewObjectId()
	if err := dao.InsertBook(book); err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusCreated, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var params models.Books
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		responseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.UpdateBook(params); err != nil {
		responseWithJson(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if err := dao.RemoveBook(id); err != nil {
		responseWithJson(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	responseWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
