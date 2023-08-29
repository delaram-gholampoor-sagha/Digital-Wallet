package user

import (
	"context"
	"reflect"
	"testing"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestGetProfile(t *testing.T) {
	mockRepo := new(protocol.MockUserRepository)
	svc := &Service{
		userRepo: mockRepo,
		logger:   zap.NewExample().Sugar(),
	}

	tests := []struct {
		name        string
		userID      int
		mockReturn  []interface{}
		want        response.GetProfile
		wantErr     bool
		wantErrType error
	}{
		// ... Your test cases
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the mock expectations for this particular test
			mockRepo.On("Get", mock.Anything, tt.userID).Return(tt.mockReturn[0], tt.mockReturn[1])

			profile, err := svc.GetProfile(context.TODO(), tt.userID)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetProfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				if !reflect.DeepEqual(err, tt.wantErrType) {
					t.Errorf("GetProfile() error = %v, wantErrType %v", err, tt.wantErrType)
				}
			}

			if !reflect.DeepEqual(profile, tt.want) {
				t.Errorf("GetProfile() got = %v, want %v", profile, tt.want)
			}

			mockRepo.AssertExpectations(t)
			mockRepo.ExpectedCalls = []*mock.Call{} // Clear the expectations
			mockRepo.Calls = []mock.Call{}          // Clear the recorded calls
		})
	}
}



func TestEditProfile(t *testing.T) {
	// Create a test context
	ctx := context.TODO()

	// Set up the dependencies
	service, mockRepo, _ := setup()

	// Set expectations on the mock repository
	mockRepo.On("Get", ctx, 42).Return(entity.User{ID: 42}, nil)
	mockRepo.On("Update", ctx, mock.AnythingOfType("entity.User")).Return(nil)

	// Call the function
	err := service.EditProfile(ctx, request.EditProfile{
		UserID:    42,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "johndoe@example.com",
		Cellphone: "1234567890",
		Password:  "newpassword",
	})

	// Assert that no error occurred
	assert.NoError(t, err)

	// Assert that the mock repository methods were called with the expected arguments
	mockRepo.AssertCalled(t, "Get", ctx, 42)
	mockRepo.AssertCalled(t, "Update", ctx, mock.AnythingOfType("entity.User"))

	// Assert any other expectations as needed
	// ...
}
