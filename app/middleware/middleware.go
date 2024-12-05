package middleware

import (
	"go-fiber-template-v2/app/database"
	"log"
	"os"
	"strings"
	"time"
	"unicode"

	"regexp"

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

	userID, err := GetSessionCookie(c)

	if c.Locals("user_id") != nil && err == nil {
		user := database.SearchUserById(c.Locals("user_id").(uint))

		data["IsAdmin"] = user.IsAdmin
		data["UserID"] = userID
		data["Name"] = user.Name
		data["Email"] = user.Email
	}

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
	userID, err := GetSessionCookie(c)
	if err != nil {
		log.Printf("Session error: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to authenticate user",
		})
	}

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

func UnknownAuth(c *fiber.Ctx) error {
	userID, err := GetSessionCookie(c)
	if err != nil || userID == nil {
		return c.Next()
	}
	c.Locals("user_id", userID.(uint))
	return c.Next()
}

func AdminAuth(c *fiber.Ctx) error {
	user := database.SearchUserById(c.Locals("user_id").(uint))
	if user.IsAdmin {
		return Redirect(c, "pages/admin", "/admin")
	}
	return Render(c, "pages/404")
}

// SetSessionCookie stores the user ID in the session
func SetSessionCookie(c *fiber.Ctx, id uint) {
	session, err := Session.Get(c)
	if err != nil {
		log.Println("Failed to get session.")
	}

	// Saves the user_id as a cookie in the user's browser
	session.Set("user_id", id)
	session.Save()
}

// GetSessionCookie retrieves the user ID from the session
func GetSessionCookie(c *fiber.Ctx) (interface{}, error) {
	session, err := Session.Get(c)
	if err != nil {
		return nil, err
	}
	return session.Get("user_id"), nil
}

// ClearSessionCookie removes the user ID from the session and clears the cookie
func ClearSessionCookie(c *fiber.Ctx) error {
	session, err := Session.Get(c)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	session.Delete("user_id")
	session.Save()
	return nil
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

// ValidateInput validates user input for registration
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidateInput(name, email, password string) []ValidationError {
	var errors []ValidationError

	// Validate name
	if len(name) < 3 {
		errors = append(errors, ValidationError{
			Field:   "name",
			Message: "Name must be at least 3 characters long",
		})
	}

	// Validate email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		errors = append(errors, ValidationError{
			Field:   "email",
			Message: "Invalid email format",
		})
	}

	// Validate password
	if len(password) < 8 {
		errors = append(errors, ValidationError{
			Field:   "password",
			Message: "Password must be at least 8 characters long",
		})
	}

	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSpecial := false
	specialChars := "!@#$%^&*"

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case strings.ContainsRune(specialChars, char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		errors = append(errors, ValidationError{
			Field:   "password",
			Message: "Password must contain at least one uppercase letter",
		})
	}
	if !hasLower {
		errors = append(errors, ValidationError{
			Field:   "password",
			Message: "Password must contain at least one lowercase letter",
		})
	}
	if !hasNumber {
		errors = append(errors, ValidationError{
			Field:   "password",
			Message: "Password must contain at least one number",
		})
	}
	if !hasSpecial {
		errors = append(errors, ValidationError{
			Field:   "password",
			Message: "Password must contain at least one special character (!@#$%^&*)",
		})
	}

	return errors
}

// RegisterUser handles user registration with input validation
func RegisterUser(c *fiber.Ctx) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	// Validate all inputs
	if errors := ValidateInput(name, email, password); len(errors) > 0 {
		// Return first error as string for HTMX
		return c.Status(fiber.StatusBadRequest).SendString(errors[0].Message)
	}

	// Hash password and create user
	hashedPassword := HashPassword(password)
	user := database.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	if err := database.Database.Create(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating user")
	}
	return nil
}
