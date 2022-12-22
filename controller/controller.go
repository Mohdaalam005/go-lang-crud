package controller

import (
	"github.com/gorilla/mux"
	"github.com/mohdaalam005/crud-with-postgres/service"
	"net/http"
)

func Route() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", service.GetAllBook).Methods("GET")
	router.HandleFunc("/books/{id}", service.GetBook).Methods("GET")
	router.HandleFunc("/books/{id}", service.DeleteBook).Methods("DELETE")
	router.HandleFunc("/books/{id}", service.UpdateBook).Methods("PUT")
	router.HandleFunc("/books", service.CreateBook).Methods("POST")
	http.ListenAndServe(":8080", router)
	return router
}
