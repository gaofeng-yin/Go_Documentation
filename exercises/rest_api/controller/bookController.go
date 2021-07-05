package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"rest_api/model"

	"github.com/gorilla/mux"
)

var Books = []model.Book{
	{Id: "1", Title: "Brothers Karamazov", Author: "Fiodor Dostoviesky", Genre: "romance", Isbn: "32132-das23-3213-23"},
	{Id: "2", Title: "1984", Author: "George Orwell", Genre: "drama", Isbn: "3133-dsad-a2-13-321"},
	{Id: "3", Title: "12 Rules for life", Author: "Jordan B.Peterson", Genre: "self help", Isbn: "321-4545-7342-123"},
	//more books
}

func ReturnAllBooks(w http.ResponseWriter, r *http.Request) {
	//this way is faster
	//json.NewEncoder(w).Encode(Books)

	//pretty output
	prettyJsonEncode := json.NewEncoder(w)
	prettyJsonEncode.SetIndent("", "   ")
	prettyJsonEncode.Encode(Books)
}

//return book with path variable {id}
func ReturnSingleBook(w http.ResponseWriter, r *http.Request) {
	//Get request variables
	vars := mux.Vars(r)
	//Get varaible value
	value := vars["id"]

	for _, book := range Books {
		if book.Id == value {
			json.NewEncoder(w).Encode(book)
		}
	}
}

//create new book and append to Books slice
func CreateNewBook(w http.ResponseWriter, r *http.Request) {
	//read request body content to varaible requestBody
	requestBody, _ := ioutil.ReadAll((r.Body))

	var book model.Book

	//parse JSON encoded data and store result in the pointer
	json.Unmarshal(requestBody, &book)

	//append to Books slice
	Books = append(Books, book)

	json.NewEncoder(w).Encode(book)
}

//Delete book by variable {id}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	value := vars["id"]

	//Loop through slice and when id match, append everything from beginning to index and next value of the index to the end.
	for index, book := range Books {
		if book.Id == value {
			Books = append(Books[:index], Books[index+1:]...)
		}
	}
}

//Update book by variable {id}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	value := vars["id"]

	requestBody, _ := ioutil.ReadAll((r.Body))

	var newBook model.Book

	json.Unmarshal(requestBody, &newBook)

	for index, article := range Books {
		if article.Id == value {
			Books[index].Title = newBook.Title
			Books[index].Author = newBook.Author
			Books[index].Genre = newBook.Genre
			Books[index].Isbn = newBook.Isbn
		}
	}
}
