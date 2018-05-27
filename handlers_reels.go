package main

import (
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func handleListReels(c echo.Context) error {
	userID := c.Param("userID")
	var reels []Reel
	DB.Where("user_id = ?", userID).Find(&reels)
	return c.JSON(200, reels)
}

func handleGetReel(c echo.Context) error {
	userID := c.Param("userID")
	reelID := c.Param("reelID")
	var r Reel
	DB.Where("user_id = ? AND id = ?", userID, reelID).First(&r)
	if r.ID == 0 {
		return echo.NewHTTPError(404)
	}
	return c.JSON(200, r)
}

func handleCreateReel(c echo.Context) error {
	var r Reel
	if err := c.Bind(&r); err != nil {
		log.Error(err)
		return echo.NewHTTPError(400)
	}
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(400)
	}
	r.UserID = uint(userID)
	DB.Create(&r)
	return c.JSON(201, r)
}

func handleUpdateReel(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("userID"), 10, 32)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(400)
	}
	reelID, err := strconv.ParseUint(c.Param("reelID"), 10, 32)
	if err != nil {
		log.Error(err)
		return echo.NewHTTPError(400)
	}
	var r Reel
	DB.Where("user_id = ? AND id = ?", userID, reelID).First(&r)
	if r.ID == 0 {
		return echo.NewHTTPError(404)
	}
	var updatedReel Reel
	if err = c.Bind(&updatedReel); err != nil {
		log.Error(err)
		return echo.NewHTTPError(400)
	}
	updatedReel.ID = uint(reelID)
	DB.Save(&updatedReel)
	return c.JSON(200, updatedReel)
}

func handleDeleteReel(c echo.Context) error {
	userID := c.Param("userID")
	reelID := c.Param("reelID")
	var r Reel
	DB.Where("user_id = ? AND id = ?", userID, reelID).First(&r)
	if r.ID == 0 {
		return echo.NewHTTPError(404)
	}
	DB.Delete(&r)
	return c.JSON(204, nil)
}
