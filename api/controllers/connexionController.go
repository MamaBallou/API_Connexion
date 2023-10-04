package controllers

import (
	"net/http"
	"encoding/json"
	"connexion/api/utils"
	"connexion/api/models"
	"connexion/api/auth"
	"connexion/config"
)

type Response struct{
	Token string `json:"token"`
}

var secretKey = config.ApiKey

func Connexion(w http.ResponseWriter, r *http.Request){
	apiKey := utils.ExtractApiKey(r)

	if apiKey != secretKey{
		utils.ToJson(w, "Wrong Api Key", http.StatusUnauthorized)
        return
	}

	body := utils.BodyParser(r)
    var user models.User
    err := json.Unmarshal(body, &user)

    if err != nil {
        utils.ToJson(w, err.Error(), http.StatusUnprocessableEntity)
        return
    }

	token, err := auth.SignIn(user.Email, user.Password)
    if err != nil {
        utils.ToJson(w, err.Error(), http.StatusUnauthorized)
        return
    }

	response := Response {
		Token: token,
	}

	utils.ToJson(w, response, http.StatusOK)
}