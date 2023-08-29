package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestJWT(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/api/protected", nil)

	// Create a new HTTP response recorder
	rec := httptest.NewRecorder()

	// Define the secret for JWT
	secret := "my-secret-key"

	// Define the JWT middleware
	jwtMiddleware := JWT(secret)

	// Create a new Echo group and attach the JWT middleware
	g := e.Group("/api")
	g.Use(jwtMiddleware)

	// Define a handler function for the protected route
	handler := func(c echo.Context) error {
		return c.String(http.StatusOK, "Protected route")
	}

	// Register the handler function for the protected route
	g.GET("/protected", handler)

	// Perform the HTTP request
	e.ServeHTTP(rec, req)

	// Assert that the response code is 401 Unauthorized
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
