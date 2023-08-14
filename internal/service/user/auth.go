package user

import (
	"context"
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

	exist, err := s.userRepo.IsUsernameExist(ctx, req.Username)
	if err != nil {
		s.logger.Errorw("service.user.SignUp.userRepo.IsUsernameExist", "error", err.Error())
		return response.SignUp{}, derror.NewInternalSystemError()
	}

	if exist {
		return response.SignUp{}, derror.NewBadRequestError(message.DuplicateUsername)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Errorw("service.user.SignUp.bcrypt.GenerateFromPassword", "error", err.Error())
		return response.SignUp{}, derror.NewInternalSystemError()
	}

	user := entity.User{
		Username:           req.Username,
		Password:           string(password),
		FirstName:          req.FirstName,
		LastName:           req.LastName,
		Email:              req.Email,
		ValidatedEmail:     false,
		Cellphone:          req.Cellphone,
		ValidatedCellphone: false,
		Admin:              false,
	}

	if err := s.userRepo.Insert(ctx, user); err != nil {
		s.logger.Errorw("service.user.SignUp.userRepo.Create", "error", err.Error())
		return response.SignUp{}, derror.NewInternalSystemError()
	}

	user, err = s.userRepo.GetByUsername(ctx, user.Username)
	if err != nil {
		s.logger.Errorw("service.user.SignUp.userRepo.GetByUsername", "error", err.Error())
		return response.SignUp{}, derror.NewInternalSystemError()
	}

	accessClaims := entity.JWTClaims{
		UserID: user.ID,
		Admin:  false,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.AccessTokenExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	accessToken, err := generateToken(accessClaims, s.cfg.Secret)
	if err != nil {
		s.logger.Errorw("service.user.SignUp.jwt.GenerateToken", "error", err.Error(), "user_id", user.ID)
		return response.SignUp{}, derror.NewInternalSystemError()
	}

	refreshClaims := entity.JWTClaims{
		UserID: user.ID,
		Admin:  false,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.RefreshTokenExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	refreshToken, err := generateToken(refreshClaims, s.cfg.Secret)
	if err != nil {
		s.logger.Errorw("service.user.SignUp.jwt.GenerateToken", "error", err.Error(), "user_id", user.ID)
		return response.SignUp{}, derror.NewInternalSystemError()
	}

	return response.SignUp{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *Service) SignIn(ctx context.Context, req request.SignIn) (response.SignIn, error) {
	if err := req.Validate(); err != nil {
		return response.SignIn{}, err
	}

	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		if derror.IsNotFound(err) {
			return response.SignIn{}, derror.NewBadRequestError(message.InvalidUsernameOrPassword)
		}
		s.logger.Errorw("service.user.SignIn.userRepo.GetByUsername", "error", err.Error())
		return response.SignIn{}, derror.NewInternalSystemError()
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return response.SignIn{}, derror.NewBadRequestError(message.InvalidUsernameOrPassword)
	}

	accessClaims := entity.JWTClaims{
		UserID: user.ID,
		Admin:  user.Admin,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.AccessTokenExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken, err := generateToken(accessClaims, s.cfg.Secret)
	if err != nil {
		s.logger.Errorw("service.user.SignIn.jwt.GenerateToken", "error", err.Error(), "user_id", user.ID)
		return response.SignIn{}, derror.NewInternalSystemError()
	}

	refreshClaims := entity.JWTClaims{
		UserID: user.ID,
		Admin:  user.Admin,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.RefreshTokenExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	refreshToken, err := generateToken(refreshClaims, s.cfg.Secret)
	if err != nil {
		s.logger.Errorw("service.user.SignIn.jwt.GenerateToken", "error", err.Error(), "user_id", user.ID)
		return response.SignIn{}, derror.NewInternalSystemError()
	}

	return response.SignIn{AccessToken: accessToken, RefreshToken: refreshToken}, nil
}

func (s *Service) RefreshToken(ctx context.Context, userID int) (response.RefreshToken, error) {
	user, err := s.userRepo.Get(ctx, userID)
	if err != nil {
		if derror.IsNotFound(err) {
			return response.RefreshToken{}, derror.NewNotFoundError(message.UserNotFound)
		}
		s.logger.Errorw("service.user.RefreshToken.userRepo.Get", "error", err, "user_id", userID)
		return response.RefreshToken{}, derror.NewInternalSystemError()
	}

	accessClaims := entity.JWTClaims{
		UserID: user.ID,
		Admin:  user.Admin,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        uuid.New().String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.cfg.AccessTokenExp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	accessToken, err := generateToken(accessClaims, s.cfg.Secret)
	if err != nil {
		s.logger.Errorw("service.user.RefreshToken.jwt.GenerateToken", "error", err.Error(), "user_id", user.ID)
		return response.RefreshToken{}, derror.NewInternalSystemError()
	}

	return response.RefreshToken{AccessToken: accessToken}, nil
}

func generateToken(claims entity.JWTClaims, secret string) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
}
