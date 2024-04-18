package user

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	Type  string `json:"type"`
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if username != "adminTax" && password != "admin!" {
		return echo.ErrUnauthorized
	}

	claims := &jwtCustomClaims{
		"admin", true, "accessToken", jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 30)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	_, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "Login")
}
