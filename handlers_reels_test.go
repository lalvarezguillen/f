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

func TestHandleListReelsEmpty(t *testing.T) {
	DB.AutoMigrate(&User{}, &Reel{})
	defer DB.DropTable(&User{}, &Reel{})
	u := createDummyUser()
	DB.Create(&u)

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/:userID/reels/", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	c.SetParamNames("userID")
	c.SetParamValues(fmt.Sprint(u.ID))
	c.Set("user", u)

	if assert.NoError(t, handleListReels(c)) {
		assert.Equal(t, 200, resp.Code)
		var r []Reel
		json.Unmarshal(resp.Body.Bytes(), &r)
		assert.Empty(t, r)
	}
}

func TestHandleListReels(t *testing.T) {
	DB.AutoMigrate(&User{}, &Reel{})
	defer DB.DropTable(&User{}, &Reel{})
	u := createDummyUser()
	DB.Create(&u)
	reelsCount := 2
	for idx := 0; idx < reelsCount; idx++ {
		r := createDummyReel()
		r.UserID = u.ID
		DB.Create(&r)
	}

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/:userID/reels/", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	c.SetParamNames("userID")
	c.SetParamValues(fmt.Sprint(u.ID))
	c.Set("user", u)

	if assert.NoError(t, handleListReels(c)) {
		assert.Equal(t, 200, resp.Code)
		var reels []Reel
		json.Unmarshal(resp.Body.Bytes(), &reels)
		assert.Equal(t, reelsCount, len(reels))
	}
}

func TestHandleGetReel(t *testing.T) {
	DB.AutoMigrate(&User{}, &Reel{})
	defer DB.DropTable(&User{}, &Reel{})
	u := createDummyUser()
	DB.Create(&u)
	r := createDummyReel()
	r.UserID = u.ID
	DB.Create(&r)

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/:userID/reels/:reelID", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	c.SetParamNames("userID", "reelID")
	c.SetParamValues(fmt.Sprint(u.ID), fmt.Sprint(r.ID))
	c.Set("user", u)

	if assert.NoError(t, handleGetReel(c)) {
		assert.Equal(t, 200, resp.Code)
		var reel Reel
		json.Unmarshal(resp.Body.Bytes(), &reel)
		assert.Equal(t, r.ID, reel.ID)
	}
}

func TestHandleGetNonexistentReel(t *testing.T) {
	DB.AutoMigrate(&User{}, &Reel{})
	defer DB.DropTable(&User{}, &Reel{})
	u := createDummyUser()
	DB.Create(&u)

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/:userID/reels/:reelID", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	c.SetParamNames("userID", "reelID")
	c.SetParamValues(fmt.Sprint(u.ID), "nonexistent-reel")
	c.Set("user", u)

	if err := handleGetReel(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 404, httpError.Code)
		}
	}
}

func TestHandleCreateReel(t *testing.T) {
	DB.AutoMigrate(&User{}, &Reel{})
	defer DB.DropTable(&User{}, &Reel{})
	u := createDummyUser()
	DB.Create(&u)
	r := createDummyReel()
	jsonReel, _ := json.Marshal(&r)

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/users/:userID/reels/",
		strings.NewReader(string(jsonReel)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	c.SetParamNames("userID")
	c.SetParamValues(fmt.Sprint(u.ID))

	if assert.NoError(t, handleCreateReel(c)) {
		assert.Equal(t, 201, resp.Code)
		var reel Reel
		json.Unmarshal(resp.Body.Bytes(), &reel)
		assert.Equal(t, u.ID, reel.UserID)
	}
}

func TestHandleCreateReelMalformed(t *testing.T) {
	DB.AutoMigrate(&User{}, &Reel{})
	defer DB.DropTable(&User{}, &Reel{})
	u := createDummyUser()
	DB.Create(&u)
	jsonPayload, _ := json.Marshal([]int{1, 2, 3})

	e := echo.New()
	req := httptest.NewRequest(echo.POST, "/users/:userID/reels/",
		strings.NewReader(string(jsonPayload)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	c.SetParamNames("userID")
	c.SetParamValues(fmt.Sprint(u.ID))
	c.Set("user", u)

	if err := handleCreateReel(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 400, httpError.Code)
		}
	}
}

func TestHandleUpdateReel(t *testing.T) {
	DB.AutoMigrate(&User{}, &Reel{})
	defer DB.DropTable(&User{}, &Reel{})
	u := createDummyUser()
	DB.Create(&u)
	r := createDummyReel()
	r.UserID = u.ID
	DB.Create(&r)
	r.Title = "Updated Title"
	jsonReel, _ := json.Marshal(&r)

	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/users/:userID/reels/:reelID",
		strings.NewReader(string(jsonReel)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	c.SetParamNames("userID", "reelID")
	c.SetParamValues(fmt.Sprint(u.ID), fmt.Sprint(r.ID))
	c.Set("user", u)

	if assert.NoError(t, handleUpdateReel(c)) {
		var updatedReel Reel
		json.Unmarshal(resp.Body.Bytes(), &updatedReel)
		assert.Equal(t, r.Title, updatedReel.Title)
	}
}

func TestHandleUpdateNonexistentReel(t *testing.T) {
	DB.AutoMigrate(&User{}, &Reel{})
	defer DB.DropTable(&User{}, &Reel{})
	u := createDummyUser()
	DB.Create(&u)
	r := createDummyReel()
	jsonReel, _ := json.Marshal(&r)

	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/users/:userID/reels/:reelID",
		strings.NewReader(string(jsonReel)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	c.SetParamNames("userID", "reelID")
	c.SetParamValues(fmt.Sprint(u.ID), "9000")
	c.Set("user", u)

	if err := handleUpdateReel(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 404, httpError.Code)
		}
	}
}

func TestHandleUpdateReelMalformed(t *testing.T) {
	DB.AutoMigrate(&User{}, &Reel{})
	defer DB.DropTable(&User{}, &Reel{})
	u := createDummyUser()
	DB.Create(&u)
	r := createDummyReel()
	DB.Create(&r)
	jsonPayload, _ := json.Marshal([]int{1, 2, 3})

	e := echo.New()
	req := httptest.NewRequest(echo.PUT, "/users/:userID/reels/:reelID",
		strings.NewReader(string(jsonPayload)))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	c.SetParamNames("userID", "reelID")
	c.SetParamValues(fmt.Sprint())
	c.Set("user", u)

	if err := handleUpdateReel(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 400, httpError.Code)
		}
	}
}

func TestHandleDeleteReel(t *testing.T) {
	DB.AutoMigrate(&User{}, &Reel{})
	defer DB.DropTable(&User{}, &Reel{})
	u := createDummyUser()
	DB.Create(&u)
	r := createDummyReel()
	r.UserID = u.ID
	DB.Create(&r)

	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/users/:userID/reels/:reelID", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	c.SetParamNames("userID", "reelID")
	c.SetParamValues(fmt.Sprint(u.ID), fmt.Sprint(r.ID))
	c.Set("user", u)

	if assert.NoError(t, handleDeleteReel(c)) {
		assert.Equal(t, 204, resp.Code)
	}
}

func TestHandleDeleteNonexistentReel(t *testing.T) {
	DB.AutoMigrate(&User{}, &Reel{})
	defer DB.DropTable(&User{}, &Reel{})
	u := createDummyUser()
	DB.Create(&u)

	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/users/:userID/reels/:reelID", nil)
	resp := httptest.NewRecorder()
	c := e.NewContext(req, resp)
	c.SetParamNames("userID", "reelID")
	c.SetParamValues(fmt.Sprint(u.ID), "nonexistent-reel")
	c.Set("user", u)

	if err := handleDeleteReel(c); assert.Error(t, err) {
		httpError, ok := err.(*echo.HTTPError)
		if assert.True(t, ok) {
			assert.Equal(t, 404, httpError.Code)
		}
	}
}
