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

	mockUserService := new(protocol.MockUserRepository)

	reqBody := &request.SignUp{
		Username: "testuser",
		Password: "testpassword",
	}

	expectedRespBody := &response.SignUp{
		AccessToken:  "test_access_token",
		RefreshToken: "test_refresh_token",
	}

	mockUserService.On("SignUp", mock.Anything, *reqBody).Return(*expectedRespBody, nil)

	h := SignUpHandler(mockUserService)

	reqBodyJSON, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/signup", bytes.NewReader(reqBodyJSON))
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
	fmt.Printf("Response: %+v\n", respBody)

	if respBody["data"] != nil {
		dataMap := respBody["data"].(map[string]interface{})
		assert.Equal(t, "success", respBody["message"])
		assert.Equal(t, expectedRespBody.AccessToken, dataMap["access_token"])
		assert.Equal(t, expectedRespBody.RefreshToken, dataMap["refresh_token"])
	} else {
		t.Fatal("Data field is nil")
	}

	mockUserService.AssertExpectations(t)
}

const (
	mockUserID       = 42
	mockToken        = "new_token"
	mockRefreshToken = "new_refresh_token"
)

func TestRefreshTokenHandler(t *testing.T) {

	e := echo.New()

	mockUserService := new(protocol.MockUserRepository)
	handler := RefreshTokenHandler(mockUserService)

	req := httptest.NewRequest(http.MethodGet, "/refresh-token", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockClaims := &entity.JWTClaims{
		UserID: mockUserID,
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = mockClaims
	c.Set("user", token)

	expectedTokens := response.RefreshToken{
		AccessToken: mockToken,
	}
	mockUserService.On("RefreshToken", mock.Anything, mockUserID).Return(expectedTokens, nil)

	if err := handler(c); err != nil {
		t.Errorf("handler returned an error: %v", err)
	}

	assert.Equal(t, http.StatusOK, rec.Code)

	var resp protocol.Success
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Errorf("failed to unmarshal response: %v", err)
	}

	actualTokens, ok := resp.Data.(map[string]interface{})
	if !ok {
		t.Errorf("Could not assert response data to expected type")
	}

	expectedAccessToken := expectedTokens.AccessToken
	actualAccessToken := actualTokens["access_token"].(string)

	assert.Equal(t, expectedAccessToken, actualAccessToken)

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

}

func TestGetProfileHandler(t *testing.T) {

	e := echo.New()

	mockUserService := new(protocol.MockUserRepository)
	handler := GetProfileHandler(mockUserService)

	req := httptest.NewRequest(http.MethodGet, "/profile", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	token := &jwt.Token{
		Claims: &entity.JWTClaims{
			UserID: mockUserID,
		},
	}
	c.Set("user", token)

	expectedProfile := response.GetProfile{}
	mockUserService.On("GetProfile", mock.Anything, mockUserID).Return(expectedProfile, nil)

	err := handler(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	mockUserService.AssertExpectations(t)
}
func TestEditProfileHandler(t *testing.T) {

	e := echo.New()

	mockUserService := new(protocol.MockUserRepository)
	handler := EditProfileHandler(mockUserService)

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

	token := &jwt.Token{
		Claims: &entity.JWTClaims{
			UserID: mockUserID,
		},
	}
	c.Set("user", token)

	mockEditProfileRequest := request.EditProfile{
		UserID:    42,
		Password:  "mfkldfl",
		FirstName: "",
		LastName:  "",
		Email:     "majesticdelaram@yahoo.com",
		Cellphone: "3894983748",
	}
	mockUserService.On("EditProfile", mock.Anything, mockEditProfileRequest).Return(nil)

	err := handler(c)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	mockUserService.AssertExpectations(t)
}
