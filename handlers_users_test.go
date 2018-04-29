package main

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestHandleListUsersEmpty(t *testing.T) {
	// Setup
	DB.AutoMigrate(&User{})
	defer DB.DropTable(&User{})

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users/", strings.NewReader(""))
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	// Test
	if assert.NoError(t, handleListUsers(c)) {
		assert.Equal(t, 200, res.Code)
		var respUsers []User
		json.Unmarshal(res.Body.Bytes(), respUsers)
		assert.Empty(t, respUsers)
	}
}
