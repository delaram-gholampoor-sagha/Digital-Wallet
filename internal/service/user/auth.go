package user

import (
	"context"
	"fmt"
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror/message"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *Service) SignUp(ctx context.Context, req request.SignUp) (response.SignUp, error) {
	if err := req.Validate(); err != nil {
		return response.SignUp{}, err
	}

	// TODO: concurrent safe ??
	if err := s.ensureUsernameIsUnique(ctx, req.Username); err != nil {
		return response.SignUp{}, err
	}

	hashedPassword, err := s.hasher.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Errorw("service.user.SignUp.bcrypt.GenerateFromPassword", "error", err.Error())
		return response.SignUp{}, derror.NewInternalSystemError()
	}

	if err := s.insertUser(ctx, req, hashedPassword); err != nil {
		return response.SignUp{}, err
	}

	user, err := s.getUserByUsername(ctx, req.Username)
	if err != nil {
		return response.SignUp{}, err
	}

	accessToken, err := s.generateJWTToken(ctx, user, s.cfg.AccessTokenExp)
	if err != nil {
		return response.SignUp{}, err
	}

	refreshToken, err := s.generateJWTToken(ctx, user, s.cfg.RefreshTokenExp)
	if err != nil {
		return response.SignUp{}, err
	}

	return response.SignUp{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *Service) ensureUsernameIsUnique(ctx context.Context, username string) error {
	exists, err := s.userRepo.IsUsernameExist(ctx, username)
	if err != nil {
		s.logger.Errorw("service.user.SignUp.userRepo.IsUsernameExist", "error", err.Error())
		return derror.NewInternalSystemError()
	}
	if exists {
		return derror.NewBadRequestError(message.DuplicateUsername)
	}
	return nil
}

func (s *Service) insertUser(ctx context.Context, req request.SignUp, hashedPassword []byte) error {
	user := entity.User{
		Username:           req.Username,
		Password:           string(hashedPassword),
		FirstName:          req.FirstName,
		LastName:           req.LastName,
		Email:              req.Email,
		ValidatedEmail:     false,
		Cellphone:          req.Cellphone,
		ValidatedCellphone: false,
		Admin:              false,
	}

	if err := s.userRepo.Insert(ctx, user); err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}

func (s *Service) getUserByUsername(ctx context.Context, username string) (entity.User, error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return entity.User{}, fmt.Errorf("failed to get user by username: %w", err)
	}
	return user, nil
}

func (s *Service) generateJWTToken(ctx context.Context, user entity.User, exp time.Duration) (string, error) {
	claims := entity.JWTClaims{
		UserID: user.ID,
		Admin:  false,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	return s.tokenGen.GenerateToken(claims, s.cfg.Secret)
}

func (s *Service) SignIn(ctx context.Context, req request.SignIn) (response.SignIn, error) {
	if err := req.Validate(); err != nil {
		return response.SignIn{}, err
	}

	user, err := s.validateUserCredentials(ctx, req.Username, req.Password)
	if err != nil {
		return response.SignIn{}, err
	}

	accessToken, err := s.generateUserAccessToken(ctx, user)
	if err != nil {
		return response.SignIn{}, err
	}

	refreshToken, err := s.generateUserRefreshToken(ctx, user)
	if err != nil {
		return response.SignIn{}, err
	}

	return response.SignIn{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *Service) validateUserCredentials(ctx context.Context, username, password string) (entity.User, error) {
	user, err := s.getUserByUsername(ctx, username)
	if err != nil {
		if derror.IsNotFound(err) {
			return entity.User{}, derror.NewBadRequestError(message.InvalidUsernameOrPassword)
		}
		s.logger.Errorw("Failed to retrieve user by username", "username", username, "error", err.Error())
		return entity.User{}, derror.NewInternalSystemError()
	}

	// Check if user.Password is empty
	if user.Password == "" {
		s.logger.Errorw("Retrieved user password is empty", "username", username)
		return entity.User{}, derror.NewInternalSystemError() // or any appropriate error message indicating a data integrity issue
	}

	if err := s.hasher.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		s.logger.Errorw("Failed password hash comparison", "username", username)
		return entity.User{}, derror.NewBadRequestError(message.InvalidUsernameOrPassword)
	}
	return user, nil
}

func (s *Service) generateUserAccessToken(ctx context.Context, user entity.User) (string, error) {
	return s.generateJWTToken(ctx, user, s.cfg.AccessTokenExp)
}

func (s *Service) generateUserRefreshToken(ctx context.Context, user entity.User) (string, error) {
	return s.generateJWTToken(ctx, user, s.cfg.RefreshTokenExp)
}

func (s *Service) RefreshToken(ctx context.Context, userID int) (response.RefreshToken, error) {

	user, err := s.getUserByID(ctx, userID)
	if err != nil {
		return response.RefreshToken{}, err
	}

	accessToken, err := s.generateUserAccessToken(ctx, user)
	if err != nil {
		return response.RefreshToken{}, err
	}

	return response.RefreshToken{AccessToken: accessToken}, nil

}

// A new function to retrieve user by ID
func (s *Service) getUserByID(ctx context.Context, userID int) (entity.User, error) {
	user, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		s.logger.Errorw("service.user.getUserByID", "error", err.Error())
		return entity.User{}, derror.NewInternalSystemError()
	}
	return user, nil
}
