package bankbranch

import (
	"context"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
	"go.uber.org/zap"
)

func (s *Service) AddBranch(ctx context.Context, req request.AddBranch) (response.RegisterBranch, error) {

	if err := req.Validate(); err != nil {
		s.logger.Error("Failed to validate request",
			zap.Error(err),
			zap.String("method", "AddBranch"),
		)
		return response.RegisterBranch{}, derror.NewBadRequestError("Invalid request parameters: %v", err)
	}

	exists, err := s.bank.IsBankExist(ctx, req.BankID)
	if err != nil {
		s.logger.Error("Failed to check if bank exists",
			zap.Error(err),
			zap.String("method", "AddBranch"),
		)
		return response.RegisterBranch{}, derror.NewInternalSystemError()
	}
	if !exists {
		return response.RegisterBranch{}, derror.NewBadRequestError("Bank ID does not exist")
	}

	exists, err = s.bankBranchRepo.IsBranchCodeExist(ctx, req.BranchCode)
	if err != nil {
		s.logger.Error("Failed to check if branch code exists",
			zap.Error(err),
			zap.String("method", "AddBranch"),
		)
		return response.RegisterBranch{}, derror.NewInternalSystemError()
	}
	if exists {
		return response.RegisterBranch{}, derror.NewBadRequestError("Branch code already exists")
	}

	branch := entity.BankBranch{
		BankID:      req.BankID,
		BranchName:  req.BranchName,
		BranchCode:  req.BranchCode,
		Address:     req.Address,
		City:        req.City,
		Province:    req.Province,
		PostalCode:  req.PostalCode,
		PhoneNumber: req.PhoneNumber,
		Status:      req.Status,
	}

	if err := s.bankBranchRepo.Insert(ctx, &branch); err != nil {
		s.logger.Error("Failed to insert branch",
			zap.Error(err),
			zap.String("method", "AddBranch"),
		)
		return response.RegisterBranch{}, derror.NewInternalSystemError()
	}

	return response.RegisterBranch{BranchID: branch.BranchID}, nil
}

func (s *Service) GetBranchByID(ctx context.Context, branchID int) (response.GetBranch, error) {
	branch, err := s.bankBranchRepo.GetByID(ctx, branchID)
	if err != nil {
		s.logger.Error("Failed to get branch by ID",
			zap.Error(err),
			zap.String("method", "GetBranchByID"),
		)
		return response.GetBranch{}, derror.NewInternalSystemError()
	}
	if branch == nil {
		return response.GetBranch{}, derror.NewNotFoundError("Branch not found")
	}

	res := response.GetBranch{
		BranchID:    branch.BranchID,
		BankID:      branch.BankID,
		BranchName:  branch.BranchName,
		BranchCode:  branch.BranchCode,
		Address:     branch.Address,
		City:        branch.City,
		Province:    branch.Province,
		PostalCode:  branch.PostalCode,
		PhoneNumber: branch.PhoneNumber,
		Status:      branch.Status,
		CreatedAt:   branch.CreatedAt,
		UpdatedAt:   branch.UpdatedAt,
		DeletedAt:   branch.DeletedAt,
	}

	return res, nil
}

const updateBranchMethod = "UpdateBranch"

func (s *Service) UpdateBranch(ctx context.Context, req request.UpdateBranch) error {

	if err := req.Validate(); err != nil {
		s.logger.Error("Failed to validate request",
			zap.Error(err),
			zap.String("method", "UpdateBranch"),
		)
		return derror.NewBadRequestError("Invalid request parameters: %v", err)
	}

	branch, err := s.bankBranchRepo.GetByID(ctx, req.BranchID)
	if err != nil {
		s.logger.Error("Failed to get branch by ID",
			zap.Error(err),
			zap.String("method", "UpdateBranch"),
		)
		return derror.NewInternalSystemError()
	}

	if branch == nil {
		return derror.NewNotFoundError("Branch not found")
	}

	branch.BankID = req.BankID
	branch.BranchName = req.BranchName
	branch.BranchCode = req.BranchCode
	branch.Address = req.Address
	branch.City = req.City
	branch.Province = req.Province
	branch.PostalCode = req.PostalCode
	branch.PhoneNumber = req.PhoneNumber
	branch.Status = req.Status

	if err := s.bankBranchRepo.Update(ctx, branch); err != nil {
		s.logger.Error("Failed to update branch",
			zap.Error(err),
			zap.String("method", "UpdateBranch"),
		)
		return derror.NewInternalSystemError()
	}
	return nil
}

const getBranchByNameMethod = "GetBranchByName"
const getBranchByCodeMethod = "GetBranchByCode"

func (s *Service) GetBranchByName(ctx context.Context, branchName string) (response.GetBranch, error) {

	branch, err := s.bankBranchRepo.GetByName(ctx, branchName)
	if err != nil {
		s.logger.Error("Failed to get branch by name",
			zap.Error(err),
			zap.String("method", getBranchByNameMethod),
		)
		return response.GetBranch{}, derror.NewInternalSystemError()
	}

	if branch == nil {
		return response.GetBranch{}, derror.NewNotFoundError("Branch not found")
	}

	return response.GetBranch{
		BranchID:    branch.BranchID,
		BankID:      branch.BankID,
		BranchName:  branch.BranchName,
		BranchCode:  branch.BranchCode,
		Address:     branch.Address,
		City:        branch.City,
		Province:    branch.Province,
		PostalCode:  branch.PostalCode,
		PhoneNumber: branch.PhoneNumber,
		Status:      branch.Status,
		CreatedAt:   branch.CreatedAt,
		UpdatedAt:   branch.UpdatedAt,
		DeletedAt:   branch.DeletedAt,
	}, nil
}

func (s *Service) GetBranchByCode(ctx context.Context, branchCode string) (response.GetBranch, error) {

	branch, err := s.bankBranchRepo.GetByCode(ctx, branchCode)
	if err != nil {
		s.logger.Error("Failed to get branch by code",
			zap.Error(err),
			zap.String("method", getBranchByCodeMethod),
		)
		return response.GetBranch{}, derror.NewInternalSystemError()
	}

	if branch == nil {
		return response.GetBranch{}, derror.NewNotFoundError("Branch not found")
	}

	return response.GetBranch{
		BranchID:    branch.BranchID,
		BankID:      branch.BankID,
		BranchName:  branch.BranchName,
		BranchCode:  branch.BranchCode,
		Address:     branch.Address,
		City:        branch.City,
		Province:    branch.Province,
		PostalCode:  branch.PostalCode,
		PhoneNumber: branch.PhoneNumber,
		Status:      branch.Status,
		CreatedAt:   branch.CreatedAt,
		UpdatedAt:   branch.UpdatedAt,
		DeletedAt:   branch.DeletedAt,
	}, nil
}

func (s *Service) ListAllBranches(ctx context.Context) ([]*entity.BankBranch, error) {
	branches, err := s.bankBranchRepo.ListAll(ctx)
	if err != nil {
		s.logger.Error("Failed to list all branches",
			zap.Error(err),
			zap.String("method", "ListAllBranches"),
		)
		return nil, derror.NewInternalSystemError()
	}
	return branches, nil
}

func (s *Service) ListBranchesByStatus(ctx context.Context, status enum.BankBranchStatus) ([]*entity.BankBranch, error) {
	branches, err := s.bankBranchRepo.ListByStatus(ctx, status)
	if err != nil {
		s.logger.Error("Failed to list branches by status",
			zap.Error(err),
			zap.String("method", "ListBranchesByStatus"),
		)
		return nil, derror.NewInternalSystemError()
	}
	return branches, nil
}

func (s *Service) ListBranchesByBankID(ctx context.Context, bankID int) ([]*entity.BankBranch, error) {
	branches, err := s.bankBranchRepo.ListByBankID(ctx, bankID)
	if err != nil {
		s.logger.Error("Failed to list branches by bank ID",
			zap.Error(err),
			zap.String("method", "ListBranchesByBankID"),
		)
		return nil, derror.NewInternalSystemError()
	}
	return branches, nil
}

const deleteBranchMethod = "DeleteBranch"

func (s *Service) DeleteBranch(ctx context.Context, branchID int) error {

	branch, err := s.bankBranchRepo.GetByID(ctx, branchID)
	if err != nil {
		s.logger.Error("Failed to get branch by ID",
			zap.Error(err),
			zap.String("method", deleteBranchMethod),
		)
		return derror.NewInternalSystemError()
	}

	if branch == nil {
		return derror.NewNotFoundError("Branch not found")
	}

	if err := s.bankBranchRepo.Delete(ctx, branchID); err != nil {
		s.logger.Error("Failed to delete branch",
			zap.Error(err),
			zap.String("method", deleteBranchMethod),
		)
		return derror.NewInternalSystemError()
	}

	return nil
}

func (s *Service) IsBankBranchExist(ctx context.Context, bankBranchID int) (bool, error) {

	return false, nil
}
