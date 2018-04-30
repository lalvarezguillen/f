package main

import "github.com/labstack/echo"

func handleListUserPictures(c echo.Context) error {
	userID := c.Param("userID")
	var u User
	var userCount int
	DB.Where("id = ?", userID).First(&u).Count(&userCount)
	if userCount == 0 {
		return c.JSON(404, nil)
	}
	var pics []Picture
	DB.Model(&u).Related(&pics)
	return c.JSON(200, pics)
}

func handleGetUserPicture(c echo.Context) error {
	return nil
}

func handleCreateUserPicture(c echo.Context) error {
	return nil
}

func handleUpdateUserPicture(c echo.Context) error {
	return nil
}

func handleDeleteUserPicture(c echo.Context) error {
	return nil
}
