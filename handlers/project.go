package handlers

import (
	"encoding/json"
	"github.com/1zidorius/project-management-api/dao"
	"github.com/1zidorius/project-management-api/models"
	"net/http"
)

func CreateProjectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	task := models.Task{}
	err := dec.Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	task, err = dao.CreateTask(task)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

