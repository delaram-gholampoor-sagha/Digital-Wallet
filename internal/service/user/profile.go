package user

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror/message"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) GetProfile(ctx context.Context, userID int) (response.GetProfile, error) {
	user, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		if derror.IsNotFound(err) {
			return response.GetProfile{}, derror.NewNotFoundError(message.UserNotFound)
		}
		s.logger.Errorw("service.user.GetProfile.userRepo.Get", "error", err)
		return response.GetProfile{}, derror.NewInternalSystemError()
	}

	profile := response.GetProfile{
		Username:           user.Username,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		Email:              user.Email,
		ValidatedEmail:     user.ValidatedEmail,
		Cellphone:          user.Cellphone,
		ValidatedCellphone: user.ValidatedCellphone,
		CreatedAt:          user.CreatedAt.Unix(),
	}

	return profile, nil
}

func (s *Service) EditProfile(ctx context.Context, req request.EditProfile) error {
	if err := req.Validate(); err != nil {
		return err
	}

	user, err := s.userRepo.Get(ctx, req.UserID)
	if err != nil {
		if derror.IsNotFound(err) {
			return derror.NewBadRequestError(message.UserNotFound)
		}
		s.logger.Errorw("service.user.EditProfile.userRepo.Get", "error", err)
		return derror.NewInternalSystemError()
	}

	if req.Password != "" {
		password, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
		if err != nil {
			s.logger.Errorw("service.user.EditProfile.bcrypt.GenerateFromPassword", "error", err)
			return derror.NewInternalSystemError()
		}
		user.Password = string(password)
	}

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}

	if req.LastName != "" {
		user.LastName = req.LastName
	}

	if req.Email != "" {
		user.Email = req.Email
		user.ValidatedEmail = false
	}

	if req.Cellphone != "" {
		user.Cellphone = req.Cellphone
		user.ValidatedCellphone = false
	}

	if err := s.userRepo.Update(ctx, user); err != nil {
		s.logger.Errorw("service.user.EditProfile.userRepo.Update", "error", err, "user_id", user.ID)
		return derror.NewInternalSystemError()
	}

	return nil
}
