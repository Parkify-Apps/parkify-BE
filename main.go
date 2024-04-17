package main

import (
	// "log"
	"parkify-BE/config"
	pd "parkify-BE/features/parking/data"
	ph "parkify-BE/features/parking/handler"
	ps "parkify-BE/features/parking/services"
	parking_slot_data "parkify-BE/features/parkingslot/data"
	parking_slot_handler "parkify-BE/features/parkingslot/handler"
	parking_slot_services "parkify-BE/features/parkingslot/services"
	reservation_data "parkify-BE/features/reservation/data"
	reservation_handler "parkify-BE/features/reservation/handler"
	reservation_services "parkify-BE/features/reservation/services"
	tr_data "parkify-BE/features/transaction/data"
	tr_handler "parkify-BE/features/transaction/handler"
	tr_services "parkify-BE/features/transaction/services"
	"parkify-BE/features/user/data"
	"parkify-BE/features/user/handler"
	"parkify-BE/features/user/services"
	"parkify-BE/middlewares"
	"parkify-BE/routes"
	"parkify-BE/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitSQL(cfg)

	md := utils.NewMidtrans(cfg.MDKey)
	// log.Print(md)

	userData := data.New(db)
	userService := services.NewService(userData)
	userHandler := handler.NewUserHandler(userService)

	parkingData := pd.New(db)
	parkingService := ps.NewService(parkingData)
	parkingHandler := ph.NewHandler(parkingService)

	parkingSlotData := parking_slot_data.New(db)
	parkingSlotService := parking_slot_services.ParkingSlotService(parkingSlotData, middlewares.NewMidlewareJWT())
	parkingSlotHandler := parking_slot_handler.NewHandler(parkingSlotService)

	reservationData := reservation_data.New(db)
	reservationService := reservation_services.ReservationService(reservationData)
	reservationHandler := reservation_handler.NewHandler(reservationService)

	transactionData := tr_data.New(db)
	transactionService := tr_services.NewServices(transactionData, md)
	transactionHandler := tr_handler.NewHandler(transactionService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	routes.InitRoute(e, userHandler, parkingHandler, parkingSlotHandler, reservationHandler, transactionHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
