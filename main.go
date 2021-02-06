package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Here is RamsarH Project!!")
	})
	e.Logger.Fatal(e.Start(":9091")) //port:9091
}

