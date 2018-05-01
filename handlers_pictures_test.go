package main

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHandleListUserPicturesEmpty(t *testing.T) {
	DB.AutoMigrate(&User{}, &Picture{})
	defer DB.DropTable(&User{}, &Picture{})

	u := createDummyUser()
	DB.Create(&u)
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/:userID/pictures/",
		strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID")
	c.SetParamValues(fmt.Sprint(u.ID))

	if assert.NoError(t, handleListUserPictures(c)) {
		assert.Equal(t, 200, res.Code)
		var pictures []Picture
		json.Unmarshal(res.Body.Bytes(), &pictures)
		assert.Empty(t, pictures)
	}
}

func TestHandleListUserPictures(t *testing.T) {
	DB.AutoMigrate(&User{}, &Picture{})
	defer DB.DropTable(&User{}, &Picture{})

	u := createDummyUser()
	DB.Create(&u)
	for n := 0; n < 2; n++ {
		p := createDummyPicture()
		p.UserID = u.ID
		DB.Create(&p)
	}
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/:userID/pictures/",
		strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID")
	c.SetParamValues(fmt.Sprint(u.ID))

	if assert.NoError(t, handleListUserPictures(c)) {
		assert.Equal(t, 200, res.Code)
		var pictures []Picture
		json.Unmarshal(res.Body.Bytes(), &pictures)
		assert.Equal(t, 2, len(pictures))
	}
}

func TestHandleListPicturesNonexistentUser(t *testing.T) {
	DB.AutoMigrate(&User{}, &Picture{})
	defer DB.DropTable(&User{}, &Picture{})

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/user/:userID/pictures/",
		strings.NewReader(")"))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID")
	c.SetParamValues("200")

	if assert.NoError(t, handleListUserPictures(c)) {
		assert.Equal(t, 404, res.Code)
	}
}

func TestHandleGetUserPicture(t *testing.T) {
	DB.AutoMigrate(&User{}, &Picture{})
	defer DB.DropTable(&User{}, &Picture{})

	u := createDummyUser()
	DB.Create(&u)
	p := createDummyPicture()
	p.UserID = u.ID
	DB.Create(&p)
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/user/:userID/pictures/:pictureID",
		strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID", "pictureID")
	c.SetParamValues(fmt.Sprint(u.ID), fmt.Sprint(p.ID))

	if assert.NoError(t, handleGetUserPicture(c)) {
		assert.Equal(t, 200, res.Code)
		var respPic Picture
		json.Unmarshal(res.Body.Bytes(), &respPic)
		assert.Equal(t, p.ID, respPic.ID)
	}
}

func TestHandleGetNonexistentUserPic(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/:userID/pictures/:pictureID",
		strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID", "pictureID")
	c.SetParamValues("200", "2")

	if assert.NoError(t, handleGetUserPicture(c)) {
		assert.Equal(t, 404, res.Code)
	}
}

func TestHandleUserNonexistentPicture(t *testing.T) {
	DB.AutoMigrate(&User{}, &Picture{})
	defer DB.DropTable(&User{}, &Picture{})

	u := createDummyUser()
	DB.Create(&u)
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/:userID/pictures/:pictureID",
		strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID", "pictureID")
	c.SetParamValues(fmt.Sprint(u.ID), "2")

	if assert.NoError(t, handleGetUserPicture(c)) {
		assert.Equal(t, 404, res.Code)
	}
}
