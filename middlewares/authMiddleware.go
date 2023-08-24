package middlewares

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"api/models"
	"api/database"
	"github.com/google/uuid"

)

func Init() {
	// Initialize anything you need here
}

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authorization := c.Request().Header.Get("Authorization")

		if authorization == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please provide a token")
		}

		token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, echo.ErrUnauthorized
			}
			return []byte("secret"), nil
		})

		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			user := models.User{}
			d :=database.GetDB()
            
			d.QueryRow("SELECT id, username, created_at, updated_at FROM users WHERE id = $1", claims["id"]).Scan(&user.ID, &user.Username, &user.CreatedAt, &user.UpdatedAt)
            if user.ID == uuid.Nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "User not found")
			} 
			

			c.Set("user", user)
			return next(c)
		}

		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}
}

