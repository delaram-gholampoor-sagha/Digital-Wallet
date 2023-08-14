package protocol

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
)

type User interface {
	SignUp(ctx context.Context, req request.SignUp) (response.SignUp, error)
	SignIn(ctx context.Context, req request.SignIn) (response.SignIn, error)
	RefreshToken(ctx context.Context, userID int) (response.RefreshToken, error)
	GetProfile(ctx context.Context, userID int) (response.GetProfile, error)
	EditProfile(ctx context.Context, req request.EditProfile) error
	Get(ctx context.Context, id int) (entity.User, error)
}

type UserRepository interface {
	Get(ctx context.Context, id int) (entity.User, error)
	GetByUsername(ctx context.Context, username string) (entity.User, error)
	Insert(ctx context.Context, u entity.User) error
	Update(ctx context.Context, u entity.User) error
	IsUsernameExist(ctx context.Context, username string) (bool, error)
	IsExist(ctx context.Context, id int) (bool, error)
}
