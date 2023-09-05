package user

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
)

func (s *Service) Get(ctx context.Context, id int) (entity.User, error) {
	//TODO: implement me
	panic("implement me")
}

func (s *Service) IsUserExist(ctx context.Context, userID int) (bool, error) {
	//TODO: implement me
	// panic("implement me")

	return false, nil
}
