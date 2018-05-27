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

	r := api.Group("/reels")
	r.GET("/", handleListReels)
	r.GET("/:id", handleGetReel)
	r.POST("/", handleCreateReel)
	r.PUT("/:id", handleUpdateReel)
	r.DELETE("/:id", handleDeleteUser)

	u.GET("/:userID/pictures/", handleListPictures)
	u.GET("/:userID/pictures/:pictureID", handleGetPicture)
	u.POST("/:userID/pictures/", handleCreatePicture)
	u.PUT("/:userID/pictures/:pictureID", handleUpdatePicture)
	u.DELETE("/:userID/pictures/:pictureID", handleDeletePicture)
}
