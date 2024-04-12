package handler

import (
	"errors"
	user "parkify-BE/features/user"
	"parkify-BE/helper"

	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type controller struct {
	service user.UserService
}

func NewUserHandler(s user.UserService) user.UserController {
	return &controller{
		service: s,
	}
}

func (ct *controller) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input user.User
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}
		err = ct.service.Register(input)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") {
				code = http.StatusBadRequest
			}
			if strings.Contains(err.Error(), "Field validation for 'Nama'") {
				return c.JSON(http.StatusBadRequest,
					helper.ResponseFormat(http.StatusBadRequest, "Nama wajib diisi", nil))
			}
			if strings.Contains(err.Error(), "Field validation for 'Email'") {
				return c.JSON(http.StatusBadRequest,
					helper.ResponseFormat(http.StatusBadRequest, "Format harus berupa email", nil))
			}
			if strings.Contains(err.Error(), "Field validation for 'Password'") {
				return c.JSON(http.StatusBadRequest,
					helper.ResponseFormat(http.StatusBadRequest, "Password minimal 8", nil))
			}
			return c.JSON(code,
				helper.ResponseFormat(code, err.Error(), nil))
		}
		return c.JSON(http.StatusCreated,
			helper.ResponseFormat(http.StatusCreated, "selamat data sudah terdaftar", nil))
	}
}

func (ct *controller) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input LoginRequest
		err := c.Bind(&input)
		if err != nil {
			if strings.Contains(err.Error(), "unsupport") {
				log.Print("error unsupport: ", err.Error())
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, "format data tidak didukung", nil))
			}
			log.Print("error bad request input: ", err.Error())
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, "data yang dikirmkan tidak sesuai", nil))
		}

		var processData user.User
		processData.Email = input.Email
		processData.Password = input.Password

		result, token, err := ct.service.Login(processData)
		if err != nil {
			var code = http.StatusInternalServerError
			if strings.Contains(err.Error(), "validation") || strings.Contains(err.Error(), "cek kembali") {
				code = http.StatusBadRequest
			}
			if strings.Contains(err.Error(), "Field validation for 'Email'") {
				return c.JSON(http.StatusBadRequest,
					helper.ResponseFormat(http.StatusBadRequest, "Format harus berupa email", nil))
			}
			if strings.Contains(err.Error(), "Field validation for 'Password'") {
				return c.JSON(http.StatusBadRequest,
					helper.ResponseFormat(http.StatusBadRequest, "Password minimal 8", nil))
			}
			return c.JSON(code,
				helper.ResponseFormat(code, err.Error(), nil))
		}

		var responseData LoginResponse
		responseData.Fullname = result.Fullname
		responseData.Email = result.Email
		responseData.Token = token

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "berhasil login", responseData))

	}
}

func (ct *controller) Profile() echo.HandlerFunc {
	return func(c echo.Context) error {

		token, ok := c.Get("user").(*jwt.Token)
		defer func() {
			if err := recover(); err != nil {
				log.Println("error jwt creation:", err)

			}
		}()
		if !ok {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		result, err := ct.service.Profile(token)
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

		var response ProfileResponse
		response.ID = result.ID
		response.Fullname = result.Fullname
		response.Email = result.Email

		return c.JSON(http.StatusOK,
			helper.ResponseFormat(http.StatusOK, "berhasil mendapatkan data", response))
	}
}

func (ct *controller) UpdateProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		defer func() {
			if err := recover(); err != nil {
				log.Println("error jwt creation:", err)

			}
		}()
		if !ok {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		var input user.User
		err := c.Bind(&input)
		if err != nil {
			log.Println("error bind data:", err.Error())
			if strings.Contains(err.Error(), "unsupport") {
				return c.JSON(http.StatusUnsupportedMediaType,
					helper.ResponseFormat(http.StatusUnsupportedMediaType, helper.UserInputFormatError, nil))
			}
			return c.JSON(http.StatusBadRequest,
				helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		err = ct.service.UpdateProfile(token, input)
		if err != nil {
			log.Println("error update profile not found:", err.Error())
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusNotFound,
					helper.ResponseFormat(http.StatusNotFound, "data tidak ditemukan", nil))
			}
			if strings.Contains(err.Error(), "Field validation for 'Email'") {
				return c.JSON(http.StatusBadRequest,
					helper.ResponseFormat(http.StatusBadRequest, "Format harus berupa email", nil))
			}
			log.Println("error update profile forbidden:", err.Error())
			return c.JSON(http.StatusForbidden,
				helper.ResponseFormat(http.StatusForbidden, "Anda tidak diizinkan mengakses profil pengguna lain", nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "berhasil merubah profile", nil))
	}
}

func (ct *controller) DeleteAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		token, ok := c.Get("user").(*jwt.Token)
		defer func() {
			if err := recover(); err != nil {
				log.Println("error jwt creation:", err)

			}
		}()
		if !ok {
			return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, helper.UserInputError, nil))
		}

		err := ct.service.DeleteAccount(token)
		if err != nil {
			log.Println("error deleting account:", err.Error())
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusNotFound,
					helper.ResponseFormat(http.StatusNotFound, "data tidak ditemukan", nil))
			}
			// Jika terjadi kesalahan lain selain "record not found",
			// kembalikan respons forbidden
			log.Println("error Deleting profile:", err.Error())
			return c.JSON(http.StatusForbidden,
				helper.ResponseFormat(http.StatusForbidden, "Anda tidak diizinkan menghapus profil pengguna lain", nil))
		}

		return c.JSON(http.StatusCreated, helper.ResponseFormat(http.StatusCreated, "berhasil menghapus akun", nil))
	}
}
