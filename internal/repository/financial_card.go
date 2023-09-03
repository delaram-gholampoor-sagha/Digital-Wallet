package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type FinancialCard struct {
	cli *sql.DB
}

func (repo *FinancialCard) Insert(ctx context.Context, card *entity.FinancialCard) error {
	query := `
        INSERT INTO financial_card 
        (account_id, card_type, card_number, expiration_date, card_holder_name, cvv, status, issued_date, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
        RETURNING card_id
    `
	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.FinancialCard.Insert.PrepareContext: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRowContext(ctx,
		card.AccountID,
		card.CardType,
		card.CardNumber,
		card.ExpirationDate,
		card.CardHolderName,
		card.CVV,
		card.Status,
		card.IssuedDate,
	).Scan(&card.CardID)
	if err != nil {
		return fmt.Errorf("repository.FinancialCard.Insert.QueryRowContext: %w", err)
	}

	return nil
}

func (repo *FinancialCard) Update(ctx context.Context, card *entity.FinancialCard) error {
	query := `
        UPDATE financial_card
        SET account_id = $1, card_type = $2, card_number = $3, expiration_date = $4, card_holder_name = $5, cvv = $6, status = $7, issued_date = $8, updated_at = CURRENT_TIMESTAMP
        WHERE card_id = $9
    `
	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.FinancialCard.Update.PrepareContext: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx,
		card.AccountID,
		card.CardType,
		card.CardNumber,
		card.ExpirationDate,
		card.CardHolderName,
		card.CVV,
		card.Status,
		card.IssuedDate,
		card.CardID,
	)
	if err != nil {
		return fmt.Errorf("repository.FinancialCard.Update.ExecContext: %w", err)
	}

	return nil
}

func (repo *FinancialCard) Delete(ctx context.Context, cardID int64) error {
	query := "DELETE FROM financial_card WHERE card_id = $1"
	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.FinancialCard.Delete.PrepareContext: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, cardID)
	if err != nil {
		return fmt.Errorf("repository.FinancialCard.Delete.ExecContext: %w", err)
	}

	return nil
}

func (repo *FinancialCard) GetByID(ctx context.Context, cardID int64) (*entity.FinancialCard, error) {
	query := `
		SELECT card_id, account_id, card_type, card_number, expiration_date, card_holder_name, cvv, status, issued_date, created_at, updated_at, deleted_at
		FROM financial_card
		WHERE card_id = $1
	`

	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query for getting a FinancialCard by ID: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, cardID)

	var card entity.FinancialCard
	var deletedAt sql.NullTime // Handle NULL time

	if err := row.Scan(
		&card.CardID,
		&card.AccountID,
		&card.CardNumber,
		&card.CardType,
		&card.ExpirationDate,
		&card.CardHolderName,
		&card.CVV,
		&card.Status,
		&card.IssuedDate,
		&card.CreatedAt,
		&card.UpdatedAt,
		&deletedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("FinancialCard not found")
		}
		return nil, fmt.Errorf("failed to scan FinancialCard data: %w", err)
	}

	if deletedAt.Valid {
		card.DeletedAt = &deletedAt.Time
	}

	return &card, nil
}

func (repo *FinancialCard) ListByAccountID(ctx context.Context, accountID int) ([]*entity.FinancialCard, error) {
	query := `
	SELECT card_id, account_id, card_type, card_number, expiration_date, status, created_at, updated_at
	FROM financial_card
	WHERE account_id = $1
`

	rows, err := repo.cli.QueryContext(ctx, query, accountID)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialCard.ListByAccountID.QueryContext: %w", err)
	}
	defer rows.Close()

	var cards []*entity.FinancialCard
	for rows.Next() {
		var card entity.FinancialCard
		if err := rows.Scan(
			&card.CardID,
			&card.AccountID,
			&card.CardNumber,
			&card.CardType,
			&card.ExpirationDate,
			&card.CardHolderName,
			&card.CVV,
			&card.Status,
			&card.IssuedDate,
			&card.CreatedAt,
			&card.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("repository.FinancialCard.ListByAccountID.Scan: %w", err)
		}
		cards = append(cards, &card)
	}

	return cards, nil
}

func (repo *FinancialCard) ListByCardType(ctx context.Context, cardType enum.FinancialCardType) ([]*entity.FinancialCard, error) {
	query := `
	SELECT card_id, account_id, card_type, card_number, expiration_date, status, created_at, updated_at
	FROM financial_card
	WHERE card_type = $1
`

	rows, err := repo.cli.QueryContext(ctx, query, cardType)
	if err != nil {
		return nil, fmt.Errorf("repository.FinancialCard.ListByCardType.QueryContext: %w", err)
	}
	defer rows.Close()

	var cards []*entity.FinancialCard
	for rows.Next() {
		var card entity.FinancialCard
		if err := rows.Scan(
			&card.CardID,
			&card.AccountID,
			&card.CardNumber,
			&card.CardType,
			&card.ExpirationDate,
			&card.CardHolderName,
			&card.CVV,
			&card.Status,
			&card.IssuedDate,
			&card.CreatedAt,
			&card.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("repository.FinancialCard.ListByCardType.Scan: %w", err)
		}
		cards = append(cards, &card)
	}

	return cards, nil
}
