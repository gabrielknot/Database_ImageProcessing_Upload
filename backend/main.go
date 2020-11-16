package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Task struct {
	Id          int       `json:"id,omitempty"`
	Description string    `json:"description,omitempty"`
	Done        string    `json:"done,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}

var Tasks []Task = []Task{
	Task{
		Id:          1,
		Description: "Capinar Mato",
		Done:        "false",
		CreatedAt:   time.Now(),
	},
	Task{
		Id:          2,
		Description: "Entender suas limitações",
		Done:        "false",
		CreatedAt:   time.Now(),
	},
	Task{
		Id:          3,
		Description: "Não agir se não tiver convicção",
		Done:        "false",
		CreatedAt:   time.Now(),
	},
	Task{
		Id:          4,
		Description: "Não agir se não tiver convicção",
		Done:        "false",
		CreatedAt:   time.Now(),
	},
}

func getDataBase(w http.ResponseWriter, r *http.Request) {
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

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "your-password"
	dbname   = "calhounio_demo"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	configureServer()
}
