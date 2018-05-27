package main

import "github.com/labstack/echo"

func handleGetAPIInfo(c echo.Context) error {
	return c.JSON(200, "F API v1")
}
