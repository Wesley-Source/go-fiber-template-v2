package middleware

import (
	"go-fiber-template-v2/app/database"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3/v2"
	"github.com/gofiber/utils"
	"golang.org/x/crypto/bcrypt"
)

var Session *session.Store

func Render(c *fiber.Ctx, view string, partial ...bool) error {

	data := fiber.Map{}

	data["Title"] = os.Getenv("TITLE")

	if partial != nil && partial[0] {
		return c.Render(view, data)
	}
	return c.Render(view, data, "layouts/main")
}

func Redirect(c *fiber.Ctx, view, route string) error {
	if c.Get("HX-Request") == "true" {
		c.Set("HX-Redirect", route)
		return c.SendStatus(fiber.StatusOK)
	}

	return Render(c, view)
}

func ConnectSessionsDB() {
	storage := sqlite3.New(sqlite3.Config{
		Table:    "session_storage",
		Database: "./app/database/sessions.db",
	})
	Session = session.New(session.Config{
		Storage:        storage,
		Expiration:     24 * time.Hour,
		KeyLookup:      "cookie:session_id",
		CookieSecure:   true,
		CookieHTTPOnly: true,
		CookieSameSite: "Strict",
		KeyGenerator:   utils.UUID,
	})
}

func Auth(c *fiber.Ctx) error {
	session, err := Session.Get(c)
	if err != nil {
		log.Printf("Session error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to authenticate user",
		})
	}

	userID := session.Get("user_id")

	switch c.Path() {

	// If the user is logged in, redirect them to the index page instead of /login or /register
	case "/login", "/register":

		// Redirects the user to the index page if they're already logged in
		if userID != nil {
			return Redirect(c, "pages/index", "/")
		} else {
			return c.Next()
		}
	}

	if userID == nil {
		return Redirect(c, "pages/index", "/")
	}
	c.Locals("user_id", userID.(uint))
	return c.Next()
}

// GetSessionCookie retrieves the user ID from the session
func GetSessionCookie(c *fiber.Ctx) interface{} {
	session, err := Session.Get(c)
	if err != nil {
		log.Println("Failed to get session.")
	}

	return session.Get("user_id")
}

// ClearSessionCookie removes the user ID from the session and clears the cookie
func ClearSessionCookie(c *fiber.Ctx) {
	session, err := Session.Get(c)
	if err != nil {
		log.Println("Failed to get session.")
	} else {
		session.Delete("user_id")
		session.Save()
	}
	c.ClearCookie("user_id")
}

// HashPassword generates a secure hash of the provided password
func HashPassword(password string) string {
	// Returns a hashed and salted password
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashBytes)
}

// ValidatePassword checks if a password matches its hashed version
func ValidatePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func RegisterUser(c *fiber.Ctx) {
	user := database.User{
		Name:     c.FormValue("name"),
		Email:    c.FormValue("email"),
		Password: HashPassword(c.FormValue("password")),
	}
	database.Database.Create(&user)
}
