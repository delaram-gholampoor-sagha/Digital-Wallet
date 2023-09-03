package financialcard

import (
	"context"
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
	"go.uber.org/zap"
)

const registerCardMethod = "RegisterCard"
const updateCardMethod = "UpdateCard"

func (s *Service) RegisterCard(ctx context.Context, req *request.RegisterFinancialCard) (int, error) {
	if err := req.Validate(); err != nil {
		s.logger.Error("Failed to validate request",
			zap.Error(err),
			zap.String("method", registerCardMethod),
		)
		return 0, derror.NewBadRequestError("Invalid request parameters: %v", err)
	}

	exists, err := s.financialAccountService.IsAccountExist(ctx, req.AccountID)
	if err != nil {
		s.logger.Error("Failed to check if account exists",
			zap.Error(err),
			zap.String("method", registerCardMethod),
		)
		return 0, derror.NewInternalSystemError()
	}
	if !exists {
		return 0, derror.NewBadRequestError("Account ID does not exist")
	}

	card := entity.FinancialCard{
		AccountID:      req.AccountID,
		CardNumber:     req.CardNumber,
		CardType:       req.CardType,
		ExpirationDate: req.ExpirationDate,
		CardHolderName: req.CardHolderName,
		CVV:            req.CVV,
		Status:         req.Status,
		IssuedDate:     time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	if err := s.financialCardRepo.Insert(ctx, &card); err != nil {
		s.logger.Error("Failed to register card",
			zap.Error(err),
			zap.String("method", registerCardMethod),
		)
		return 0, derror.NewInternalSystemError()
	}

	return card.CardID, nil
}

func (s *Service) UpdateCard(ctx context.Context, req *request.UpdateFinancialCard) error {
	if err := req.Validate(); err != nil {
		s.logger.Error("Failed to validate request",
			zap.Error(err),
			zap.String("method", updateCardMethod),
		)
		return derror.NewBadRequestError("Invalid request parameters: %v", err)
	}

	card, err := s.financialCardRepo.GetByID(ctx, req.CardID)
	if err != nil {
		s.logger.Error("Failed to get card by ID",
			zap.Error(err),
			zap.String("method", updateCardMethod),
		)
		return derror.NewInternalSystemError()
	}

	if card == nil {
		return derror.NewNotFoundError("Card not found")
	}

	card.CardNumber = req.CardNumber
	card.CardType = req.CardType
	card.ExpirationDate = req.ExpirationDate
	card.CardHolderName = req.CardHolderName
	card.CVV = req.CVV
	card.Status = req.Status
	card.UpdatedAt = time.Now()

	if err := s.financialCardRepo.Update(ctx, card); err != nil {
		s.logger.Error("Failed to update card",
			zap.Error(err),
			zap.String("method", updateCardMethod),
		)
		return derror.NewInternalSystemError()
	}

	return nil
}

const deleteCardMethod = "DeleteCard"

func (s *Service) DeleteCard(ctx context.Context, cardID int64) error {
	card, err := s.financialCardRepo.GetByID(ctx, cardID)
	if err != nil {
		s.logger.Error("Failed to get card by ID",
			zap.Error(err),
			zap.String("method", deleteCardMethod),
		)
		return derror.NewInternalSystemError()
	}

	if card == nil {
		return derror.NewNotFoundError("Card not found")
	}

	if err := s.financialCardRepo.Delete(ctx, cardID); err != nil {
		s.logger.Error("Failed to delete card",
			zap.Error(err),
			zap.String("method", deleteCardMethod),
		)
		return derror.NewInternalSystemError()
	}

	return nil
}

const getCardByIDMethod = "GetCardByID"

func (s *Service) GetCardByID(ctx context.Context, cardID int64) (*entity.FinancialCard, error) {
	card, err := s.financialCardRepo.GetByID(ctx, cardID)
	if err != nil {
		s.logger.Error("Failed to get card by ID",
			zap.Error(err),
			zap.String("method", getCardByIDMethod),
		)
		return nil, derror.NewInternalSystemError()
	}

	if card == nil {
		return nil, derror.NewNotFoundError("Card not found")
	}

	return card, nil
}

const listCardsByAccountIDMethod = "ListCardsByAccountID"

func (s *Service) ListCardsByAccountID(ctx context.Context, accountID int) ([]*entity.FinancialCard, error) {
	cards, err := s.financialCardRepo.ListByAccountID(ctx, accountID)
	if err != nil {
		s.logger.Error("Failed to list cards by Account ID",
			zap.Error(err),
			zap.String("method", listCardsByAccountIDMethod),
		)
		return nil, derror.NewInternalSystemError()
	}
	return cards, nil
}

const listCardsByTypeMethod = "ListCardsByType"

func (s *Service) ListCardsByType(ctx context.Context, cardType enum.FinancialCardType) ([]*entity.FinancialCard, error) {
	cards, err := s.financialCardRepo.ListByCardType(ctx, cardType)
	if err != nil {
		s.logger.Error("Failed to list cards by type",
			zap.Error(err),
			zap.String("method", listCardsByTypeMethod),
		)
		return nil, derror.NewInternalSystemError()
	}
	return cards, nil
}
