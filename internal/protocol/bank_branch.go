package protocol

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
)

type BankBranch interface {
	AddBranch(ctx context.Context, req request.AddBranch) (response.RegisterBranch, error)
	GetBranchByID(ctx context.Context, branchID int) (response.GetBranch, error)
	GetBranchByName(ctx context.Context, branchName string) (response.GetBranch, error)
	GetBranchByCode(ctx context.Context, branchCode string) (response.GetBranch, error)
	IsBankBranchExist(ctx context.Context, bankBranchID int) (bool, error)
	UpdateBranch(ctx context.Context, req request.UpdateBranch) error
	DeleteBranch(ctx context.Context, branchID int) error
	ListAllBranches(ctx context.Context) ([]*entity.BankBranch, error)
	ListBranchesByStatus(ctx context.Context, status enum.BankBranchStatus) ([]*entity.BankBranch, error)
	ListBranchesByBankID(ctx context.Context, bankID int) ([]*entity.BankBranch, error)
}

type BankBranchRepository interface {
	GetByID(ctx context.Context, branchID int) (*entity.BankBranch, error)
	GetByName(ctx context.Context, branchName string) (*entity.BankBranch, error)
	GetByCode(ctx context.Context, branchCode string) (*entity.BankBranch, error)
	Insert(ctx context.Context, branch *entity.BankBranch) error
	Update(ctx context.Context, branch *entity.BankBranch) error
	Delete(ctx context.Context, branchID int) error
	ListAll(ctx context.Context) ([]*entity.BankBranch, error)
	ListByStatus(ctx context.Context, status enum.BankBranchStatus) ([]*entity.BankBranch, error)
	ListByBankID(ctx context.Context, bankID int) ([]*entity.BankBranch, error)
	IsBranchCodeExist(ctx context.Context, branchCode string) (bool, error)
	IsBankBranchExist(ctx context.Context, bankBranchID int) (bool, error)
}
