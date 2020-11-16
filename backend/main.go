package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
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

func getTask(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(Tasks)
}

func postTask(w http.ResponseWriter, r *http.Request) {

	body, erro := ioutil.ReadAll(r.Body)
	var new_Task Task

	if erro != nil {
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

func deleteTask(w http.ResponseWriter, r *http.Request, URL_slices []string) {

	index := searchTask(w, r, URL_slices)
	if index < 0 {
		w.WriteHeader(http.StatusNoContent)
	} else {
		for index < len(Tasks)-1 {
			Tasks[index] = Tasks[index+1]
			index++
		}
		Tasks = Tasks[:len(Tasks)-1]
	}
}

func putTask(w http.ResponseWriter, r *http.Request, URL_slices []string) {
	index := searchTask(w, r, URL_slices)
	if index < 0 {
		postTask(w, r)
	} else {
		body, _ := ioutil.ReadAll(r.Body)
		var new_Task Task
		json.Unmarshal(body, &new_Task)
		new_Task.Id, _ = strconv.Atoi(URL_slices[3])
		Tasks[index] = new_Task
		json.NewEncoder(w).Encode(new_Task)
	}
}

func searchTask(w http.ResponseWriter, r *http.Request, URL_slices []string) int {

	w.Header().Set("Content-Type", "application/json")

	if len(URL_slices) < 4 {

		w.WriteHeader(http.StatusNoContent)
		return -1
	} else if len(URL_slices) == 4 || URL_slices[len(URL_slices)-1] == "" {

		id, _ := strconv.Atoi(URL_slices[3])

		for i := 0; i < len(Tasks); i++ {
			if Tasks[i].Id == id {

				w.WriteHeader(http.StatusFound)
				json.NewEncoder(w).Encode(Tasks[i])
				return i

			}
		}

		w.WriteHeader(http.StatusNoContent)
	}
	return -1
}

func http_Tasks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")

	if r.Method == "OPTIONS" {

		return
	}
	URL_slices := strings.Split(r.URL.Path, "/")

	if len(URL_slices) < 4 || len(URL_slices) == 4 && URL_slices[3] == "" {

		if r.Method == "GET" {
			getTask(w, r)

		} else if r.Method == "POST" {
			postTask(w, r)

		}
	} else if len(URL_slices) >= 4 || URL_slices[len(URL_slices)-1] == "" {
		if r.Method == "GET" {
			searchTask(w, r, URL_slices)

		} else if r.Method == "DELETE" {
			deleteTask(w, r, URL_slices)

		} else if r.Method == "PUT" {
			putTask(w, r, URL_slices)
		}
	} else {

		w.WriteHeader(http.StatusNotFound)
	}
}

func configureRoutes() {
	http.HandleFunc("/api/todos/", http_Tasks)
	http.HandleFunc("/api/todos", http_Tasks)

}

func configureServer() {
	configureRoutes()

	log.Fatal(http.ListenAndServe(":3003", nil))
}

func main() {
	configureServer()
}
