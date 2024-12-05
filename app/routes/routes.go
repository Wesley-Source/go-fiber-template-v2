package routes

import (
	"go-fiber-template-v2/app/database"
	"go-fiber-template-v2/app/middleware"
	"log"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return middleware.Redirect(c, "pages/index", "/")
}

func Register(c *fiber.Ctx) error {

	switch c.Method() {
	case "GET":
		return middleware.Redirect(c, "pages/register", "/register")

	case "POST":
		email := c.FormValue("email")
		if !database.UserExistsbyEmail(email) {
			err := middleware.RegisterUser(c)
			if err != nil {
				return c.SendString(err.Error())
			}
			return middleware.Redirect(c, "pages/login", "/login")
		}
	}
	return middleware.Redirect(c, "pages/404", "/404")
}

func Login(c *fiber.Ctx) error {

	switch c.Method() {
	case "GET":
		return middleware.Redirect(c, "pages/login", "/login")

	case "POST":
		email := c.FormValue("email")
		password := c.FormValue("password")
		if database.UserExistsbyEmail(email) {
			user := database.SearchUserByEmail(email)
			if middleware.ValidatePassword(user.Password, password) {
				middleware.SetSessionCookie(c, user.ID)
				return middleware.Redirect(c, "pages/index", "/")
			}
		}
	}

	return middleware.Redirect(c, "pages/login", "/login")
}

func UnknownRoute(c *fiber.Ctx) error {
	return middleware.Redirect(c, "pages/404", "/404")
}

func Logout(c *fiber.Ctx) error {

	switch c.Method() {
	case "GET":
		err := middleware.ClearSessionCookie(c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return middleware.Redirect(c, "pages/index", "/")
	}
	return UnknownRoute(c)
}

func Admin(c *fiber.Ctx) error {
	if middleware.AdminAuth(c) != nil {
		return UnknownRoute(c)
	}
	return middleware.Redirect(c, "pages/admin", "/admin")
}

func CheckEmail(c *fiber.Ctx) error {
	// Adicione mais logs para depuração
	log.Println("Checking email...")
	email := c.Query("email")
	log.Printf("Received email: %s", email)

	if email == "" {
		log.Println("Empty email")
		return c.SendString("")
	}

	exists := database.UserExistsbyEmail(email)
	log.Printf("Email exists: %v", exists)

	if exists {
		return c.SendString("Taken")
	}
	return c.SendString("Free")
}
