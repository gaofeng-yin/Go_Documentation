package app

import (
	"fmt"
	"log"
	"net/http"
	"rest_api/controller"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Books api")
	})

	router.HandleFunc("/api/books", controller.ReturnAllBooks).Methods("GET")

	router.HandleFunc("/api/book", controller.CreateNewBook).Methods("POST")

	router.HandleFunc("/api/book/{id}", controller.DeleteBook).Methods("DELETE")

	router.HandleFunc("/api/book/{id}", controller.UpdateBook).Methods("PUT")

	router.HandleFunc("/api/book/{id}", controller.ReturnSingleBook).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
