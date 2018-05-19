package main

import (
	"strconv"

	"github.com/labstack/echo"
)

func handleListUserPictures(c echo.Context) error {
	userID := c.Param("userID")
	var pics []Picture
	DB.Where("user_id = ?", userID).Find(&pics)
	return c.JSON(200, pics)
}

func handleGetUserPicture(c echo.Context) error {
	userID := c.Param("userID")
	picID := c.Param("pictureID")
	var pic Picture
	DB.Where("user_id = ? AND id = ?", userID, picID).First(&pic)
	if pic == (Picture{}) {
		return echo.NewHTTPError(404)
	}
	return c.JSON(200, pic)
}

func handleCreateUserPicture(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	var pic Picture
	err = c.Bind(&pic)
	if err != nil {
		return echo.NewHTTPError(400)
	}
	pic.UserID = uint(userID)
	DB.Create(&pic)
	return c.JSON(201, pic)
}

func handleUpdateUserPicture(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(400)
	}
	picID, err := strconv.ParseUint(c.Param("pictureID"), 10, 32)
	if err != nil {
		return echo.NewHTTPError(400)
	}
	var p Picture
	DB.Where("user_id = ? AND id = ?", userID, picID).First(&p)
	if p == (Picture{}) {
		return echo.NewHTTPError(404)
	}
	var pic Picture
	if err = c.Bind(&pic); err != nil {
		return echo.NewHTTPError(400)
	}
	pic.ID = uint(picID)
	DB.Save(&pic)
	return c.JSON(200, pic)
}

func handleDeleteUserPicture(c echo.Context) error {
	uID := c.Param("userID")
	pID := c.Param("pictureID")
	var p Picture

	DB.Where("id = ? AND user_id = ?", pID, uID).First(&p)
	if p == (Picture{}) {
		return echo.NewHTTPError(404)
	}
	DB.Delete(&p)
	return c.JSON(204, nil)
}
