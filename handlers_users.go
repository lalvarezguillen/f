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
	id := c.Param("id")
	var count int
	var u User
	DB.Where("id = ?", id).First(&u).Count(&count)
	if count == 0 {
		return c.JSON(404, nil)
	}
	return c.JSON(200, u)
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
	id := c.Param("id")
	var count int
	var u User
	DB.Where("id = ?", id).First(&u).Count(&count)
	if count == 0 {
		return c.JSON(404, nil)
	}
	var updatedUser User
	err := c.Bind(&updatedUser)
	if err != nil {
		return c.JSON(400, nil)
	}
	DB.Save(&updatedUser)
	return c.JSON(200, updatedUser)
}

func handleDeleteUser(c echo.Context) error {
	id := c.Param("id")
	var count int
	var u User
	DB.Where("id = ?", id).First(&u).Count(&count)
	if count == 0 {
		return c.JSON(404, nil)
	}
	DB.Delete(&u)
	return c.JSON(204, nil)
}
