package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/translation"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLocale(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Define a handler function for testing the middleware
	handler := func(c echo.Context) error {
		locale := c.Get(translation.Locale).(translation.Language)
		return c.String(http.StatusOK, fmt.Sprintf("%d", locale))
	}

	// Define a mapping between string values and translation.Language values
	localeMapping := map[string]translation.Language{
		"en": translation.English,
		"fa": translation.Farsi,
	}

	// Create a new HTTP request with Accept-Language header
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Accept-Language", "en-US")

	// Create a new HTTP response recorder
	rec := httptest.NewRecorder()

	// Create a new Echo instance and attach the Locale middleware
	e.Use(Locale())
	e.GET("/", handler)

	// Perform the HTTP request
	e.ServeHTTP(rec, req)

	// Assert that the response body matches the expected Language value
	expectedLocale := localeMapping["en"]
	actualLocale, _ := strconv.Atoi(rec.Body.String())
	assert.Equal(t, expectedLocale, translation.Language(actualLocale))
	assert.Equal(t, http.StatusOK, rec.Code)

	// Create a new HTTP request with Accept-Language header
	req = httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Accept-Language", "fa-IR")

	// Reset the HTTP response recorder
	rec = httptest.NewRecorder()

	// Perform the HTTP request
	e.ServeHTTP(rec, req)

	// Assert that the response body matches the expected Language value
	expectedLocale = localeMapping["fa"]
	actualLocale, _ = strconv.Atoi(rec.Body.String())
	assert.Equal(t, expectedLocale, translation.Language(actualLocale))
	assert.Equal(t, http.StatusOK, rec.Code)

	// Create a new HTTP request without Accept-Language header
	req = httptest.NewRequest(http.MethodGet, "/", nil)

	// Reset the HTTP response recorder
	rec = httptest.NewRecorder()

	// Perform the HTTP request
	e.ServeHTTP(rec, req)

	// Assert that the response body matches the expected Language value
	expectedLocale = localeMapping["fa"]
	actualLocale, _ = strconv.Atoi(rec.Body.String())
	assert.Equal(t, expectedLocale, translation.Language(actualLocale))
	assert.Equal(t, http.StatusOK, rec.Code)
}
