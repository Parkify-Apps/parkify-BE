package routes

import (
	"parkify-BE/config"
	"parkify-BE/features/parking"
	parkingslot "parkify-BE/features/parkingslot"
	user "parkify-BE/features/user"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ct1 user.UserController, pc parking.ParkingController, psc parkingslot.ParkingSlotController) {
	userRoute(c, ct1)
	parkingRoute(c, pc)
	parkingSlotRoute(c, psc)
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
	c.PUT("/parkingslot/:parkingSlotID", psc.Edit(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.DELETE("/parkingslot/:parkingSlotID", psc.Delete(), echojwt.WithConfig(echojwt.Config{
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

func parkingSlotRoute(c *echo.Echo, psc parkingslot.ParkingSlotController) {
	c.POST("/parkingslot", psc.Add(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.GET("/parkingslot", psc.AllParkingSlot(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.PUT("/parkingslot/:parkingSlotID", psc.Edit(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
	c.DELETE("/parkingslot/:parkingSlotID", psc.Delete(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(config.JWTSECRET),
	}))
}
