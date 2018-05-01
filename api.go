package main

import "github.com/labstack/echo"

func main() {
	api := echo.New()
	api.GET("/", getAPIInfo)

	u := api.Group("/users")
	u.GET("/", handleListUsers)
	u.GET("/:id", handleGetUser)
	u.POST("/", handleCreateUser)
	u.PUT("/:id", handleUpdateUser)
	u.DELETE("/:id", handleDeleteUser)

	r := api.Group("/reels")
	r.GET("/", handleListReels)
	r.GET("/:id", handleGetReel)
	r.POST("/", handleCreateReel)
	r.PUT("/:id", handleUpdateReel)
	r.DELETE("/:id", handleDeleteUser)

	u.GET("/:userID/pictures/", handleListUserPictures)
	u.GET("/:userID/pictures/:pictureID", handleGetUserPicture)
	u.POST("/:userID/pictures/", handleCreateUserPicture)
	u.PUT("/:userID/pictures/:pictureID", handleUpdateUserPicture)
	u.DELETE("/:userID/pictures/:pictureID", handleDeleteUserPicture)
}
