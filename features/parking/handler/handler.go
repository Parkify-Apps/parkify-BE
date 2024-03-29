package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"parkify-BE/features/parking"

	"parkify-BE/helper"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type controller struct {
	s parking.ParkingServices
}

func NewHandler(service parking.ParkingServices) parking.ParkingController {
	return &controller{
		s: service,
	}
}

func (ct *controller) PostParking() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		file, err := c.FormFile("imageloc")
		if err != nil && err != http.ErrMissingFile { // Check if error is not due to missing file
			log.Println("error form file: ", err.Error())
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "Invalid data! The data type must be images!", nil))
		}

		cld, err := cloudinary.NewFromURL("cloudinary://426244812151882:GBqN6L8Rm77iHHkPXiemVPP_e2Y@dlosajdpy")
		if err != nil {
			log.Print("error connect error: ", err.Error())
			return err
		}
		resp, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{})
		if err != nil {
			log.Print("error upload error: ", err.Error())
			return err
		}

		locName := c.FormValue("locname")
		location := c.FormValue("location")


		var inputProcess parking.Parking
		inputProcess.LocationName = locName
		inputProcess.Location = location
		if resp.SecureURL != "" { // add image only if URL is not empty
			inputProcess.ImageLoc = resp.SecureURL
		}
		if err := ct.s.PostParking(token, inputProcess); err != nil {
			log.Println("error update account:", err.Error())
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusNotFound,
					helper.ResponseFormat(http.StatusNotFound, "data tidak ditemukan", nil))
			}
			// Jika terjadi kesalahan lain selain "record not found",
			// kembalikan respons forbidden
			log.Println("error update profile:", err.Error())
			return c.JSON(http.StatusForbidden,
				helper.ResponseFormat(http.StatusForbidden, "Anda tidak diizinkan mengakses halaman ini", nil))
		}
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "Adding Parking Success", nil))

	}
}
