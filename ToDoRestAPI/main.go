package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	router "Alive/ToDoRestAPI/http/router"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/user/{userid}", router.GetToDos).Methods("GET")
	r.HandleFunc("/api/user/{userid}/{todoid}", router.GetToDo).Methods("GET")
	r.HandleFunc("/api/user/{userid}", router.CreateToDo).Methods("POST")
	r.HandleFunc("/api/user/{userid}/{todoid}", router.UpdateToDo).Methods("PUT")
	r.HandleFunc("/api/user/{userid}/{todoid}", router.DeleteToDo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
