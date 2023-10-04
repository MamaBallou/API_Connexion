package auth

import (
	"connexion/api/models"
	"connexion/api/security"
	"connexion/api/utils"
	"errors"
)

// Erreur pour des identifiants invalides.
var ErrInvalidCredentials = errors.New("Mot de passe incorrect.")

// Erreur pour un utilisateur non trouvé.
var ErrUserNotFound = errors.New("Email incorrect.")

// SignIn gère le processus de connexion de l'utilisateur.
func SignIn(email, password string) (string, error) {
	// Récupération de l'utilisateur en utilisant l'adresse e-mail.
	user, err := models.GetUserByEmail(email)
    if err != nil {
        return "", err
    }
    if user.Email == "" {
        return "", ErrUserNotFound 
    }

    // Vérification du mot de passe fourni.
    err = security.VerifyPassword(user.Password, password)
    if err != nil {
        return "", ErrInvalidCredentials
    }

	// Génération du token JWT.
    token, err := utils.GenerateJWT(user)
    if err != nil {
        return "", err
    }

    return token, nil
}