package user

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/config"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror/message"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

// go test github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/service/user -coverprofile=coverage.out
// go tool cover -func=coverage.out

func setup() (*Service, *protocol.MockUserRepository, *protocol.MockHasher) {
	mockRepo := new(protocol.MockUserRepository)
	mockHasher := new(protocol.MockHasher)

	cfg := config.JWT{
		AccessTokenExp:  time.Minute * 15,
		RefreshTokenExp: time.Hour * 24,
		Secret:          "testSecret",
	}

	tokenGen := utils.JWTTokenGenerator{}

	logger, _ := zap.NewProduction()
	sugaredLogger := logger.Sugar()

	service := &Service{
		userRepo: mockRepo,
		hasher:   mockHasher,
		tokenGen: &tokenGen,
		cfg:      cfg,
		logger:   sugaredLogger,
	}
	return service, mockRepo, mockHasher
}

func TestService_SignUp(t *testing.T) {
	service, mockRepo, mockHasher := setup()

	tests := []struct {
		name      string
		req       request.SignUp
		mockCalls func()
		wantErr   bool
		errType   error
	}{
		{
			name: "successful sign up",
			req: request.SignUp{
				Username:  "testUser",
				Password:  "testPass",
				FirstName: "deliiii",
				LastName:  "qobadiiii",
				Email:     "majesticdelaram@yahoo.com",
				Cellphone: "2398483984",
			},
			mockCalls: func() {
				mockRepo.On("IsUsernameExist", mock.Anything, "testUser").Return(false, nil)
				mockRepo.On("Insert", mock.Anything, mock.Anything).Return(nil)
				mockRepo.On("GetByUsername", mock.Anything, "testUser").Return(entity.User{ID: 1}, nil)

				// Add this mock expectation
				mockHasher.On("GenerateFromPassword", mock.Anything, mock.Anything).Return([]byte("hashedPassword"), nil)
			},
			wantErr: false,
		},
		{
			name: "duplicate username",
			req: request.SignUp{
				Username:  "duplicateUser",
				Password:  "testPass",
				FirstName: "deliiii",
				LastName:  "qobadiiii",
				Email:     "majesticdelaram@yahoo.com",
				Cellphone: "2398483984",
			},
			mockCalls: func() {
				mockRepo.On("IsUsernameExist", mock.Anything, "duplicateUser").Return(true, nil)
			},
			wantErr: true,
			errType: derror.NewBadRequestError(message.DuplicateUsername),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockCalls()

			_, err := service.SignUp(context.Background(), tt.req)
			if tt.wantErr {
				assert.Equal(t, tt.errType, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestGenerateJWTToken(t *testing.T) {
	service, _, _ := setup()

	user := entity.User{
		ID: 1,
	}

	token, err := service.generateJWTToken(context.Background(), user, time.Minute*15)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestGetUserByUsername(t *testing.T) {
	service, mockRepo, _ := setup()

	mockRepo.On("GetByUsername", mock.Anything, "testUser").Return(entity.User{Username: "testUser"}, nil) // Assuming no error and valid user
	user, err := service.getUserByUsername(context.Background(), "testUser")
	assert.NoError(t, err)
	assert.Equal(t, "testUser", user.Username)
	mockRepo.AssertExpectations(t)
}

func TestInsertUser(t *testing.T) {
	service, mockRepo, _ := setup()

	signUpReq := request.SignUp{
		Username:  "testUser",
		FirstName: "Test",
		LastName:  "User",
		Email:     "test@user.com",
		Cellphone: "1234567890",
	}

	hashedPassword := []byte("hashedPassword")

	mockRepo.On("Insert", mock.Anything, mock.Anything).Return(nil) // Assuming no error
	err := service.insertUser(context.Background(), signUpReq, hashedPassword)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestEnsureUsernameIsUnique(t *testing.T) {
	service, mockRepo, _ := setup()

	testCases := []struct {
		desc          string
		username      string
		repoError     error
		repoResponse  bool
		expectedError error
	}{
		{
			desc:          "username unique",
			username:      "uniqueUser",
			repoResponse:  false,
			expectedError: nil,
		},
		{
			desc:          "username exists",
			username:      "existingUser",
			repoResponse:  true,
			expectedError: derror.NewBadRequestError(message.DuplicateUsername),
		},
		{
			desc:          "repo error",
			username:      "errorUser",
			repoError:     fmt.Errorf("db error"),
			expectedError: derror.NewInternalSystemError(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			mockRepo.On("IsUsernameExist", mock.Anything, tc.username).Return(tc.repoResponse, tc.repoError)
			err := service.ensureUsernameIsUnique(context.Background(), tc.username)
			assert.Equal(t, tc.expectedError, err)
			mockRepo.AssertExpectations(t)
		})
	}
}


func TestService_SignIn(t *testing.T) {

	service, mockService, mockHasher := setup()

	// Define the user you're "expecting" to get from the mock repository
	expectedUser := entity.User{
		Username: "testUser",
		Password: "hashed_testPass",
		// ... other fields...
	}

	// Create a SignIn request
	signInReq := request.SignIn{
		Username: "testUser",
		Password: "testPass",
		// ... set other fields if there are any ...
	}

	// Set up expectations
	mockService.On("GetByUsername", mock.Anything, "testUser").Return(expectedUser, nil)
	mockService.On("CompareHashAndPassword", []byte("hashed_testPass"), []byte("testPass")).Return(nil)

	// Call the SignIn method
	user, err := service.SignIn(context.TODO(), signInReq)

	// Asserts
	assert.NoError(t, err) // Assuming no error is expected
	assert.Equal(t, expectedUser, user)

	// Verify that the expected methods were called on the mock objects
	mockService.AssertExpectations(t)
	mockHasher.AssertExpectations(t)
}



func TestValidateUserCredentials(t *testing.T) {
	service, mockRepo, mockHasher := setup()
	t.Run("valid user credentials", func(t *testing.T) {
		mockRepo.On("GetByUsername", context.Background(), "testUser").Return(entity.User{Username: "testUser", Password: "$hashedPasswordHere$"}, nil)

		mockHasher.On("CompareHashAndPassword", mock.Anything, []byte("testPassword")).Return(nil)

		user, err := service.validateUserCredentials(context.Background(), "testUser", "testPassword")
		assert.Nil(t, err)
		assert.Equal(t, "testUser", user.Username)
	})

	t.Run("user not found", func(t *testing.T) {
		mockRepo.On("GetByUsername", context.Background(), "testUser").Return(entity.User{}, derror.NewNotFoundError("user not found"))
		mockHasher.On("CompareHashAndPassword", mock.Anything, []byte("wrongPassword")).Return(errors.New("password mismatch"))

		_, err := service.validateUserCredentials(context.Background(), "testUser", "wrongPassword")
		assert.NotNil(t, err)
		assert.True(t, derror.IsHTTPError(err, http.StatusBadRequest))
	})

	t.Run("system error during user retrieval", func(t *testing.T) {
		mockRepo.On("GetByUsername", context.Background(), "testUser").Return(entity.User{}, errors.New("database error"))

		_, err := service.validateUserCredentials(context.Background(), "testUser", "testPassword")
		assert.NotNil(t, err)
		assert.True(t, derror.IsHTTPError(err, http.StatusInternalServerError))
	})

	t.Run("retrieved user's password is empty", func(t *testing.T) {
		mockRepo.On("GetByUsername", context.Background(), "testUser").Return(entity.User{Username: "testUser", Password: ""}, nil)

		_, err := service.validateUserCredentials(context.Background(), "testUser", "testPassword")
		assert.NotNil(t, err)
		assert.True(t, derror.IsHTTPError(err, http.StatusInternalServerError))
	})

	t.Run("provided password is incorrect", func(t *testing.T) {
		mockRepo.On("GetByUsername", context.Background(), "testUser").Return(entity.User{Username: "testUser", Password: "$hashedPasswordHere$"}, nil)

		_, err := service.validateUserCredentials(context.Background(), "testUser", "wrongPassword")
		assert.NotNil(t, err)
		assert.True(t, derror.IsHTTPError(err, http.StatusBadRequest))
	})

	mockRepo.AssertExpectations(t)
	mockHasher.AssertExpectations(t)

}

func TestGenerateUserAccessToken(t *testing.T) {
	service, _, _ := setup()

	user := entity.User{
		Username: "testUser",
	}

	token, err := service.generateUserAccessToken(context.Background(), user)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestGenerateUserRefreshToken(t *testing.T) {
	service, _, _ := setup()

	user := entity.User{
		Username: "testUser",
	}

	token, err := service.generateUserRefreshToken(context.Background(), user)
	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}
