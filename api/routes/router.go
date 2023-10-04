package routes

import (
	"connexion/api/controllers"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/mux"
)

// Crée un nouveau routeur en utilisant le package mux.
func NewRouter() *mux.Router {
    r := mux.NewRouter().StrictSlash(true)

    // Ajoute des middlewares globaux.
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)

    // Charge la gestion des en-têtes CORS.
    LoadCors(r)

    // Configure les routes pour diverses actions avec leurs contrôleurs associés.

	r.HandleFunc("/connexion", controllers.Connexion).Methods("POST")
    //r.HandleFunc("/user", controllers.PostUser).Methods("POST")

	return r
}