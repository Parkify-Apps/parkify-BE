package routes

import (
	"parkify-BE/config"
	"parkify-BE/features/parking"
	parkingslot "parkify-BE/features/parkingslot"
	"parkify-BE/features/reservation"
	"parkify-BE/features/transaction"
	user "parkify-BE/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ct1 user.UserController, pc parking.ParkingController, psc parkingslot.ParkingSlotController, rc reservation.ReservationController, t transaction.TransactionController) {
	userRoute(c, ct1)
	parkingRoute(c, pc)
	parkingSlotRoute(c, psc)
	reservationRoute(c, rc)
	transactionRoute(c, t)
}

func userRoute(c *echo.Echo, ct1 user.UserController) {
	c.POST("/users", ct1.Add())
	c.POST("/login", ct1.Login())
	c.GET("/users", ct1.Profile(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.PUT("/users", ct1.UpdateProfile(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.DELETE("/users", ct1.DeleteAccount(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}

func parkingSlotRoute(c *echo.Echo, psc parkingslot.ParkingSlotController) {
	c.POST("/parkingslot", psc.Add(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.GET("/parkingslot", psc.AllParkingSlot(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.PUT("/parkingslot/:id", psc.Edit(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.DELETE("/parkingslot/:id", psc.Delete(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}

func parkingRoute(c *echo.Echo, pc parking.ParkingController) {
	c.POST("/parking", pc.PostParking(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.PUT("/parking/:id", pc.UpdateParking(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.GET("/parking/:id", pc.GetParking(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.GET("/parking", pc.GetAllParking(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}

func reservationRoute(c *echo.Echo, rc reservation.ReservationController) {
	c.POST("/reservation", rc.Create(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.GET("/reservation", rc.GetHistory(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.GET("/reservation/:id", rc.GetReservationInfo(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}

func transactionRoute(c *echo.Echo, t transaction.TransactionController) {
	c.POST("/transaction", t.Transaction(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.POST("/transaction/payment-callback", t.PaymentCallback(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.GET("/transaction/:id", t.Get(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}
