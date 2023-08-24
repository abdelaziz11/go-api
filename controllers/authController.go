package controllers

import (
	"net/http"
	"github.com/labstack/echo"
	"api/models"
	"api/database"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c echo.Context) error {
	db := database.GetDB()
	username := c.FormValue("username")
	password := c.FormValue("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := models.NewUser(username, string(hashedPassword))
	_, err = db.Exec("INSERT INTO users (id, username, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", user.ID, user.Username, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	db := database.GetDB()
	input := new(LoginInput)
	if err := c.Bind(input); err != nil {
		return err
	}

	row := db.QueryRow("SELECT password FROM users WHERE username = $1", input.Username)
	storedPassword := ""
	err := row.Scan(&storedPassword)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(input.Password)); err != nil {
		return echo.ErrUnauthorized
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Successfully logged in",
	})
}
