package handlers

import (
	"encoding/json"
	"github.com/1zidorius/project-management-api/dao"
	"github.com/1zidorius/project-management-api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	u := models.User{}
	err := dec.Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err = dao.CreateUser(u)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(u)
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	payload := dao.GetAllUsers()
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(response))

}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	user, err := dao.GetUser(id)
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(user)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Println(err)
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	user := models.User{}
	err = dec.Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	response, err := dao.UpdateUser(id, user)
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	deletedCount, err := dao.DeleteUser(id)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	if deletedCount == 0 {
		http.Error(w, "", http.StatusNotFound)
		return
	}
}
