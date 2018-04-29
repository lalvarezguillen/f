package main

import (
	"github.com/labstack/echo"
)

func handleListUsers(c echo.Context) error {
	var users []User
	DB.Find(&users)
	return c.JSON(200, users)
}

func handleGetUser(c echo.Context) error {
	return nil
}

func handleCreateUser(c echo.Context) error {
	return nil
}

func handleUpdateUser(c echo.Context) error {
	return nil
}

func handleDeleteUser(c echo.Context) error {
	return nil
}
