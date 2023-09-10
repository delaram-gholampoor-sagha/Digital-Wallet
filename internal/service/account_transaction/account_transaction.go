package accounttransaction

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/request"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol/response"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
	"go.uber.org/zap"
)

// RegisterTransaction registers a new account transaction
func (s *Service) RegisterTransaction(ctx context.Context, req *request.RegisterTransactionRequest) (*response.RegisterTransactionResponse, error) {
	s.logger.Info("Starting the process to register a new transaction")

	if err := req.Validate(); err != nil {
		s.logger.Error("Failed to validate request", zap.Error(err), zap.String("method", "RegisterTransaction"))
		return &response.RegisterTransactionResponse{}, errors.New("Invalid request parameters")
	}

	// Check for context cancellation
	if ctx.Err() != nil {
		return nil, ctx.Err()
	}

	// Start a new transaction
	ctx, err := s.accountTransactionRepo.BeginTx(ctx)
	if err != nil {
		s.logger.Error("Failed to start a new transaction")
		return nil, err
	}

	transaction := &entity.AccountTransaction{
		TransactionGroupID: req.TransactionGroupID,
		FinancialAccountID: req.FinancialAccountID,
		Amount:             req.Amount,
		Description:        req.Description,
		Status:             enum.Peniding, // Assuming enum.Pending exists
	}

	// Insert the transaction into the database
	err = s.accountTransactionRepo.Insert(ctx, transaction)
	if err != nil {
		s.logger.Error("Failed to insert transaction into the database")
		s.accountTransactionRepo.RollbackTx(ctx)
		return nil, err
	}

	// Commit the transaction
	err = s.accountTransactionRepo.CommitTx(ctx)
	if err != nil {
		s.logger.Error("Failed to commit transaction")
		return nil, err
	}

	s.logger.Info("Transaction successfully registered")

	return &response.RegisterTransactionResponse{
		TransactionID: transaction.TransactionID,
		Status:        transaction.Status,
		CreatedAt:     transaction.CreatedAt,
	}, nil
}

// DeleteTransaction deletes an account transaction
func (s *Service) DeleteTransaction(ctx context.Context, transactionID int64) error {
	s.logger.Info("Starting the process to delete transaction", zap.Int64("TransactionID", transactionID))

	// Start a new transaction
	txContext, err := s.accountTransactionRepo.BeginTx(ctx)
	if err != nil {
		s.logger.Error("Failed to start a new transaction", zap.Error(err))
		return err
	}

	// Delete the transaction from the database
	if err := s.accountTransactionRepo.Delete(txContext, transactionID); err != nil {
		s.logger.Error("Failed to delete transaction", zap.Error(err))
		_ = s.accountTransactionRepo.RollbackTx(txContext)
		return err
	}

	// Commit the transaction
	if err := s.accountTransactionRepo.CommitTx(txContext); err != nil {
		s.logger.Error("Failed to commit transaction", zap.Error(err))
		return err
	}

	s.logger.Info("Transaction successfully deleted")
	return nil

}

// GetTransactionByID retrieves an account transaction by its ID
func (s *Service) GetTransactionByID(ctx context.Context, transactionID int64) (*entity.AccountTransaction, error) {
	// Input validation
	if transactionID <= 0 {
		s.logger.Error("Invalid TransactionID provided", zap.Int64("TransactionID", transactionID))
		return nil, errors.New("Invalid TransactionID")
	}

	s.logger.Info("Starting the process to retrieve transaction", zap.Int64("TransactionID", transactionID))

	transaction, err := s.accountTransactionRepo.GetByID(ctx, transactionID)
	if err != nil {
		s.logger.Error("Failed to retrieve transaction", zap.Int64("TransactionID", transactionID), zap.Error(err))
		return nil, err
	}

	if transaction == nil {
		s.logger.Warn("Transaction not found", zap.Int64("TransactionID", transactionID))
		return nil, errors.New("Transaction not found")
	}

	s.logger.Info("Transaction successfully retrieved", zap.Int64("TransactionID", transactionID))

	return transaction, nil
}

// ListTransactionsByAccountID lists all transactions for a given account ID
func (s *Service) ListTransactionsByAccountID(ctx context.Context, accountID int) ([]*entity.AccountTransaction, error) {
	// Validate input
	if accountID <= 0 {
		s.logger.Error("Invalid account ID specified", zap.Int("AccountID", accountID))
		return nil, errors.New("Invalid account ID")
	}

	s.logger.Info("Starting the process to list transactions by account ID", zap.Int("AccountID", accountID))

	// Retrieve transactions
	transactions, err := s.accountTransactionRepo.ListByAccountID(ctx, accountID)
	if err != nil {
		s.logger.Error("Failed to list transactions by account ID", zap.Int("AccountID", accountID), zap.Error(err))
		return nil, fmt.Errorf("failed to list transactions: %w", err)
	}

	// Check for empty list
	if len(transactions) == 0 {
		s.logger.Warn("No transactions found for the given account ID", zap.Int("AccountID", accountID))
		return nil, errors.New("No transactions found")
	}

	s.logger.Info("Successfully listed transactions by account ID", zap.Int("AccountID", accountID), zap.Int("TransactionCount", len(transactions)))

	return transactions, nil
}

// ListTransactionsByGroupID lists all transactions for a given group ID
func (s *Service) ListTransactionsByGroupID(ctx context.Context, groupID int) ([]*entity.AccountTransaction, error) {
	// Validate input
	if groupID <= 0 {
		s.logger.Error("Invalid group ID specified", zap.Int("GroupID", groupID))
		return nil, errors.New("Invalid group ID")
	}

	s.logger.Info("Starting the process to list transactions by group ID", zap.Int("GroupID", groupID))

	// Retrieve transactions
	transactions, err := s.accountTransactionRepo.ListByTransactionGroupID(ctx, groupID)
	if err != nil {
		s.logger.Error("Failed to list transactions by group ID", zap.Int("GroupID", groupID), zap.Error(err))
		return nil, fmt.Errorf("failed to list transactions: %w", err)
	}

	// Check for empty list
	if len(transactions) == 0 {
		s.logger.Warn("No transactions found for the given group ID", zap.Int("GroupID", groupID))
		return nil, errors.New("No transactions found")
	}

	s.logger.Info("Successfully listed transactions by group ID", zap.Int("GroupID", groupID), zap.Int("TransactionCount", len(transactions)))

	return transactions, nil

}

func (s *Service) CancelTransaction(ctx context.Context, transactionID int64) error {
	s.logger.Info("Starting to cancel transaction", zap.Int64("TransactionID", transactionID))

	// Validate the transaction ID
	if transactionID <= 0 {
		s.logger.Error("Invalid transaction ID specified", zap.Int64("TransactionID", transactionID))
		return errors.New("invalid transaction ID")
	}

	// Begin database transaction
	ctx, err := s.accountTransactionRepo.BeginTx(ctx)
	if err != nil {
		s.logger.Error("Failed to begin database transaction", zap.Error(err))
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	// Fetch the existing transaction
	existingTransaction, err := s.accountTransactionRepo.GetByID(ctx, transactionID)
	if err != nil {
		s.logger.Error("Failed to fetch existing transaction", zap.Int64("TransactionID", transactionID), zap.Error(err))
		s.accountTransactionRepo.RollbackTx(ctx)
		return fmt.Errorf("failed to fetch existing transaction: %w", err)
	}

	// Validate if the transaction is already cancelled or completed
	if existingTransaction.Status == enum.Cancelled || existingTransaction.Status == enum.Completed {
		s.logger.Error("Cannot cancel a transaction that is already cancelled or completed", zap.Int64("TransactionID", transactionID))
		s.accountTransactionRepo.RollbackTx(ctx)
		return errors.New("invalid transaction state for cancellation")
	}

	// Your business logic to cancel the transaction, for example:
	existingTransaction.Status = enum.Cancelled

	// Update the transaction status
	if err := s.accountTransactionRepo.Update(ctx, existingTransaction); err != nil {
		s.logger.Error("Failed to update the transaction status to 'cancelled'", zap.Int64("TransactionID", transactionID), zap.Error(err))
		s.accountTransactionRepo.RollbackTx(ctx)
		return fmt.Errorf("failed to update transaction: %w", err)
	}

	// Commit database transaction
	if err := s.accountTransactionRepo.CommitTx(ctx); err != nil {
		s.logger.Error("Failed to commit transaction", zap.Error(err))
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	s.logger.Info("Successfully cancelled the transaction", zap.Int64("TransactionID", transactionID))
	return nil
}

// Transfer initiates a transfer between accounts
func (s *Service) Transfer(ctx context.Context, req request.TransferRequest) (res response.TransferResponse, err error) {
	s.logger.Info("Starting the process to transfer funds",
		zap.Int("SenderAccountID", req.SenderAccountID),
		zap.Int("ReceiverAccountID", req.ReceiverAccountID),
		zap.Float64("Amount", req.Amount))

	// Validate request
	if err := req.Validate(); err != nil {
		s.logger.Error("Invalid transfer request", zap.Error(err))
		return res, err
	}

	// Start database transaction
	ctx, err = s.accountTransactionRepo.BeginTx(ctx)
	if err != nil {
		s.logger.Error("Failed to begin database transaction", zap.Error(err))
		return res, err
	}

	defer func() {
		if err != nil {
			s.accountTransactionRepo.RollbackTx(ctx)
		}
	}()

	senderCurrentBalance, err := s.accountTransactionRepo.GetCurrentBalance(ctx, req.SenderAccountID)
	if err != nil {
		s.logger.Error("Failed to fetch sender's current balance", zap.Error(err))
		return res, err
	}

	receiverCurrentBalance, err := s.accountTransactionRepo.GetCurrentBalance(ctx, req.ReceiverAccountID)
	if err != nil {
		s.logger.Error("Failed to fetch receiver's current balance", zap.Error(err))
		return res, err
	}

	// Validate that the sender has enough funds and the receiver can receive funds
	if senderCurrentBalance < req.Amount {
		s.logger.Error("Insufficient funds in the sender's account")
		err = errors.New("insufficient funds in the sender's account")
		return res, err
	}

	// Check if the sender's account is in a state that allows transactions (e.g., not frozen or closed)
	senderStatus, err := s.financialAccountService.GetAccountStatus(ctx, req.SenderAccountID)
	if err != nil {
		s.logger.Error("Failed to fetch the sender's account status", zap.Error(err))
		return res, err
	}

	if senderStatus != enum.Verified {
		s.logger.Error("Sender's account is not in a state that allows transactions", zap.String("status", senderStatus.String())) // Assuming you implement a String() method on the enum
		err = errors.New("sender's account is not in a state that allows transactions")
		return res, err
	}

	// Check if the receiver's account is in a state that allows receiving transactions (e.g., not frozen or closed)
	receiverStatus, err := s.financialAccountService.GetAccountStatus(ctx, req.ReceiverAccountID)
	if err != nil {
		s.logger.Error("Failed to fetch the receiver's account status", zap.Error(err))
		return res, err
	}

	if receiverStatus != enum.Verified {
		s.logger.Error("Receiver's account is not in a state that allows receiving transactions", zap.String("status", receiverStatus.String()))
		err = errors.New("receiver's account is not in a state that allows receiving transactions")
		return res, err
	}

	// Generate new transaction group ID
	txGroupID, err := s.accountTransactionRepo.CreateNewTransactionGroup(ctx)
	if err != nil {
		s.logger.Error("Failed to create new transaction group", zap.Error(err))
		return res, err
	}

	// Log the generated txGroupID
	s.logger.Info("Generated transaction group ID", zap.Int("TransactionGroupID", txGroupID))

	// please check docs/doc.go for more scenraios in populating the status:
	// /home/delaram/go/src/github.com/delaram-gholampoor-sagha/Digital-Wallet/docs/doc.go
	// it starts from line 102 till 571
	// // Perform a risk check on the transaction
	// isRisky, err := s.riskManagementService.PerformRiskCheck(ctx, req)
	// if err != nil {
	// 	s.logger.Error("Failed to perform risk check", zap.Error(err))
	// 	return res, err
	// }

	// // Decide the transaction status based on risk check
	// var transactionStatus enum.StatusType // Replace with your actual enum type
	// if isRisky {
	// 	transactionStatus = enum.Pending // Pending because the transaction is risky and might need manual review
	// } else {
	// 	transactionStatus = enum.Completed // Not risky, so mark it as completed
	// }

	// Debit the sender's account
	if err := s.accountTransactionRepo.DebitAccount(ctx, req.SenderAccountID, req.Amount); err != nil {
		s.logger.Error("Failed to debit the sender's account", zap.Error(err))

		return res, err
	}

	// Credit the receiver's account
	if err := s.accountTransactionRepo.CreditAccount(ctx, req.ReceiverAccountID, req.Amount); err != nil {
		s.logger.Error("Failed to credit the receiver's account", zap.Error(err))
		return res, err
	}

	// Create sender's transaction record
	senderTx := &entity.AccountTransaction{
		TransactionGroupID: txGroupID,
		FinancialAccountID: req.SenderAccountID,
		Amount:             -req.Amount,                       // Negative because it's a debit
		Balance:            senderCurrentBalance - req.Amount, // Update the balance after the transaction
		Description:        &req.Description,
		Status:             enum.Peniding, // Or enum.Completed, depending on your logic
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		DeletedAt:          nil, // Not deleted, so it's nil
	}

	if err := s.accountTransactionRepo.Insert(ctx, senderTx); err != nil {
		s.logger.Error("Failed to insert sender's transaction record", zap.Error(err))
		return res, err
	}

	// Create receiver's transaction record
	receiverTx := &entity.AccountTransaction{
		TransactionGroupID: txGroupID,
		FinancialAccountID: req.ReceiverAccountID,
		Amount:             req.Amount,
		Balance:            receiverCurrentBalance + req.Amount, // Update the balance after the transaction
		Description:        &req.Description,                    // If the receiver has a different description, update this
		Status:             enum.Peniding,                       // Or enum.Completed, depending my your logic
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
		DeletedAt:          nil, // Not deleted, so it's nil
	}

	if err := s.accountTransactionRepo.Insert(ctx, receiverTx); err != nil {
		s.logger.Error("Failed to insert receiver's transaction record", zap.Error(err))
		return res, err
	}

	// Commit transaction
	if err := s.accountTransactionRepo.CommitTx(ctx); err != nil {
		s.logger.Error("Failed to commit transaction", zap.Error(err))
		return res, err
	}

	res = response.TransferResponse{
		SenderTx:   *senderTx,
		ReceiverTx: *receiverTx,
	}

	s.logger.Info("Successfully transferred funds")

	return res, nil
}

func (s *Service) GetAccountTransactionHistory(ctx context.Context, accountID int) ([]*entity.AccountTransaction, error) {
	// Validation (replace this with your own validation logic if needed)
	if accountID <= 0 {
		s.logger.Error("Invalid account ID", zap.Int("accountID", accountID), zap.String("method", "GetAccountTransactionHistory"))
		return nil, derror.NewValidationError("Invalid account ID")
	}

	// TODO: Optional: I need to also add a check here to make sure the accountID belongs to the authenticated user.
	// ... (My authorization logic here)

	transactions, err := s.accountTransactionRepo.ListByAccountID(ctx, accountID)
	if err != nil {
		s.logger.Error("Failed to get transactions by account ID", zap.Error(err), zap.String("method", "GetAccountTransactionHistory"))
		return nil, derror.NewInternalSystemError()
	}

	if len(transactions) == 0 {
		return nil, derror.NewNotFoundError("No transactions found for the given account ID")
	}

	return transactions, nil
}
