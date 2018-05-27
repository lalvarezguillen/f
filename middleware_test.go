package main

import (
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetUserFromURL(t *testing.T) {
	DB.AutoMigrate(&User{})
	defer DB.DropTable(&User{})
	u := createDummyUser()
	DB.Create(&u)

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/:userID", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	c.SetParamNames("userID")
	c.SetParamValues(fmt.Sprint(u.ID))
	dummyHandler := func(c echo.Context) error {
		return nil
	}

	h := GetUserFromURL(dummyHandler)
	h(c)

	userFromC, ok := c.Get("user").(User)
	assert.True(t, ok)
	assert.Equal(t, u.ID, userFromC.ID)
}

func TestGetUserFromURLMalformed(t *testing.T) {
	DB.AutoMigrate(&User{})
	defer DB.DropTable(&User{})

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/:userID", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	c.SetParamNames("userID")
	c.SetParamValues("")
	dummyHandler := func(c echo.Context) error {
		return nil
	}

	h := GetUserFromURL(dummyHandler)
	if err := h(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 400, httpError.Code)
		}
	}
}

func TestGetUserFromURLNonexistent(t *testing.T) {
	DB.AutoMigrate(&User{})
	defer DB.DropTable(&User{})

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/:userID", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	c.SetParamNames("userID")
	c.SetParamValues("nonexistent-user")
	dummyHandler := func(c echo.Context) error {
		return nil
	}

	h := GetUserFromURL(dummyHandler)
	if err := h(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 404, httpError.Code)
		}
	}
}
