package user

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type UserController interface {
	Add() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
}

type UserService interface {
	Register(newData User) error
	Login(loginData User) (User, string, error)
	Profile(token *jwt.Token) (User, error)
}

type UserModel interface {
	AddUser(newData User) error
	Login(email string) (User, error)
	Profile(email string) (User, error)
}

type User struct {
	ID       uint
	Fullname string
	Email    string
	Password string
}

type Login struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type Register struct {
	Fullname string `validate:"required,alpha"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,alphanum,min=8"`
}
