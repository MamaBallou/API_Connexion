package database

import (
	"fmt"
	"log"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Utilisateur struct {
	Email       string `gorm:"column:email;type:nvarchar(max);not null"`
	Password    string `gorm:"column:password;type:nvarchar(max);not null"`
	Permissions string `gorm:"column:permissions;type:nvarchar(max);not null"`
	Role 		string `gorm:"column:role;type:nvarchar(max);not null"`
}

func ConnectSQLServer() (*gorm.DB, error) {
	// Définissez les informations de connexion
    username := "Xav"
    password := "Cybersecurity1!"
    host := "10.0.2.105"
    port := 1433
	database := "user"

    // Construisez le DSN (Data Source Name)
    dsn := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
        host, username, password, port, database)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Impossible de se connecter à la base de données : " + err.Error())
    }

    return db, nil
}

func CloseSQLServer(db *gorm.DB) {
    sqlDB, err := db.DB()
    if err != nil {
        panic("Impossible d'obtenir la connexion SQL sous-jacente : " + err.Error())
    }

    // Fermer la connexion sous-jacente
    if err := sqlDB.Close(); err != nil {
        panic("Impossible de fermer la connexion SQL : " + err.Error())
    }
}

func AutoMigrateTables(db *gorm.DB) {
    // Utilisez AutoMigrate pour créer automatiquement les tables si elles n'existent pas
	log.Println("Création des tables")
    db.AutoMigrate(&Utilisateur{})
}