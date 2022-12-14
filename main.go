package main

import (
	"cheapest-bookiego/api/book"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello World!")
}

var ApiV1 = "/api/v1"

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc(ApiV1+"/", homeLink)
	router.HandleFunc(ApiV1+"/book/search", book.SearchBook).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
