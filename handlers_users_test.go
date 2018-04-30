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

func TestHandleGetUser(t *testing.T) {
	DB.AutoMigrate(&User{})
	defer DB.DropTable(&User{})

	u := createDummyUser()
	DB.Create(&u)
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(u.ID))

	if assert.NoError(t, handleGetUser(c)) {
		assert.Equal(t, 200, res.Code)
		var respUser User
		json.Unmarshal(res.Body.Bytes(), &respUser)
		assert.Equal(t, u.ID, respUser.ID)
	}
}

func TestGetNonexistentUser(t *testing.T) {
	DB.AutoMigrate(&User{})
	defer DB.DropTable(&User{})

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("id")
	c.SetParamValues("200")

	if assert.NoError(t, handleGetUser(c)) {
		assert.Equal(t, 404, res.Code)
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
	c.SetParamNames("id")
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

func TestUpdateNonexistentUser(t *testing.T) {
	DB.AutoMigrate(&User{})
	defer DB.DropTable(&User{})

	u := createDummyUser()
	jsonUser, _ := json.Marshal(&u)
	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/users/",
		strings.NewReader(string(jsonUser)))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(u.ID))

	if assert.NoError(t, handleUpdateUser(c)) {
		assert.Equal(t, 404, res.Code)
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
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(u.ID))

	if assert.NoError(t, handleUpdateUser(c)) {
		assert.Equal(t, 400, res.Code)
	}
}
