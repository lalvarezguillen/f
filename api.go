package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	api := echo.New()
	api.Use(middleware.Logger())
	api.GET("/", getAPIInfo)

	u := api.Group("/users", GetUserFromURL)
	u.GET("/", handleListUsers)
	u.GET("/:userID", handleGetUser)
	u.POST("/", handleCreateUser)
	u.PUT("/:userID", handleUpdateUser)
	u.DELETE("/:userID", handleDeleteUser)

	u.GET("/:userID/reels/", handleListReels)
	u.GET("/:userID/reels/:reelID", handleGetReel)
	u.POST("/:userID/reels/", handleCreateReel)
	u.PUT("/:userID/reels/:reelID", handleUpdateReel)
	u.DELETE("/:userID/reels/:reelID", handleDeleteUser)

	u.GET("/:userID/pictures/", handleListPictures)
	u.GET("/:userID/pictures/:pictureID", handleGetPicture)
	u.POST("/:userID/pictures/", handleCreatePicture)
	u.PUT("/:userID/pictures/:pictureID", handleUpdatePicture)
	u.DELETE("/:userID/pictures/:pictureID", handleDeletePicture)
}
