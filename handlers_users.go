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
	var u User
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(400, nil)
	}
	DB.Create(&u)
	return c.JSON(201, u)
}

func handleUpdateUser(c echo.Context) error {
	return nil
}

func handleDeleteUser(c echo.Context) error {
	return nil
}
