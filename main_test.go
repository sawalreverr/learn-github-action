package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHTTP(t *testing.T) {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!") // 200, Hello, World!
	})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	e.ServeHTTP(recorder, request)

	assert.Equal(t, 200, recorder.Code)                      // Status harus 200
	assert.Equal(t, "Hello, World!", recorder.Body.String()) // Body harus Hello, World!
}
