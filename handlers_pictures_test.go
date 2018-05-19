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

/* func TestHandleListPicturesNonexistentUser(t *testing.T) {
	DB.AutoMigrate(&User{}, &Picture{})
	defer DB.DropTable(&User{}, &Picture{})

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/user/:userID/pictures/",
		strings.NewReader(")"))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID")
	c.SetParamValues("200")

	if err := handleListUserPictures(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 404, httpError.Code)
		}
	}
} */

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

// func TestHandleGetNonexistentUserPic(t *testing.T) {
// 	e := echo.New()
// 	req := httptest.NewRequest(echo.GET, "/users/:userID/pictures/:pictureID",
// 		strings.NewReader(""))
// 	res := httptest.NewRecorder()
// 	c := e.NewContext(req, res)
// 	c.SetParamNames("userID", "pictureID")
// 	c.SetParamValues("200", "2")

// 	if assert.NoError(t, handleGetUserPicture(c)) {
// 		assert.Equal(t, 404, res.Code)
// 	}
// }

func TestHandleGetUserNonexistentPicture(t *testing.T) {
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

	if err := handleGetUserPicture(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 404, httpError.Code)
		}
	}
}

func TestHandleCreateUserPicture(t *testing.T) {
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

	if assert.NoError(t, handleCreateUserPicture(c)) {
		assert.Equal(t, 201, res.Code)
		var respPic Picture
		json.Unmarshal(res.Body.Bytes(), &respPic)
		assert.Equal(t, u.ID, respPic.ID)
	}
}

// func TestHandleCreateNonexistentUserPicture(t *testing.T) {
// 	DB.AutoMigrate(&User{}, &Picture{})
// 	defer DB.DropTable(&User{}, &Picture{})

// 	p := createDummyPicture()
// 	jsonPic, _ := json.Marshal(p)
// 	e := echo.New()
// 	req := httptest.NewRequest(echo.POST, "/users/:userID/pictures/:pictureID",
// 		strings.NewReader(string(jsonPic)))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	res := httptest.NewRecorder()
// 	c := e.NewContext(req, res)
// 	c.SetParamNames("userID")
// 	c.SetParamValues("2")

// 	if assert.NoError(t, handleCreateUserPicture(c)) {
// 		assert.Equal(t, 404, res.Code)
// 	}
// }

func TestHandleCreateUserPictureMalformed(t *testing.T) {
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

	if err := handleCreateUserPicture(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 400, httpError.Code)
		}
	}
}

func TestHandleUpdateUserPicture(t *testing.T) {
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

	if assert.NoError(t, handleUpdateUserPicture(c)) {
		assert.Equal(t, 200, res.Code)
		var respPicture Picture
		json.Unmarshal(res.Body.Bytes(), &respPicture)
		assert.Equal(t, p.URL, respPicture.URL)
	}
}

func TestUpdateNonexistantUserPicture(t *testing.T) {
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

	if err := handleUpdateUserPicture(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 404, httpError.Code)
		}
	}
}

func TestHandleUpdateUserPictureMalformed(t *testing.T) {
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

	if err := handleUpdateUserPicture(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 400, httpError.Code)
		}
	}
}

func TestHandleDeleteUserPicture(t *testing.T) {
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

	if assert.NoError(t, handleDeleteUserPicture(c)) {
		assert.Equal(t, 204, res.Code)
	}
}

func TestHandleDeleteNonexistentUserPicture(t *testing.T) {
	DB.AutoMigrate(&User{}, &Picture{})
	defer DB.DropTable(&User{}, &Picture{})

	e := echo.New()
	req := httptest.NewRequest(echo.DELETE,
		"/user/:userID/pictures/:pictureID", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID", "pictureID")
	c.SetParamValues("nonexistent-user", "nonexistent-pic")

	if err := handleDeleteUserPicture(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 404, httpError.Code)
		}
	}
}
