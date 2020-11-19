package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"database/sql"
	"fmt"

	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

type Database struct {
	ID     int      `json:"id,omitempty"`
	Dbname string   `json:"dbname,omitempty"`
	images []string `json:"done,omitempty"`
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
	var err error

	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
		return
	}
	var errorOnCreate error

	_, errorOnCreate = db.Exec(
		"CREATE TABLE [IF NOT EXISTS] DATABASES (" +
			"ID serial PRIMARY KEY," +
			"Dbname VARCHAR ( 50 ) UNIQUE NOT NULL" +
			"images TEXT []" +
			")")

	if errorOnCreate != nil {
		panic(errorOnCreate)
		return
	}

	fmt.Println("Successfully connected!")
}

func getDataBase(w http.ResponseWriter, r *http.Request) {

	registers, errorOnGetRows := db.Query("SELECT ID, Dbname , images  FROM DATABASES")

	if errorOnGetRows != nil {
		panic(errorOnGetRows)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var databases []Database = make([]Database, 0)

	for registers.Next() {
		var database Database
		scanErorr := registers.Scan(&database.ID, &database.Dbname, &database.images)
		if scanErorr != nil {
			panic(scanErorr)
			continue
		}

		databases = append(databases, database)
	}

	closeRergistersError := registers.Close()

	if closeRergistersError != nil {
		panic(closeRergistersError)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(databases)
}

func postDataBase(w http.ResponseWriter, r *http.Request) {

	body, erro := ioutil.ReadAll(r.Body)
	var newDataBase Database

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

	registers := db.QueryRow("SELECT ID, Dbname , images  FROM DATABASES WHERE ID = ?", id)

	var database Database

	scanErorr := registers.Scan(&database.ID, &database.Dbname, &database.images)

	w.Header().Add("Content-Type", "application/json")
	if scanErorr != nil {
		panic(scanErorr)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	var mdifiedDatabase Database

	json.Unmarshal(body, &mdifiedDatabase)

	_, execError := db.Exec("UPDATE DATABASES SET Dbname = ?, images = ? WHERE ID = ?", modmdifiedDatabase.Dbname, mdifiedDatabase.images, id)
	if execErorr != nil {
		panic(execErorr)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(mdifiedDatabase)

}

func searchDataBase(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["databaseID"])

	registers := db.QueryRow("SELECT ID, Dbname , images  FROM DATABASES WHERE ID = ?", id)

	var database Database

	scanErorr := registers.Scan(&database.ID, &database.Dbname, &database.Images)

	w.Header().Add("Content-Type", "application/json")
	if scanErorr != nil {
		panic(scanErorr)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(dataBase)

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
