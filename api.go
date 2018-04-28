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

	p := api.Group("/pictures")
	p.GET("/", handleListPictures)
	p.GET("/:id", handleGetPicture)
	p.POST("/", handleCreatePicture)
	p.PUT("/:id", handleUpdatePicture)
	p.DELETE("/:id", handleDeletePicture)
}
