package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func Data(c echo.Context) error {
	// Get team and member from the query string
	param1 := c.QueryParam("param1")
	param2 := c.QueryParam("param2")
	return c.String(http.StatusOK, "param1:"+param1+", param2:"+param2)
}
