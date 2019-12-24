package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/1zidorius/project-management-api/models"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello student!"))
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	u := &models.User{}
	err := dec.Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//resp := user.Create() //Create account
	fmt.Fprintf(w, "User: %+v", u)
}

func CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	dec := json.NewDecoder(r.Body)
	task := &models.Task{}
	err := dec.Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//resp := task.Create() //Create task
	task.CreatedOn = time.Now()
	t := time.Now()
	task.UpdatedOn = &t
	fmt.Fprintf(w, "User: %+v", task)
}

func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	dec := json.NewDecoder(r.Body)
	t := &models.Task{}
	err := dec.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
