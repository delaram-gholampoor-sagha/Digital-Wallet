package protocol

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
)

type FinancialCard interface {
	RegisterCard(ctx context.Context, req *request.RegisterFinancialCard) (Id int, err error)
	UpdateCard(ctx context.Context, req *request.UpdateFinancialCard) error
	DeleteCard(ctx context.Context, cardID int64) error
	GetCardByID(ctx context.Context, cardID int64) (*entity.FinancialCard, error)
	ListCardsByAccountID(ctx context.Context, accountID int) ([]*entity.FinancialCard, error)
	ListCardsByType(ctx context.Context, cardType enum.FinancialCardType) ([]*entity.FinancialCard, error)
}

type FinancialCardRepository interface {
	Insert(ctx context.Context, card *entity.FinancialCard) error
	Update(ctx context.Context, card *entity.FinancialCard) error
	Delete(ctx context.Context, cardID int64) error
	GetByID(ctx context.Context, cardID int64) (*entity.FinancialCard, error)
	ListByAccountID(ctx context.Context, accountID int) ([]*entity.FinancialCard, error)
	ListByCardType(ctx context.Context, cardType enum.FinancialCardType) ([]*entity.FinancialCard, error)
}
