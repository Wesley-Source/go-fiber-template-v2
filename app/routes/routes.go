package routes

import (
	"go-fiber-template-v2/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return middleware.Redirect(c, "pages/index", "/")
}

func Register(c *fiber.Ctx) error {
	return middleware.Redirect(c, "pages/register", "/register")
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
