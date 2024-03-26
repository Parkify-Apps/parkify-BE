package user

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo"
)

type UserController interface {
	Add() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
}

type UserService interface {
	Register(newData User) error
	Login(loginData User) (User, string, error)
	Profile(token *jwt.Token, userID uint) (User, error)
}

type UserModel interface {
	AddUser(newData User) error
	Login(email string) (User, error)
	GetUserByID(userID uint) (User, error)
	GetLastUserID() (int, error)
}

type User struct {
	UserID   int
	Nama     string
	Email    string
	Password string
}

type Login struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type Register struct {
	UserID   int
	Nama     string `validate:"required,alpha"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,alphanum,min=8"`
}
