package routes

import (
	"go-fiber-template-v2/app/database"
	"go-fiber-template-v2/app/middleware"

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
		if !database.UserExists(email) {
			middleware.RegisterUser(c)
			return middleware.Redirect(c, "pages/login", "/login")
		}
	}
	return middleware.Redirect(c, "pages/404", "/404")
}

func Login(c *fiber.Ctx) error {
	return middleware.Redirect(c, "pages/login", "/login")
}

func UnknownRoute(c *fiber.Ctx) error {
	return middleware.Redirect(c, "pages/404", "/404")
}

func Logout(c *fiber.Ctx) error {
	middleware.ClearSessionCookie(c)
	return middleware.Redirect(c, "pages/index", "/")
}