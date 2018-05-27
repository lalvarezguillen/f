package main

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func handleListUsers(c echo.Context) error {
	var users []User
	DB.Find(&users)
	return c.JSON(200, users)
}

func handleGetUser(c echo.Context) error {
	u, ok := c.Get("user").(User)
	if !ok {
		log.Error("Unable to cast c.user to User type")
		return echo.NewHTTPError(500)
	}
	return c.JSON(200, u)
}

func handleCreateUser(c echo.Context) error {
	var u User
	if err := c.Bind(&u); err != nil {
		log.Error(err)
		return echo.NewHTTPError(400)
	}
	DB.Create(&u)
	return c.JSON(201, u)
}

func handleUpdateUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(400)
	}
	var updatedUser User
	if err = c.Bind(&updatedUser); err != nil {
		log.Error(err)
		return echo.NewHTTPError(400)
	}
	updatedUser.ID = uint(id)
	DB.Save(&updatedUser)
	return c.JSON(200, updatedUser)
}

func handleDeleteUser(c echo.Context) error {
	u, ok := c.Get("user").(User)
	if !ok {
		log.Error("Unable to cast c.user to User type")
		return echo.NewHTTPError(500)
	}
	DB.Delete(&u)
	return c.JSON(204, nil)
}
