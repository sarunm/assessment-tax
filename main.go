package main

import (
	"github.com/labstack/echo/v4"
	"github.com/sarunm/assessment-tax.git/user"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootca1mp!")
	})

	e.POST("/login", user.Login)
	e.Logger.Fatal(e.Start(":8080"))
}
