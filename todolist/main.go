package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func init() {
	ConnectToDatabase()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/todoitems", GetToDoItems).Methods("GET")
	r.HandleFunc("/api/todoitems/complete", GetToDoItemInCompleted).Methods("GET")
	r.HandleFunc("/api/todoitems/{id}", GetToDoItemById).Methods("GET")
	r.HandleFunc("/api/todoitems", CreateToDoItem).Methods("POST")
	r.HandleFunc("/api/todoitems/{id}", DeleteToDoItem).Methods("DELETE")
	r.HandleFunc("/api/todoitems/{id}", UpdateToDoItem).Methods("PUT")
	srv := &http.Server{
		Handler:      r,
		Addr:         ":3000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
