package main

import (
	"parkify-BE/config"
	"parkify-BE/features/user/data"
	"parkify-BE/features/user/handler"
	"parkify-BE/features/user/services"
	"parkify-BE/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)

	userData := data.New(db)
	userService := services.NewService(userData)
	userHandler := handler.NewUserHandler(userService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	routes.InitRoute(e, userHandler)
	e.Logger.Fatal(e.Start(":8000"))
}
