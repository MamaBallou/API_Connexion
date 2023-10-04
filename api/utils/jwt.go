package utils

import (
    "time"
    "connexion/config"
    "connexion/api/models"
    jwt "github.com/dgrijalva/jwt-go"
)

// Clé secrète pour la génération et la validation des tokens JWT.
var secretKey = config.JwtSecretKey

var application_type = "Consultation des freezebee"
var application_version = "1.0.0"

// Génère un token JWT en utilisant les informations de l'utilisateur.
func GenerateJWT(user models.User) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    claims["application_type"] = application_type
    claims["application_version"] = application_version
    claims["user_permission"] = user.Permissions
    claims["user_role"] = user.Role
    claims["exp"] = time.Now().Add(time.Minute * 120).Unix()

    // Signe le token avec la clé secrète et renvoie le token signé.
    signedToken, err := token.SignedString(secretKey)
    if err != nil {
        return "", err // Gérer l'erreur ici
    }
    
    return signedToken, nil
}