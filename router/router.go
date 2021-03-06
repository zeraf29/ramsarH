package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"

	handler "github.com/zeraf29/ramsarH/handler"
)

//Router Function
func Router() *echo.Echo {
	//echo.New() 를 통해 *Echo 리턴 받음
	e := echo.New()

	//Debug setting
	e.Debug = true

	//echo Middleware func
	e.Use(middleware.Logger())                             // setting Logger
	e.Use(middleware.Recover())                            // Recover from panic anyware in the chain
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{ //CORS Config
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	//Health check
	e.GET("/healthy", func(c echo.Context) error {
		return c.String(http.StatusOK, "Healthy!")
	})

	//Router List
	data := e.Group("/data")
	{
		data.GET("/api", handler.Data)
	}

	return e
}
