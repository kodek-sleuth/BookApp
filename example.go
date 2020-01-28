package main

import (
	"encoding/json"
	"fmt"
	_"encoding/json"
	"log"
	"math/rand"
	"net/http"
	_"math/rand"
	"strconv"
	_"strconv"
	"github.com/gorilla/mux"
)

type Book struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Company *Company `json:"company"`
}

type Company struct {
	Name string `json:"name"`
	Age  int  `json:"age"`
}

var books []Book
var book Book

func getBooks(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(books)
}

func getBook(res http.ResponseWriter, req *http.Request ) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	fmt.Println(books)
	for _, book := range books {
		if book.Id == params["id"] {
			json.NewEncoder(res).Encode(book)
			return
		}
	}

	json.NewEncoder(res).Encode(&Book{})
}

func createBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	_ = json.NewDecoder(req.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(12000000))
	books = append(books, book)
	json.NewEncoder(res).Encode(book)
}


func main(){
	fmt.Println("Hello world")

	router := mux.NewRouter()
	books = append(books, Book{ Id: "isbn231jks", Title: "Mugerwa", Company: &Company{ Name: "Mugerwa", Age:  21 } })
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", router))
}


