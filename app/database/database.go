package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	IsAdmin  bool `gorm:"default:false"`
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

func UserExistsbyEmail(email string) bool {
	var user User
	result := Database.Where("email = ?", email).First(&user)
	return result.Error == nil // Retorna true se não houver erro (usuário encontrado)
}

func SearchUserById(id uint) User {
	var user User

	Database.Where("id = ?", id).First(&user)
	return user
}

func SearchUserByEmail(email string) User {
	var user User

	Database.Where("email = ?", email).First(&user)
	return user
}

func GetAllUsers() []User {
	var users []User
	Database.Find(&users)
	return users
}
