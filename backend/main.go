package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"database/sql"
	"fmt"

	"github.com/gorilla/mux"

	"github.com/rs/cors"

	_ "github.com/lib/pq"
)

type Database struct {
	ID     int      `json:"id,omitempty"`
	Dbname string   `json:"dbname,omitempty"`
	Images []string `json:"images,omitempty"`
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
		"CREATE TABLE DATABASES (" +
			"ID serial PRIMARY KEY," +
			"Dbname VARCHAR ( 50 ) UNIQUE NOT NULL," +
			"images TEXT []" +
			")")

	if errorOnCreate != nil {
		_, errorOnGetRows := db.Query("SELECT ID, Dbname , images  FROM DATABASES")

		if errorOnGetRows != nil {
			panic(errorOnCreate)
			return
		}
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
		scanErorr := registers.Scan(&database.ID, &database.Dbname, &database.Images)
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

	if erro != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var newDataBase Database

	json.Unmarshal(body, &newDataBase)

	_, execError := db.Exec("INSERT INTO DATABASES (Dbname, images) VALUES (?, ?)", newDataBase.Dbname, newDataBase.Images)

	if execError != nil {
		panic(execError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newDataBase)
}

func deleteDataBase(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["databaseID"])

	registers := db.QueryRow("SELECT ID FROM DATABASES WHERE ID = ?", id)

	var database Database

	scanErorr := registers.Scan(&database.ID, &database.Dbname, &database.Images)

	w.Header().Add("Content-Type", "application/json")
	if scanErorr != nil {
		panic(scanErorr)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	_, execError := db.Exec("DELETE FROM DATABASES WHERE ID = ?", database.Images, id)

	if execError != nil {
		panic(execError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func putDataBase(w http.ResponseWriter, r *http.Request) {

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

	body, _ := ioutil.ReadAll(r.Body)
	var mdifiedDatabase Database

	json.Unmarshal(body, &mdifiedDatabase)

	_, execError := db.Exec("UPDATE DATABASES SET Dbname = ?, images = ? WHERE ID = ?", mdifiedDatabase.Dbname, mdifiedDatabase.Images, id)
	if execError != nil {
		panic(execError)
		w.WriteHeader(http.StatusInternalServerError)
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
	json.NewEncoder(w).Encode(database)

}

func configureServer() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/databases/{databaseID}", searchDataBase).Methods("GET")
	router.HandleFunc("/api/databases", getDataBase).Methods("GET")
	router.HandleFunc("/api/databases", postDataBase).Methods("POST")
	router.HandleFunc("/api/databases/{databaseID}", putDataBase).Methods("PUT")
	router.HandleFunc("/api/databases/{databaseID}", deleteDataBase).Methods("DELETE")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(":3003", handler))
}

func main() {
	databaseConnection()
	configureServer()
}
