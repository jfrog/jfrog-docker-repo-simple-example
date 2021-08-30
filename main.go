package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func getBooks(w http.ResponseWriter, r *http.Request) {

	log.Println("get all books")

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {

	log.Println("get single books")

	params := mux.Vars(r)
	log.Println(params)

	k, _ := strconv.Atoi(params["id"])

	for _, book := range books {
		if book.ID == k {
			json.NewEncoder(w).Encode(&book)
		}

	}
}

func addBook(w http.ResponseWriter, r *http.Request) {

	log.Println("add new book")

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)
	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {

	log.Println("update a books")

	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for k, v := range books {
		if v.ID == book.ID {
			books[k] = book
		}
	}
	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {

	log.Println("delete a books")

	params := mux.Vars(r)
	log.Println(params)

	id, _ := strconv.Atoi(params["id"])

	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)

}

func main() {

	router := mux.NewRouter()

	books = append(books,
		Book{ID: 1, Title: "Golang pointers", Author: "Mr. Golang", Year: "2010"},
		Book{ID: 2, Title: "Goroutines", Author: "Mr. Goroutine", Year: "2011"},
		Book{ID: 3, Title: "Golang routers", Author: "Mr. Router", Year: "2012"},
		Book{ID: 4, Title: "Golang concurrency", Author: "Mr. Currency", Year: "2013"},
		Book{ID: 5, Title: "Golang good parts", Author: "Mr. Good", Year: "2014"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3333", router))
}
