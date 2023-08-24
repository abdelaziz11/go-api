package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"api/config"
	"api/database"
	"api/middlewares"
	"api/routes"
)



func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Database connection
	config.Connect()
	database.Init()

	// Routes
	routes.Init(e)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}

func init() {
	middlewares.Init()
}
