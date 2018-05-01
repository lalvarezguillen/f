package main

import (
	"github.com/labstack/echo"
)

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
	userID := c.Param("userID")
	var u User
	var userCount int
	DB.Where("id = ?", userID).First(&u).Count(&userCount)
	if userCount == 0 {
		return c.JSON(404, nil)
	}
	picID := c.Param("pictureID")
	var pic Picture
	var picsCount int
	DB.Where("id = ?", picID).First(&pic).Count(&picsCount)
	if picsCount == 0 {
		return c.JSON(404, nil)
	}
	return c.JSON(200, pic)
}

func handleCreateUserPicture(c echo.Context) error {
	var u User
	var userCount int
	userID := c.Param("userID")
	DB.Where("id = ?", userID).First(&u).Count(&userCount)
	if userCount == 0 {
		return c.JSON(404, nil)
	}
	var pic Picture
	err := c.Bind(&pic)
	if err != nil {
		return c.JSON(400, nil)
	}
	pic.UserID = u.ID
	DB.Create(&pic)
	return c.JSON(201, pic)
}

func handleUpdateUserPicture(c echo.Context) error {
	return nil
}

func handleDeleteUserPicture(c echo.Context) error {
	return nil
}
