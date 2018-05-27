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

func TestHandleListPicturesEmpty(t *testing.T) {
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

	if assert.NoError(t, handleListPictures(c)) {
		assert.Equal(t, 200, res.Code)
		var pictures []Picture
		json.Unmarshal(res.Body.Bytes(), &pictures)
		assert.Empty(t, pictures)
	}
}

func TestHandleListPictures(t *testing.T) {
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

	if assert.NoError(t, handleListPictures(c)) {
		assert.Equal(t, 200, res.Code)
		var pictures []Picture
		json.Unmarshal(res.Body.Bytes(), &pictures)
		assert.Equal(t, 2, len(pictures))
	}
}

func TestHandleGetPicture(t *testing.T) {
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

	if assert.NoError(t, handleGetPicture(c)) {
		assert.Equal(t, 200, res.Code)
		var respPic Picture
		json.Unmarshal(res.Body.Bytes(), &respPic)
		assert.Equal(t, p.ID, respPic.ID)
	}
}

func TestHandleGetNonexistentPicture(t *testing.T) {
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

	if err := handleGetPicture(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 404, httpError.Code)
		}
	}
}

func TestHandleCreatePicture(t *testing.T) {
	DB.AutoMigrate(&User{}, &Picture{})
	defer DB.DropTable(&User{}, &Picture{})

	u := createDummyUser()
	DB.Create(&u)
	p := createDummyPicture()
	jsonPic, _ := json.Marshal(p)
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/users/:userID/pictures/:pictureID",
		strings.NewReader(string(jsonPic)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID")
	c.SetParamValues(fmt.Sprint(u.ID))

	if assert.NoError(t, handleCreatePicture(c)) {
		assert.Equal(t, 201, res.Code)
		var respPic Picture
		json.Unmarshal(res.Body.Bytes(), &respPic)
		assert.Equal(t, u.ID, respPic.ID)
	}
}

func TestHandleCreatePictureMalformed(t *testing.T) {
	DB.AutoMigrate(&User{}, &Picture{})
	defer DB.DropTable(&User{}, &Picture{})

	u := createDummyUser()
	DB.Create(&u)
	payload := map[string]int{"Caption": 1}
	jsonPayload, _ := json.Marshal(payload)
	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/users/:userID/pictures",
		strings.NewReader(string(jsonPayload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID")
	c.SetParamValues(fmt.Sprint(u.ID))

	if err := handleCreatePicture(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 400, httpError.Code)
		}
	}
}

func TestHandleUpdatePicture(t *testing.T) {
	DB.AutoMigrate(&User{}, &Picture{})
	defer DB.DropTable(&User{}, &Picture{})

	u := createDummyUser()
	DB.Create(&u)
	p := createDummyPicture()
	p.UserID = u.ID
	DB.Create(&p)
	p.URL = "http://updated-url.com/dummy-pic.jpeg"
	jsonPayload, _ := json.Marshal(p)

	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/users/:userID/pictures/:pictureID",
		strings.NewReader(string(jsonPayload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID", "pictureID")
	c.SetParamValues(fmt.Sprint(u.ID), fmt.Sprint(p.ID))

	if assert.NoError(t, handleUpdatePicture(c)) {
		assert.Equal(t, 200, res.Code)
		var respPicture Picture
		json.Unmarshal(res.Body.Bytes(), &respPicture)
		assert.Equal(t, p.URL, respPicture.URL)
	}
}

func TestUpdateNonexistantPicture(t *testing.T) {
	DB.AutoMigrate(&User{}, &Picture{})
	defer DB.DropTable(&User{}, &Picture{})

	u := createDummyUser()
	DB.Create(&u)
	p := createDummyPicture()
	p.URL = "http://updated-url.com/dummy-pic.jpeg"
	jsonPayload, _ := json.Marshal(p)

	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/users/:userID/pictures/:pictureID",
		strings.NewReader(string(jsonPayload)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID", "pictureID")
	c.SetParamValues(fmt.Sprint(u.ID), fmt.Sprint(500))

	if err := handleUpdatePicture(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 404, httpError.Code)
		}
	}
}

func TestHandleUpdatePictureMalformed(t *testing.T) {
	DB.AutoMigrate(&User{}, &Picture{})
	defer DB.DropTable(&User{}, &Picture{})

	u := createDummyUser()
	DB.Create(&u)
	p := createDummyPicture
	DB.Create(&p)
	payload := []string{"some", "invalid", "payload"}
	jsonPayload, _ := json.Marshal(payload)

	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/users/:userID/pictures/:pictureID",
		strings.NewReader(string(jsonPayload)))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	if err := handleUpdatePicture(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 400, httpError.Code)
		}
	}
}

func TestHandleDeletePicture(t *testing.T) {
	DB.AutoMigrate(&User{}, &Picture{})
	defer DB.DropTable(&User{}, &Picture{})

	u := createDummyUser()
	DB.Create(&u)
	p := createDummyPicture()
	p.UserID = u.ID
	DB.Create(&p)

	e := echo.New()
	req := httptest.NewRequest(echo.DELETE,
		"/users/:userID/pictures/:pictureID", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID", "pictureID")
	c.SetParamValues(fmt.Sprint(u.ID), fmt.Sprint(p.ID))

	if assert.NoError(t, handleDeletePicture(c)) {
		assert.Equal(t, 204, res.Code)
	}
}

func TestHandleDeleteNonexistentPicture(t *testing.T) {
	DB.AutoMigrate(&User{}, &Picture{})
	defer DB.DropTable(&User{}, &Picture{})

	e := echo.New()
	req := httptest.NewRequest(echo.DELETE,
		"/user/:userID/pictures/:pictureID", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID", "pictureID")
	c.SetParamValues("nonexistent-user", "nonexistent-pic")

	if err := handleDeletePicture(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 404, httpError.Code)
		}
	}
}
