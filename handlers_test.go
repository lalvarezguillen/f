package main

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func testGetAPIInfo(t *testing.T) {
	req := httptest.NewRequest(echo.GET, "/", nil)
	resp := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, resp)

	if assert.NoError(t, handleGetAPIInfo(c)) {
		assert.Equal(t, 200, resp.Code)
	}
}
