package main

import (
	"github.com/labstack/echo"
)

// GetUserFromURL is a middleware that resolves the user specified by the
// userID portion of a URL, and embeds it in the request's context
func GetUserFromURL(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := c.Param("userID")
		if userID == "" {
			return nil
		}
		var u User
		DB.Where("id = ?", userID).First(&u)
		if u == (User{}) {
			return echo.NewHTTPError(404)
		}
		c.Set("user", u)
		return next(c)
	}
}
