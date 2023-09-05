package protocol

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Get(ctx context.Context, id int) (entity.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) Insert(ctx context.Context, u entity.User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

func (m *MockUserRepository) Update(ctx context.Context, u entity.User) error {
	args := m.Called(ctx, u)
	return args.Error(0)
}

func (m *MockUserRepository) IsUsernameExist(ctx context.Context, username string) (bool, error) {
	args := m.Called(ctx, username)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) IsExist(ctx context.Context, id int) (bool, error) {
	args := m.Called(ctx, id)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) ValidateUserCredentials(ctx context.Context, username, password string) (entity.User, error) {
	args := m.Called(ctx, username, password)
	return args.Get(0).(entity.User), args.Error(1)
}

func (m *MockUserRepository) GenerateUserAccessToken(ctx context.Context, user entity.User) (string, error) {
	args := m.Called(ctx, user)
	return args.String(0), args.Error(1)
}

func (m *MockUserRepository) GenerateUserRefreshToken(ctx context.Context, user entity.User) (string, error) {
	args := m.Called(ctx, user)
	return args.String(0), args.Error(1)
}

func (m *MockUserRepository) SignUp(ctx context.Context, req request.SignUp) (response.SignUp, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(response.SignUp), args.Error(1)
}

func (m *MockUserRepository) SignIn(ctx context.Context, req request.SignIn) (response.SignIn, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(response.SignIn), args.Error(1)
}

func (m *MockUserRepository) RefreshToken(ctx context.Context, userID int) (response.RefreshToken, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(response.RefreshToken), args.Error(1)
}

func (m *MockUserRepository) GetProfile(ctx context.Context, userID int) (response.GetProfile, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).(response.GetProfile), args.Error(1)
}

func (m *MockUserRepository) EditProfile(ctx context.Context, req request.EditProfile) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockUserRepository) IsUserExist(ctx context.Context, userID int) (bool, error) {
	args := m.Called(ctx, userID)
	return args.Bool(0), args.Error(1)
}
