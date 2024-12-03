package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}

var Database *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("./app/database/models.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&User{}) // Include models as you create them

	Database = db // Moving the variable to the global scope
}

func UserExists(email string) bool {
	var user User
	result := Database.Where("email = ?", email).First(&user)
	return result.Error == nil // Retorna true se não houver erro (usuário encontrado)
}
