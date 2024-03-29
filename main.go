package main

import (

	// "github.com/labstack/echo"
	"parkify-BE/config"
	"parkify-BE/features/user/data"
	"parkify-BE/features/user/handler"
	"parkify-BE/features/user/services"
	pd "parkify-BE/features/parking/data"
	ph "parkify-BE/features/parking/handler"
	ps "parkify-BE/features/parking/services"
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

	parkingData := pd.New(db)
	parkingService := ps.NewService(parkingData)
	parkingHandler := ph.NewHandler(parkingService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	routes.InitRoute(e, userHandler, parkingHandler)
	e.Logger.Fatal(e.Start(":8000"))
}
