package router

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	database "Alive/ToDoRestAPI/database"
	req "Alive/ToDoRestAPI/http/request"
	res "Alive/ToDoRestAPI/http/response"
)

const (
	toDosPerPage = 3
)

func GetToDos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	page, err := strconv.Atoi(r.FormValue("page"))
	if err != nil {
		log.Println("Failed converting String to Int")
	}

	tag := r.FormValue("tag")
	if tag == "" {
		tag = "%"
	}

	params := mux.Vars(r)
	db := database.ConnectDatabase()
	rows, err := db.Query("SELECT * FROM Todo WHERE UserId = ? AND Tag LIKE ? LIMIT ? OFFSET ?",
		params["userid"], tag, toDosPerPage, (page-1)*toDosPerPage)
	if err != nil {
		log.Fatal(err)
	}

	var toDosResponse res.ToDosResponse
	for rows.Next() {
		var toDo res.ToDo
		err = rows.Scan(&toDo.Id, &toDo.UserId, &toDo.Title, &toDo.Description, &toDo.Status, &toDo.Tag, &toDo.CreatedDate)
		if err != nil {
			log.Fatal(err)
		} else {
			toDosResponse.ToDos = append(toDosResponse.ToDos, &toDo)
		}
	}

	json.NewEncoder(w).Encode(toDosResponse.ToDos)
}

func GetToDo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := database.ConnectDatabase()
	rows, err := db.Query("SELECT * FROM Todo WHERE UserId = ? AND Id = ?",
		params["userid"], params["todoid"])
	if err != nil {
		log.Fatal(err)
	}

	var toDosResponse res.ToDosResponse
	for rows.Next() {
		var toDo res.ToDo
		err = rows.Scan(&toDo.Id, &toDo.UserId, &toDo.Title, &toDo.Description, &toDo.Status, &toDo.Tag, &toDo.CreatedDate)
		if err != nil {
			log.Fatal(err)
		} else {
			toDosResponse.ToDos = append(toDosResponse.ToDos, &toDo)
		}
	}

	json.NewEncoder(w).Encode(toDosResponse.ToDos)
}

func CreateToDo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var createRequest req.CreateRequest
	_ = json.NewDecoder(r.Body).Decode(&createRequest)

	params := mux.Vars(r)
	db := database.ConnectDatabase()
	_, err := db.Exec("INSERT INTO Todo (UserId, Title, Description, Status, Tag) VALUES (?, ?, ?, ?, ?)",
		params["userid"], createRequest.Title, createRequest.Description, createRequest.Status, createRequest.Tag)
	if err != nil {
		log.Fatal(err)
	}
}

func UpdateToDo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var updateRequest req.UpdateRequest
	_ = json.NewDecoder(r.Body).Decode(&updateRequest)

	params := mux.Vars(r)
	db := database.ConnectDatabase()
	_, err := db.Exec("UPDATE Todo SET Title = ?, Description = ?, Status = ?, Tag = ? WHERE UserId = ? AND Id = ?",
		updateRequest.Title, updateRequest.Description, updateRequest.Status, updateRequest.Tag, params["userid"], params["todoid"])
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteToDo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	db := database.ConnectDatabase()
	_, err := db.Exec("DELETE FROM Todo WHERE UserId = ? AND Id = ?",
		params["userid"], params["todoid"])
	if err != nil {
		log.Fatal(err)
	}
}
