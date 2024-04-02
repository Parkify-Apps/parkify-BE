package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"parkify-BE/features/parking"
	"strconv"
	"strings"

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

		city := c.FormValue("city")
		location := c.FormValue("location")

		var inputProcess parking.Parking
		inputProcess.City = city
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
			if err != nil {
				return c.JSON(http.StatusInternalServerError,
					helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
			}
		}
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "Adding Parking Success", nil))

	}
}

func (ct *controller) UpdateParking() echo.HandlerFunc {
	return func(c echo.Context) error {
		parkingID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		// Retrieve the uploaded file from the request.
		file, err := c.FormFile("imageloc")
		if err != nil && err != http.ErrMissingFile { // Check if error is not due to missing file
			log.Println("error form file: ", err.Error())
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "Invalid data! The data type must be images!", nil))
		}

		if file != nil { // Check if file not nil
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

			location := c.FormValue("location")
			city := c.FormValue("city")

			var updateProcess parking.Parking
			updateProcess.Location = location
			updateProcess.City = city

			// log.Print(resp.SecureURL)
			// var updateProcess user.User
			if resp.SecureURL != "" { // Update picture only if URL is not empty
				updateProcess.ImageLoc = resp.SecureURL
			}
			if err := ct.s.UpdateParking(parkingID, token, updateProcess); err != nil {
				log.Println("error update account:", err.Error())
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return c.JSON(http.StatusNotFound,
						helper.ResponseFormat(http.StatusNotFound, "data tidak ditemukan", nil))
				}
				// Jika terjadi kesalahan lain selain "record not found",
				// kembalikan respons forbidden
				log.Println("error update profile:", err.Error())
				return c.JSON(http.StatusForbidden,
					helper.ResponseFormat(http.StatusForbidden, "Anda tidak diizinkan mengakses profil pengguna lain", nil))
			}
			return c.JSON(http.StatusOK,
				helper.ResponseFormat(http.StatusOK, "Update Profile Success", nil))
		} else if file == nil {
			res, err := ct.s.GetPicture(parkingID)
			if err != nil {
				return c.JSON(http.StatusInternalServerError,
					helper.ResponseFormat(http.StatusInternalServerError, helper.ServerGeneralError, nil))
			}

			location := c.FormValue("location")
			city := c.FormValue("city")

			var updateProcess parking.Parking
			updateProcess.Location = location
			updateProcess.City = city
			updateProcess.City = res.ImageLoc

			if err := ct.s.UpdateParking(parkingID, token, updateProcess); err != nil {
				log.Println("error update account:", err.Error())
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return c.JSON(http.StatusNotFound,
						helper.ResponseFormat(http.StatusNotFound, "data tidak ditemukan", nil))
				}
				// Jika terjadi kesalahan lain selain "record not found",
				// kembalikan respons forbidden
				log.Println("error update profile:", err.Error())
				return c.JSON(http.StatusForbidden,
					helper.ResponseFormat(http.StatusForbidden, "Anda tidak diizinkan mengakses profil pengguna lain", nil))
			}

			return c.JSON(http.StatusOK,
				helper.ResponseFormat(http.StatusOK, "Update Profile Success", nil))

		}
		return nil
	}
}

func (ct *controller) GetParking() echo.HandlerFunc {
	return func(c echo.Context) error {
		parkingID, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			log.Println("error param:", err.Error())
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		token, ok := c.Get("user").(*jwt.Token)
		if !ok {
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		result, err := ct.s.GetParking(token, uint(parkingID))
		if err != nil {
			// Jika tidak ada data profil yang ditemukan, kembalikan respons "not found"
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusNotFound,
					helper.ResponseFormat(http.StatusNotFound, "data tidak ditemukan", nil))
			}
			// Jika terjadi kesalahan lain selain "record not found",
			// kembalikan respons forbidden
			log.Println("error getting profile:", err.Error())
			return c.JSON(http.StatusForbidden,
				helper.ResponseFormat(http.StatusForbidden, "Anda tidak diizinkan mengakses profil pengguna lain", nil))
		}

		var response ParkingResponse
		response.Location = result.Location
		response.City = result.City
		response.ImageLoc = result.ImageLoc
		// response.ParkingSlot = result.ParkingSlot

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "berhasil mendapatkan data", response))
	}
}

func (ct *controller) GetAllParking() echo.HandlerFunc {
	return func(c echo.Context) error {
		var parkingID, err = strconv.Atoi(c.QueryParam("id"))
		result, err := ct.s.GetAllParking(parkingID)
		if err != nil {
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				return c.JSON(http.StatusInternalServerError,
					helper.ResponseFormat(http.StatusInternalServerError, err.Error(), nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, err.Error(), nil))
		}
		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "successfully get parking", result))
	}
}
