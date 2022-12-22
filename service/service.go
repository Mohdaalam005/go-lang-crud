package service

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"github.com/mohdaalam005/crud-with-postgres/database"
	"github.com/mohdaalam005/crud-with-postgres/models"
	"log"
	"net/http"
)

func GetAllBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("fetching the book.................")
	db := database.Connect()
	defer db.Close()
	var book []models.Book
	err := db.Model(&book).Select()

	database.CheckNilErr(err)

	json.NewEncoder(w).Encode(book)
	log.Println("called")

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.Connect()
	defer db.Close()
	book := &models.Book{
		ID: uuid.New().String(),
	}
	_ = json.NewDecoder(r.Body).Decode(&book)

	_, err := db.Model(book).Insert()

	database.CheckNilErr(err)

	json.NewEncoder(w).Encode(book)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.Connect()
	defer db.Close()
	param := mux.Vars(r)

	book := &models.Book{ID: param["id"]}

	db.Model(book).Where("id = ?", book.ID).Select()

	json.NewEncoder(w).Encode(book)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.Connect()
	defer db.Close()
	param := mux.Vars(r)

	book := models.Book{ID: param["id"]}
	_ = json.NewDecoder(r.Body).Decode(&book)

	_, err := db.Model(&book).Where("id = ?", book.ID).Update()
	database.CheckNilErr(err)
	json.NewEncoder(w).Encode(book)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db := database.Connect()
	defer db.Close()
	param := mux.Vars(r)

	book := models.Book{
		ID: param["id"],
	}
	_, err := db.Model(&book).WherePK().Delete()
	database.CheckNilErr(err)

	json.NewEncoder(w).Encode("book deleted with" + book.ID)

}
