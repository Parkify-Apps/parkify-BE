package middlewares

import (
	"errors"
	"log"
	"parkify-BE/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(email string, role string) (string, error) {
	var data = jwt.MapClaims{}
	// custom data
	data["email"] = email
	data["role"] = role
	// mandatory data
	data["iat"] = time.Now().Unix()
	data["exp"] = time.Now().Add(time.Hour * 3).Unix()

	var proccessToken = jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	result, err := proccessToken.SignedString([]byte(config.JWTSECRET))

	if err != nil {
		defer func() {
			if err := recover(); err != nil {
				log.Println("error jwt creation:", err)

			}
		}()
		return "", errors.New("terjadi masalah pembuatan te")
	}

	return result, nil
}

func DecodeToken(token *jwt.Token) string {
	var result string
	var claim = token.Claims.(jwt.MapClaims)

	if val, found := claim["email"]; found {
		result = val.(string)
	}

	return result
}

func DecodeRole(token *jwt.Token) string {
	var result string
	var claim = token.Claims.(jwt.MapClaims)

	if val, found := claim["role"]; found {
		result = val.(string)
	}

	return result
}
