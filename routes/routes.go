package routes

import (
	"github.com/labstack/echo"
	"github.com/yourusername/yourprojectname/controllers"
	"github.com/yourusername/yourprojectname/middlewares"
)

func Init(e *echo.Echo) {
	// Public routes
	e.POST("/login", controllers.Login)

	// Restricted routes
	r := e.Group("/auth")
	r.Use(middlewares.IsAuthenticated)

	r.GET("/user", controllers.GetUser)
	r.PUT("/user", controllers.UpdateUser)
	r.DELETE("/user", controllers.DeleteUser)
}
