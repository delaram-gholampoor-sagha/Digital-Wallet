package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestSignUpHandler(t *testing.T) {
	// Create a new Echo instance for testing
	e := echo.New()

	// Create a mock UserService object
	mockUserService := new(protocol.MockUserRepository)

	// Create a new SignUp request object
	reqBody := &request.SignUp{
		Username: "testuser",
		Password: "testpassword",
	}

	// Create a response object that you expect to get back after successful sign-up
	expectedRespBody := &response.SignUp{
		AccessToken:  "test_access_token",
		RefreshToken: "test_refresh_token",
	}

	// Set up the mock UserService's expectations
	mockUserService.On("SignUp", mock.Anything, *reqBody).Return(*expectedRespBody, nil)

	// Create a SignUpHandler with our mock UserService
	h := SignUpHandler(mockUserService)

	// Marshal the request body to JSON
	reqBodyJSON, _ := json.Marshal(reqBody)

	// Create a request and response recorder
	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(reqBodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	// Create a new Echo context for the request
	c := e.NewContext(req, rec)

	// Call the SignUpHandler
	if err := h(c); err != nil {
		t.Fatalf("Handler returned an error: %v", err)
		return
	}

	// Assert status code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Parse and print the response body
	var respBody map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &respBody); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
		return
	}
	fmt.Printf("Response: %+v\n", respBody) // Debug print

	// Assert the response body
	if respBody["data"] != nil {
		dataMap := respBody["data"].(map[string]interface{})
		assert.Equal(t, "success", respBody["message"])
		assert.Equal(t, expectedRespBody.AccessToken, dataMap["access_token"])
		assert.Equal(t, expectedRespBody.RefreshToken, dataMap["refresh_token"])
	} else {
		t.Fatal("Data field is nil")
	}

	// Assert Expectations on the mock object (Optional)
	mockUserService.AssertExpectations(t)
}

const (
	mockUserID       = 42
	mockToken        = "new_token"
	mockRefreshToken = "new_refresh_token"
)

func TestRefreshTokenHandler(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Set up the dependencies
	mockUserService := new(protocol.MockUserRepository)
	handler := RefreshTokenHandler(mockUserService)

	// Create a request
	req := httptest.NewRequest(http.MethodGet, "/refresh-token", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Mock JWT Token
	mockClaims := &entity.JWTClaims{
		UserID: mockUserID,
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = mockClaims
	c.Set("user", token)

	// Set up the expected response from the mock user service
	expectedTokens := response.RefreshToken{
		AccessToken: mockToken,
	}
	mockUserService.On("RefreshToken", mock.Anything, mockUserID).Return(expectedTokens, nil)

	// Call the handler
	if err := handler(c); err != nil {
		t.Errorf("handler returned an error: %v", err)
	}

	// Check the response code
	assert.Equal(t, http.StatusOK, rec.Code)

	// Unmarshal the response
	var resp protocol.Success
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}

	// Type assertion for 'Data'
	actualTokens, ok := resp.Data.(map[string]interface{})
	if !ok {
		t.Errorf("Could not assert response data to expected type")
	}

	// Now, manually construct your expected type from the map
	expectedAccessToken := expectedTokens.AccessToken
	actualAccessToken := actualTokens["access_token"].(string)

	// Now compare the individual fields
	assert.Equal(t, expectedAccessToken, actualAccessToken)

	// Assert that the methods on the mock user service were called as expected
	mockUserService.AssertExpectations(t)
}
func TestSignInHandler(t *testing.T) {
	e := echo.New()
	mockUserService := new(protocol.MockUserRepository)
	reqBody := &request.SignIn{
		Username: "testuser",
		Password: "testpassword",
	}
	expectedRespBody := &response.SignIn{
		AccessToken:  "test_access_token",
		RefreshToken: "test_refresh_token",
	}

	mockUserService.On("SignIn", mock.Anything, *reqBody).Return(*expectedRespBody, nil)
	h := SignInHandler(mockUserService)
	reqBodyJSON, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/signin", bytes.NewReader(reqBodyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if err := h(c); err != nil {
		t.Fatalf("Handler returned an error: %v", err)
		return
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	var respBody map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &respBody); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
		return
	}

	// Your assertions here (similar to the SignUpHandler test)
	// ...
}

func TestGetProfileHandler(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Set up the dependencies
	mockUserService := new(protocol.MockUserRepository)
	handler := GetProfileHandler(mockUserService)

	// Create a request
	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up a mock JWT Token and set it in the context
	token := &jwt.Token{
		Claims: &entity.JWTClaims{
			UserID: mockUserID,
		},
	}
	c.Set("user", token)

	// Set up the expected response from the mock user service
	expectedProfile := response.GetProfile{
		// populate this as needed
	}
	mockUserService.On("GetProfile", mock.Anything, mockUserID).Return(expectedProfile, nil)

	// Call the handler
	err := handler(c)
	assert.NoError(t, err)

	// Check the response
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Assert other checks for the response...

	// Assert that the methods on the mock user service were called as expected
	mockUserService.AssertExpectations(t)
}
func TestEditProfileHandler(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Set up the dependencies
	mockUserService := new(protocol.MockUserRepository)
	handler := EditProfileHandler(mockUserService)

	// Create a request with some JSON payload
	req := httptest.NewRequest(http.MethodPost, "/edit-profile", strings.NewReader(`{
		"UserID": 42,
		"Password": "mfkldfl",
		"FirstName": "sajad",
		"LastName": "vaezi",
		"Email": "majesticdelaram@yahoo.com",
		"Cellphone": "3894983748"
	  }`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Set up a mock JWT Token and set it in the context
	token := &jwt.Token{
		Claims: &entity.JWTClaims{
			UserID: mockUserID,
		},
	}
	c.Set("user", token)

	mockEditProfileRequest := request.EditProfile{
		UserID:    42,
		Password:  "mfkldfl",
		FirstName: "", // note the change here
		LastName:  "", // and here
		Email:     "majesticdelaram@yahoo.com",
		Cellphone: "3894983748",
	}
	mockUserService.On("EditProfile", mock.Anything, mockEditProfileRequest).Return(nil)

	// Call the handler
	err := handler(c)
	assert.NoError(t, err)

	// Check the response
	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Assert other checks for the response...

	// Assert that the methods on the mock user service were called as expected
	mockUserService.AssertExpectations(t)
}
