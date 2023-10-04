package models

import (
	"connexion/api/database"
	"connexion/api/security"
	"errors"
	"fmt"
)

type User struct {
	Email       string `gorm:"size:60;not null" json:"Email"`
	Password    string `gorm:"size:1000;" json:"Password"`
	Permissions string `gorm:"size:60;" json:"Permissions"`
	Role 		string `gorm:"size:60;" json:"Role"`
}

/*func GetUserByEmail(email string) (User, error) {

	// Hash du mot de passe
    hashedPassword, _ := security.Hash("azer")

	user := User{
		Email: email,
		Password: string(hashedPassword),
	}

	return user, nil
}*/

func GetUserByEmail(email string) (User, error) {
	db, err := database.ConnectSQLServer()
	if err != nil {
		return User{}, errors.New(err.Error())
	}
	defer database.CloseSQLServer(db)

	var user User

	// Recherchez l'utilisateur dans la table "user_dbs" en utilisant l'e-mail
	db.Table("dbo.utilisateurs").Where("email = ?", email).First(&user)

	return user, nil
}

func CreateUser(user User) error {
	db, err := database.ConnectSQLServer()
    if err != nil {
        return errors.New(err.Error())
    }
    defer database.CloseSQLServer(db)

	hashedPassword, err := security.Hash(user.Password)
    if err != nil {
        return errors.New(err.Error())
    }
    user.Password = string(hashedPassword)

	user_db := database.Utilisateur{
		Email: user.Email,
		Password: user.Password,
		Permissions: user.Permissions,
		Role: user.Role,
	}

	rs := db.Create(&user_db)
    if rs.Error != nil {
        return errors.New(fmt.Sprintf("Failed to create user:", rs.Error))
    }

	return nil
}
