package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	ID     int      `json:"id,omitempty"`
	Dbname string   `json:"dbname,omitempty"`
	Images []string `json:"done,omitempty"`
}

const (
	host     = "localhost"
	port     = 3001
	user     = "docker"
	password = "docker"
	dbname   = "docker"
)
var db *sql.DB

func databaseConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err = sql.ErrConnDone
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	_, errorOnCreate := db.Exec(
		"CREATE TABLE [IF NOT EXISTS] DATABASES ("+
		"ID serial PRIMARY KEY,"+
		"Dbname VARCHAR ( 50 ) UNIQUE NOT NULL"+
		"images TEXT []"+
		")")
	if errorOnCreate != nil {
		panic(errorOnCreate)
	}

	fmt.Println("Successfully connected!")
}




func getDataBase(w http.ResponseWriter, r *http.Request) {
	registers, errorOnCreate := db.Query('SELECT * FROM DATABASES')

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Tasks)
}

func postDataBase(w http.ResponseWriter, r *http.Request) {

	body, erro := ioutil.ReadAll(r.Body)
	var new_Task Task

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		return

	} else {
		w.WriteHeader(http.StatusCreated)

		json.Unmarshal(body, &new_Task)
		new_Task.Id = Tasks[len(Tasks)-1].Id + 1
		new_Task.Done = "false"
		new_Task.CreatedAt = time.Now()
		Tasks = append(Tasks, new_Task)

	}

	json.NewEncoder(w).Encode(new_Task)
}

func deleteDataBase(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["databaseID"])
	if id <= 0 {
		w.WriteHeader(http.StatusNoContent)
	} else {
		for id < len(Tasks)-1 {
			Tasks[id] = Tasks[id+1]
			id++
		}
		Tasks = Tasks[:len(Tasks)-1]
	}
}

func putDataBase(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["databaseID"])
	if id <= 0 {
		postDataBase(w, r)
	} else {
		body, _ := ioutil.ReadAll(r.Body)
		var new_Task Task
		json.Unmarshal(body, &new_Task)
		new_Task.Id = id
		Tasks[id] = new_Task
		json.NewEncoder(w).Encode(new_Task)
	}
}

func searchDataBase(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["databaseID"])

	for _, dataBase := range Tasks {
		if dataBase.Id == id {

			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(dataBase)
			return
		}
	}
	w.WriteHeader(http.StatusNoContent)

}

func configureRoutes(router *mux.Router) {
	router.HandleFunc("/api/databases/{databaseID}", searchDataBase).Methods("GET")
	router.HandleFunc("/api/databases", getDataBase).Methods("GET")
	router.HandleFunc("/api/databases", postDataBase).Methods("POST")
	router.HandleFunc("/api/databases/{databaseID}", putDataBase).Methods("PUT")
	router.HandleFunc("/api/databases/{databaseID}", deleteDataBase).Methods("DELETE")
}

func configureServer() {
	router := mux.NewRouter().StrictSlash(true)

	configureRoutes(router)

	log.Fatal(http.ListenAndServe(":3003", router))
}

func main() {
	databaseConnection()
	configureServer()
}
