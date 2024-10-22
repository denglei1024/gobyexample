package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetToDoItems(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("select id,name,is_complete from t_todo_item")
	if err != nil {
		panic(err.Error())
	}
	var todos []Todo
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.Id, &todo.Name, &todo.IsComplete)
		if err != nil {
			log.Printf("Error happened in scan row. Err: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}
	jsonResp, err := json.Marshal(todos)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func GetToDoItemInCompleted(w http.ResponseWriter, r *http.Request) {
	rows, err := DB.Query("select id,name,is_complete from t_todo_item where is_complete = true")
	if err != nil {
		panic(err.Error())
	}
	var todos []Todo
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.Id, &todo.Name, &todo.IsComplete)
		if err != nil {
			log.Printf("Error happened in scan row. Err: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		todos = append(todos, todo)
	}
	jsonResp, err := json.Marshal(todos)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func CreateToDoItem(w http.ResponseWriter, r *http.Request) {
	// convert todo struct from body
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error happened in JSON unmarshal. Err: %s", err)
		return
	}
	// insert into db
	err = DB.QueryRow("insert into t_todo_item (name, is_complete) values ($1, $2) RETURNING id", todo.Name, false).Scan(&todo.Id)
	if err != nil {
		log.Printf("Error happened in insert into db. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// return created todo
	jsonResp, _ := json.Marshal(todo)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func GetToDoItemById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	row := DB.QueryRow("select id,name, is_complete from t_todo_item where id = $1", id)
	var todo Todo
	row.Scan(&todo.Id, &todo.Name, &todo.IsComplete)
	jsonResp, err := json.Marshal(todo)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func DeleteToDoItem(w http.ResponseWriter, r *http.Request) {
	// check id from url
	vars := mux.Vars(r)
	// convert id to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error happened in convert id to int. Err: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// if id less then 0, return 400
	if condition := id < 0; condition {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// delete from db
	_, err = DB.Exec("delete from t_todo_item where id = $1", id)
	if err != nil {
		log.Printf("Error happened in delete from db. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateToDoItem(w http.ResponseWriter, r *http.Request) {
	// check id from url
	vars := mux.Vars(r)
	// convert id to int
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Error happened in convert id to int. Err: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// if id less then 0, return 400
	if condition := id < 0; condition {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// convert todo struct from body
	var todo Todo
	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Error happened in JSON unmarshal. Err: %s", err)
	}
	todo.Id = id
	// update db
	_, err = DB.Exec("update t_todo_item set name = $1, is_complete = $2 where id = $3", todo.Name, todo.IsComplete, id)
	if err != nil {
		log.Printf("Error happened in update db. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
