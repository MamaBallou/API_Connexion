package security

import (
    "golang.org/x/crypto/bcrypt"
)

// Hash hache un mot de passe en utilisant l'algorithme bcrypt.
func Hash(password string) ([]byte, error) {
    return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword vérifie si le mot de passe haché correspond au mot de passe fourni.
func VerifyPassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}