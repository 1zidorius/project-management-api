package handlers

import (
	"encoding/json"
	"github.com/project-management-api/models"
	"github.com/project-management-api/utils"
	"net/http"
)

func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user) //decode the request body into struct and failed if any error occur
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}

	//resp := user.Create() //Create account
	resp := utils.Message(true, "success")
	utils.Respond(w, resp)
}