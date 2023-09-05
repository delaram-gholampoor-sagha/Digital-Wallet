package financialaccount

import (
	"context"
	"reflect"
	"regexp"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
	"go.uber.org/zap"
)

func (s *Service) CreateAccount(ctx context.Context, req request.RegisterFinancialAccount) (response.RegisterFinancialAccount, error) {

	if err := req.Validate(); err != nil {
		s.logger.Error("Failed to validate request", zap.Error(err), zap.String("method", "CreateAccount"))
		return response.RegisterFinancialAccount{}, derror.NewBadRequestError("Invalid request parameters: %v", err)
	}

	exists, err := s.bankService.IsBankExist(ctx, req.BankID)
	if err != nil {

		return response.RegisterFinancialAccount{}, derror.NewInternalSystemError()
	}
	if !exists {
		return response.RegisterFinancialAccount{}, derror.NewBadRequestError("Bank ID does not exist")
	}

	exists, err = s.bankBranchService.IsBankBranchExist(ctx, req.BranchID)
	if err != nil {

		return response.RegisterFinancialAccount{}, derror.NewInternalSystemError()
	}
	if !exists {
		return response.RegisterFinancialAccount{}, derror.NewBadRequestError("Branch ID does not exist")
	}

	exists, err = s.userService.IsUserExist(ctx, req.UserID)
	if err != nil {
		return response.RegisterFinancialAccount{}, derror.NewInternalSystemError()
	}
	if !exists {
		return response.RegisterFinancialAccount{}, derror.NewBadRequestError("User ID does not exist")
	}

	exists, err = s.currencyService.IsCurrrencyExist(ctx, req.CurrencyID)
	if err != nil {
		return response.RegisterFinancialAccount{}, derror.NewInternalSystemError()
	}
	if !exists {
		return response.RegisterFinancialAccount{}, derror.NewBadRequestError("Currency ID does not exist")
	}

	account := entity.FinancialAccount{
		UserID:        req.UserID,
		CurrencyID:    req.CurrencyID,
		BankID:        req.BankID,
		BranchID:      req.BranchID,
		AccountNumber: req.AccountNumber,
		ShabaNumber:   req.ShabaNumber,
		AccountName:   req.AccountName,
		AccountType:   req.AccountType,
		CurrencyCode:  req.CurrencyCode,
		Status:        req.Status,
	}

	if err := s.financialAccountRepo.Insert(ctx, &account); err != nil {
		s.logger.Error("Failed to insert account", zap.Error(err), zap.String("method", "CreateAccount"))
		return response.RegisterFinancialAccount{}, derror.NewInternalSystemError()
	}

	return response.RegisterFinancialAccount{AccountID: account.AccountID}, nil
}

func (s *Service) GetAccountByID(ctx context.Context, accountID int) (response.GetFinancialAccount, error) {

	// Input validation
	if accountID <= 0 {
		s.logger.Error("Invalid account ID", zap.Int("accountID", accountID), zap.String("method", "GetAccountByID"))
		return response.GetFinancialAccount{}, derror.NewBadRequestError("Invalid account ID")
	}

	// Authorization (simplified, implement as needed)
	// authorized, err := s.isAuthorized(ctx, accountID)
	// if err != nil {
	// 	s.logger.Error("Authorization check failed", zap.Error(err), zap.String("method", "GetAccountByID"))
	// 	return response.GetFinancialAccount{}, derror.NewInternalSystemError()
	// }
	// if !authorized {
	// 	return response.GetFinancialAccount{}, derror.NewUnauthorizedError("Not authorized to access this account")
	// }

	account, err := s.financialAccountRepo.GetByID(ctx, accountID)
	if err != nil {
		s.logger.Error("Failed to get account by ID", zap.Error(err), zap.String("method", "GetAccountByID"))
		return response.GetFinancialAccount{}, derror.NewInternalSystemError()
	}

	if account == nil {
		return response.GetFinancialAccount{}, derror.NewNotFoundError("Account not found")
	}

	return response.GetFinancialAccount{
		AccountID:     account.AccountID,
		UserID:        account.UserID,
		CurrencyID:    account.CurrencyID,
		BankID:        account.BankID,
		BranchID:      account.BranchID,
		AccountNumber: account.AccountNumber,
		ShabaNumber:   account.ShabaNumber,
		AccountName:   account.AccountName,
		AccountType:   account.AccountType,
		CurrencyCode:  account.CurrencyCode,
		Status:        account.Status,
		CreatedAt:     account.CreatedAt,
		UpdatedAt:     account.UpdatedAt,
		DeletedAt:     account.DeletedAt,
	}, nil
}

func (s *Service) IsAccountExist(ctx context.Context, accountID int) (bool, error) {

	if accountID <= 0 {
		s.logger.Error("Invalid account ID", zap.Int("accountID", accountID), zap.String("method", "DoesAccountExist"))
		return false, derror.NewBadRequestError("Invalid account ID")
	}

	exists, err := s.financialAccountRepo.IsAccountExist(ctx, accountID)
	if err != nil {
		s.logger.Error("Failed to check if account exists", zap.Error(err), zap.String("method", "IsAccountExist"))
		return false, derror.NewInternalSystemError()
	}
	return exists, nil
}

func (s *Service) UpdateAccount(ctx context.Context, req request.UpdateFinancialAccount) error {

	if err := req.Validate(); err != nil {
		s.logger.Error("Failed to validate request", zap.Error(err), zap.String("method", "UpdateAccount"))
		return derror.NewBadRequestError("Invalid request parameters: %v", err)
	}

	// Authorization and entity checks (simplified)
	// TODO: Implement this part according to your needs
	// authorized, relatedEntitiesExist := true, true
	// if !authorized || !relatedEntitiesExist {
	// 	return derror.NewUnauthorizedError("Unauthorized or related entities not found")
	// }

	existingAccount, err := s.financialAccountRepo.GetByID(ctx, req.AccountID)
	if err != nil {
		s.logger.Error("Failed to get existing account", zap.Error(err), zap.String("method", "UpdateAccount"))
		return derror.NewInternalSystemError()
	}

	if existingAccount == nil {
		return derror.NewNotFoundError("Account not found")
	}

	updatedAccount := *existingAccount // Make a copy
	updatedAccount.UserID = req.UserID
	updatedAccount.CurrencyID = req.CurrencyID
	updatedAccount.BankID = req.BankID
	updatedAccount.BranchID = req.BranchID
	updatedAccount.AccountNumber = req.AccountNumber
	updatedAccount.ShabaNumber = req.ShabaNumber
	updatedAccount.AccountName = req.AccountName
	updatedAccount.AccountType = req.AccountType
	updatedAccount.CurrencyCode = req.CurrencyCode
	updatedAccount.Status = req.Status

	// Only update if something has changed
	if !reflect.DeepEqual(*existingAccount, updatedAccount) {
		if err := s.financialAccountRepo.Update(ctx, &updatedAccount); err != nil {
			s.logger.Error("Failed to update account", zap.Error(err), zap.String("method", "UpdateAccount"))
			return derror.NewInternalSystemError()
		}
	}

	if err := s.financialAccountRepo.Update(ctx, existingAccount); err != nil {
		s.logger.Error("Failed to update account", zap.Error(err), zap.String("method", "UpdateAccount"))
		return derror.NewInternalSystemError()
	}
	return nil
}

func (s *Service) DeleteAccount(ctx context.Context, accountID int) error {
	if accountID <= 0 {
		s.logger.Error("Invalid account ID", zap.String("method", "DeleteAccount"))
		return derror.NewBadRequestError("Invalid account ID")
	}

	// Authorization check (skipped  but should be implemented)

	if err := s.financialAccountRepo.Delete(ctx, accountID); err != nil {
		s.logger.Error("Failed to delete account", zap.Error(err), zap.String("method", "DeleteAccount"))
		return derror.NewInternalSystemError()
	}
	return nil
}

func (s *Service) ListAccountsByUserID(ctx context.Context, userID int) ([]entity.FinancialAccount, error) {
	if userID <= 0 {
		s.logger.Error("Invalid user ID", zap.String("method", "ListAccountsByUserID"))
		return nil, derror.NewBadRequestError("Invalid user ID")
	}

	// Authorization check (skipped  but should be implemented)

	accounts, err := s.financialAccountRepo.GetByUserID(ctx, userID)
	if err != nil {
		s.logger.Error("Failed to list accounts", zap.Error(err), zap.String("method", "ListAccountsByUserID"))
		return nil, derror.NewInternalSystemError()
	}
	return accounts, nil
}

func (s *Service) ListAccountsByStatus(ctx context.Context, status enum.FinancialAccountStatus) ([]*entity.FinancialAccount, error) {

	// Validate status (assuming IsValid is a method that checks if the status is a valid enum)
	if !status.IsValid() {
		s.logger.Error("Invalid status", zap.String("status", status.String()), zap.String("method", "ListAccountsByStatus"))
		return nil, derror.NewBadRequestError("Invalid status")
	}
	accounts, err := s.financialAccountRepo.ListByStatus(ctx, status)
	if err != nil {
		s.logger.Error("Failed to list accounts by status", zap.Error(err), zap.String("method", "ListAccountsByStatus"))
		return nil, derror.NewInternalSystemError()
	}

	// currentHour := time.Now().Hour()
	// if 0 <= currentHour && currentHour <= 5 { // Between midnight and 5am
	// 	return []*entity.FinancialAccount{}, nil
	// } else {
	// 	return nil, derror.NewNotFoundError("No accounts found for the given status")
	// }

	// if status == "closed" {
	// 	return []*entity.FinancialAccount{}, nil
	// } else {
	// 	return nil, derror.NewNotFoundError("No accounts found for the given status")
	// }

	// if featureflags.IsEnabled("ReturnEmptyListIfNoAccounts") {
	// 	return []*entity.FinancialAccount{}, nil
	// } else {
	// 	return nil, derror.NewNotFoundError("No accounts found for the given status")
	// }

	// Decide whether to return an error or empty list when no accounts are found
	if len(accounts) == 0 {
		return []*entity.FinancialAccount{}, nil
	}

	return accounts, nil
}

func (s *Service) VerifyAccount(ctx context.Context, accountID int) error {
	if accountID <= 0 {
		s.logger.Error("Invalid account ID", zap.Int("accountID", accountID), zap.String("method", "VerifyAccount"))
		return derror.NewBadRequestError("Invalid account ID")
	}

	account, err := s.financialAccountRepo.GetByID(ctx, accountID)
	if err != nil {
		s.logger.Error("Failed to get account by ID", zap.Error(err), zap.String("method", "VerifyAccount"))
		return derror.NewInternalSystemError()
	}

	if account == nil {
		return derror.NewNotFoundError("Account not found")
	}

	if account.Status == enum.Verified {
		s.logger.Info("Account already verified", zap.Int("accountID", accountID), zap.String("method", "VerifyAccount"))
		return nil //TODO: or return appropriate message ??
	}

	if err := s.financialAccountRepo.UpdateStatus(ctx, accountID, "verified"); err != nil {
		s.logger.Error("Failed to update account status", zap.Error(err), zap.String("method", "VerifyAccount"))
		return derror.NewInternalSystemError()
	}

	return nil
}

func (s *Service) GetAccountByShaba(ctx context.Context, shabaNumber string) (response.GetFinancialAccount, error) {
	const shabaPattern = `^[a-zA-Z0-9]{26}$`
	isValid := regexp.MustCompile(shabaPattern).MatchString(shabaNumber)
	if !isValid {
		s.logger.Error("Invalid ShabaNumber", zap.String("method", "GetAccountByShaba"), zap.String("shabaNumber", shabaNumber))
		return response.GetFinancialAccount{}, derror.NewValidationError("Invalid Shaba number")
	}

	account, err := s.financialAccountRepo.GetByShabaNumber(ctx, shabaNumber)
	if err != nil {
		s.logger.Error("Failed to get account by ShabaNumber", zap.Error(err), zap.String("method", "GetAccountByShaba"))
		return response.GetFinancialAccount{}, derror.NewInternalSystemError()
	}

	if account == nil {
		return response.GetFinancialAccount{}, derror.NewNotFoundError("Account not found")
	}

	return response.GetFinancialAccount{
		AccountID:     account.AccountID,
		UserID:        account.UserID,
		CurrencyID:    account.CurrencyID,
		BankID:        account.BankID,
		BranchID:      account.BranchID,
		AccountNumber: account.AccountNumber,
		ShabaNumber:   account.ShabaNumber,
		AccountName:   account.AccountName,
		AccountType:   account.AccountType,
		CurrencyCode:  account.CurrencyCode,
		Status:        account.Status,
		CreatedAt:     account.CreatedAt,
		UpdatedAt:     account.UpdatedAt,
		DeletedAt:     account.DeletedAt,
	}, nil
}

func (s *Service) GetAccountTransactionHistory(ctx context.Context, accountID int) ([]*entity.FinancialAccountTransaction, error) {
	// Validation (replace this with your own validation logic if needed)
	if accountID <= 0 {
		s.logger.Error("Invalid account ID", zap.Int("accountID", accountID), zap.String("method", "GetAccountTransactionHistory"))
		return nil, derror.NewValidationError("Invalid account ID")
	}

	// TODO: Optional: I need to also add a check here to make sure the accountID belongs to the authenticated user.
	// ... (My authorization logic here)

	transactions, err := s.financialAccountTransactionService.ListTransactionsByAccountID(ctx, accountID)
	if err != nil {
		s.logger.Error("Failed to get transactions by account ID", zap.Error(err), zap.String("method", "GetAccountTransactionHistory"))
		return nil, derror.NewInternalSystemError()
	}

	if len(transactions) == 0 {
		return nil, derror.NewNotFoundError("No transactions found for the given account ID")
	}

	return transactions, nil
}

func (s *Service) ListAccountsByType(ctx context.Context, accountType string) ([]*entity.FinancialAccount, error) {
	// Validate accountType (assuming 'checking' and 'savings' are the only valid types)
	validTypes := map[string]bool{"checking": true, "savings": true}
	if !validTypes[accountType] {
		s.logger.Error("Invalid account type", zap.String("accountType", accountType), zap.String("method", "ListAccountsByType"))
		return nil, derror.NewValidationError("Invalid account type")
	}

	// TODO: Optional: I need to also add a check here to make sure the accountID belongs to the authenticated user.
	// ... (My authorization logic here)

	accounts, err := s.financialAccountRepo.ListByType(ctx, accountType)
	if err != nil {
		s.logger.Error("Failed to list accounts by type", zap.Error(err), zap.String("method", "ListAccountsByType"))
		return nil, derror.NewInternalSystemError()
	}

	if len(accounts) == 0 {
		return nil, derror.NewNotFoundError("No accounts found for the given type")
	}

	return accounts, nil
}

func (s *Service) GetAccountCurrency(ctx context.Context, accountID int) (response.GetCurrency, error) {

	if accountID <= 0 {
		s.logger.Error("Invalid account ID", zap.Int("accountID", accountID), zap.String("method", "GetAccountCurrency"))
		return response.GetCurrency{}, derror.NewValidationError("Invalid account ID")
	}

	account, err := s.financialAccountRepo.GetByID(ctx, accountID)
	if err != nil {
		s.logger.Error("Failed to get account by ID", zap.Error(err), zap.String("method", "GetAccountCurrency"))
		return response.GetCurrency{}, derror.NewInternalSystemError()
	}
	if account == nil {
		return response.GetCurrency{}, derror.NewNotFoundError("Account not found")
	}

	currency, err := s.currencyService.GetCurrency(ctx, account.CurrencyID)
	if err != nil {
		s.logger.Error("Failed to get currency by ID", zap.Error(err), zap.String("method", "GetAccountCurrency"))
		return response.GetCurrency{}, derror.NewInternalSystemError()
	}
	if currency.CurrencyID == 0 {
		return response.GetCurrency{}, derror.NewNotFoundError("Currency not found")
	}

	return response.GetCurrency{
		CurrencyID:   currency.CurrencyID,
		CurrencyCode: currency.CurrencyCode,
		CurrencyName: currency.CurrencyName,
		Symbol:       currency.Symbol,
		ExchangeRate: currency.ExchangeRate,
		CreatedAt:    currency.CreatedAt,
		UpdatedAt:    currency.UpdatedAt,
		DeletedAt:    currency.DeletedAt,
	}, nil
}

func (s *Service) GetBranchForAccount(ctx context.Context, accountID int) (response.GetBankBranch, error) {
	account, err := s.financialAccountRepo.GetByID(ctx, accountID)
	if err != nil {
		s.logger.Error("Failed to get account by ID", zap.Error(err), zap.String("method", "GetBranchForAccount"))
		return response.GetBankBranch{}, derror.NewInternalSystemError()
	}

	if account == nil || account.AccountID == 0 {
		return response.GetBankBranch{}, derror.NewNotFoundError("Account not found")
	}

	branch, err := s.financialAccountRepo.GetBranchByAccountID(ctx, accountID)
	if err != nil {
		s.logger.Error("Failed to get branch by account ID", zap.Error(err), zap.String("method", "GetBranchForAccount"))
		return response.GetBankBranch{}, derror.NewInternalSystemError()
	}

	if branch == nil || branch.BranchID == 0 {
		return response.GetBankBranch{}, derror.NewNotFoundError("Branch not found")
	}

	return response.GetBankBranch{
		BranchID:   branch.BranchID,
		BankID:     branch.BankID,
		BranchCode: branch.BranchCode,
		BranchName: branch.BranchName,
		Address:    branch.Address,
		CreatedAt:  branch.CreatedAt,
		UpdatedAt:  branch.UpdatedAt,
	}, nil
}
func (s *Service) GetBankForAccount(ctx context.Context, accountID int) (response.GetBank, error) {
	account, err := s.financialAccountRepo.GetByID(ctx, accountID)
	if err != nil {
		s.logger.Error("Failed to get account by ID", zap.Error(err), zap.String("method", "GetBankForAccount"))
		return response.GetBank{}, derror.NewInternalSystemError()
	}

	if account == nil || account.AccountID == 0 {
		return response.GetBank{}, derror.NewNotFoundError("Account not found")
	}

	bank, err := s.financialAccountRepo.GetBankByAccountID(ctx, accountID)
	if err != nil {
		s.logger.Error("Failed to get bank by account ID", zap.Error(err), zap.String("method", "GetBankForAccount"))
		return response.GetBank{}, derror.NewInternalSystemError()
	}

	if bank == nil || bank.BankID == 0 {
		return response.GetBank{}, derror.NewNotFoundError("Bank not found")
	}

	return response.GetBank{
		BankID:    bank.BankID,
		Name:      bank.Name,
		BankCode:  bank.BankCode,
		Status:    bank.Status,
		CreatedAt: bank.CreatedAt,
		UpdatedAt: bank.UpdatedAt,
		DeletedAt: bank.DeletedAt,
	}, nil
}
