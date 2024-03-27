package user

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type UserController interface {
	Add() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	UpdateProfile() echo.HandlerFunc
	DeleteAccount() echo.HandlerFunc
}

type UserService interface {
	Register(newData User) error
	Login(loginData User) (User, string, error)
	Profile(token *jwt.Token) (User, error)
	UpdateProfile(userID int, token *jwt.Token, newData User) error
	DeleteAccount(userID uint, token *jwt.Token) error
}

type UserModel interface {
	AddUser(newData User) error
	Login(email string) (User, error)
	Profile(email string) (User, error)
	UpdateProfile(userID int, email string, data User) error
	GetUserByID(userID uint) (User, error)
	Delete(userID uint, email string) error
}

type User struct {
	gorm.Model
	Fullname string
	Email    string
	Password string
	Role     bool
}

type Login struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}

type Register struct {
	Fullname string `validate:"required,alpha"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,alphanum,min=8"`
	Role     bool   `validate:"required"`
}

type Update struct {
	Fullname string
	Password string
}
