package controllers

import (
	"net/http"
	"encoding/json"
	"connexion/api/utils"
	"connexion/api/models"
)

func PostUser(w http.ResponseWriter, r *http.Request){

	body := utils.BodyParser(r)
    var user models.User
    err := json.Unmarshal(body, &user)

    if err != nil {
        utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
        return
    }

    err = models.CreateUser(user)
    if err != nil {
        utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
        return
    }
    utils.ToJson(w, "User Created", http.StatusOK)
}