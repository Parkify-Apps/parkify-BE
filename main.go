package main

import (
	"parkify-BE/config"
	pd "parkify-BE/features/parking/data"
	ph "parkify-BE/features/parking/handler"
	ps "parkify-BE/features/parking/services"
	parking_slot_data "parkify-BE/features/parkingslot/data"
	parking_slot_handler "parkify-BE/features/parkingslot/handler"
	parking_slot_services "parkify-BE/features/parkingslot/services"
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

	parkingData := pd.New(db)
	parkingService := ps.NewService(parkingData)
	parkingHandler := ph.NewHandler(parkingService)

	parkingSlotData := parking_slot_data.New(db)
	parkingSlotService := parking_slot_services.ParkingSlotService(parkingSlotData)
	parkingSlotHandler := parking_slot_handler.NewHandler(parkingSlotService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	routes.InitRoute(e, userHandler, parkingHandler, parkingSlotHandler)
  
	e.Logger.Fatal(e.Start(":8000"))
}
