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

func TestHandleListUsersEmpty(t *testing.T) {
	DB.AutoMigrate(&User{})
	defer DB.DropTable(&User{})

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	if assert.NoError(t, handleListUsers(c)) {
		assert.Equal(t, 200, res.Code)
		var respUsers []User
		json.Unmarshal(res.Body.Bytes(), &respUsers)
		assert.Empty(t, respUsers)
	}
}

func TestHandleListUsers(t *testing.T) {
	DB.AutoMigrate(&User{})
	defer DB.DropTable(&User{})
	for n := 0; n < 2; n++ {
		u := createDummyUser()
		DB.Create(&u)
	}

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	if assert.NoError(t, handleListUsers(c)) {
		assert.Equal(t, 200, res.Code)
		var respUsers []User
		json.Unmarshal(res.Body.Bytes(), &respUsers)
		assert.Equal(t, 2, len(respUsers))
	}
}

func TestHandleCreateUser(t *testing.T) {
	DB.AutoMigrate(&User{})
	defer DB.DropTable(&User{})

	e := echo.New()
	u := createDummyUser()
	uj, _ := json.Marshal(&u)
	req := httptest.NewRequest(echo.POST, "/users/",
		strings.NewReader(string(uj)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	if assert.NoError(t, handleCreateUser(c)) {
		assert.Equal(t, 201, res.Code)
		var respUser User
		json.Unmarshal(res.Body.Bytes(), &respUser)
		assert.Equal(t, u.FirstName, respUser.FirstName)

		var c int
		DB.Model(&User{}).Count(&c)
		assert.Equal(t, 1, c)
	}
}

func TestHandleCreateUserMalformed(t *testing.T) {
	DB.AutoMigrate(&User{})
	defer DB.DropTable(&User{})

	e := echo.New()
	uj, _ := json.Marshal([]string{"malformed", "payload"})
	req := httptest.NewRequest(echo.POST, "/users/",
		strings.NewReader(string(uj)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	if err := handleCreateUser(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 400, httpError.Code)
		}
	}
}

func TestHandleGetUser(t *testing.T) {
	DB.AutoMigrate(&User{})
	defer DB.DropTable(&User{})

	u := createDummyUser()
	DB.Create(&u)
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID")
	c.SetParamValues(fmt.Sprint(u.ID))
	c.Set("user", u)

	if assert.NoError(t, handleGetUser(c)) {
		assert.Equal(t, 200, res.Code)
		var respUser User
		json.Unmarshal(res.Body.Bytes(), &respUser)
		assert.Equal(t, u.ID, respUser.ID)
	}
}

func TestUpdateUser(t *testing.T) {
	DB.AutoMigrate(&User{})
	defer DB.DropTable(&User{})

	u := createDummyUser()
	DB.Create(&u)
	u.FirstName = "Renamed"
	jsonUser, _ := json.Marshal(&u)
	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/users/",
		strings.NewReader(string(jsonUser)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID")
	c.SetParamValues(fmt.Sprint(u.ID))

	if assert.NoError(t, handleUpdateUser(c)) {
		assert.Equal(t, 200, res.Code)
		var respUser User
		json.Unmarshal(res.Body.Bytes(), &respUser)
		assert.Equal(t, u.FirstName, respUser.FirstName)

		var count int
		DB.Model(&User{}).Count(&count)
	}
}

func TestUpdateUserMalformed(t *testing.T) {
	DB.AutoMigrate(&User{})
	defer DB.DropTable(&User{})

	u := createDummyUser()
	DB.Create(&u)
	payload := map[string]string{"team": "Barcelona"}
	jsonPayload, _ := json.Marshal(&payload)
	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/users/",
		strings.NewReader(string(jsonPayload)))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID")
	c.SetParamValues(fmt.Sprint(u.ID))

	if err := handleUpdateUser(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if ok {
			assert.Equal(t, 400, httpError.Code)
		}
	}
}

func TestDeleteUser(t *testing.T) {
	DB.AutoMigrate(&User{})
	defer DB.DropTable(&User{})

	u := createDummyUser()
	DB.Create(&u)
	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/users/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("userID")
	c.SetParamValues(fmt.Sprint(u.ID))
	c.Set("user", u)

	if assert.NoError(t, handleDeleteUser(c)) {
		assert.Equal(t, 204, res.Code)

		var count int
		DB.Model(&User{}).Count(&count)
		assert.Equal(t, 0, count)
	}
}
