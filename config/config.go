package config

import (
	"fmt"
	"os"
	parking "parkify-BE/features/parking/data"
	parkingslot "parkify-BE/features/parkingslot/data"
	reservation "parkify-BE/features/reservation/data"
	"parkify-BE/features/transaction"
	user "parkify-BE/features/user/data"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var JWTSECRET = ""

type AppConfig struct {
	DBUsername string
	DBPassword string
	DBPort     string
	DBHost     string
	DBName     string
	MDKey      string
}

func AssignEnv(c AppConfig) (AppConfig, bool) {
	var missing = false

	if val, found := os.LookupEnv("DBUsername"); found {
		c.DBUsername = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("DBPassword"); found {
		c.DBPassword = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("DBPort"); found {
		c.DBPort = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("DBHost"); found {
		c.DBHost = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("DBName"); found {
		c.DBName = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("JWT_SECRET"); found {
		JWTSECRET = val
	} else {
		missing = true
	}
	if val, found := os.LookupEnv("MDKey"); found {
		c.MDKey = val
	} else {
		missing = true
	}
	// log.Print(c.MDKey)
	return c, missing
}

func InitConfig() AppConfig {
	var result AppConfig
	var missing = false
	result, missing = AssignEnv(result)
	if missing {
		godotenv.Load(".env")
		result, _ = AssignEnv(result)
	}

	return result
}

func InitSQL(c AppConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.DBUsername, c.DBPassword, c.DBHost, c.DBPort, c.DBName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("terjadi error", err.Error())
		return nil
	}

	db.AutoMigrate(&user.User{}, &parkingslot.ParkingSlot{}, &parking.Parking{}, &reservation.Reservation{}, &transaction.Transaction{})

	return db
}
