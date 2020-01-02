package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/1zidorius/project-management-api/dao"
	"github.com/1zidorius/project-management-api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello world!"))
}

func GetUsersHandlers(w http.ResponseWriter, r *http.Request) {
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
	params := mux.Vars(r)
	payload := dao.GetAllUsers()
	for _, user := range payload {
		if params["id"] == "1" {
			//if uuid.Equal(user.Id, params["id"]){
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(user)
			if err != nil {
				http.Error(w, "", http.StatusBadRequest)
				return
			}
			return
		}
	}
	http.Error(w, "", http.StatusNotFound)
	return
}

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
	dao.CreateUser(u)
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


// TODO: test update handler
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		log.Println(err)
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1048576)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	u := models.User{}
	u.Id = id
	err = dec.Decode(&u)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = dao.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, err.Error())
	}
	fmt.Fprintf(w, "User: %+v", u)
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

func PrintRequestJson(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Fprintln(w, bodyString)
}
