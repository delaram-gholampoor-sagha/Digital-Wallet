package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/pkg/derror"
)

type User struct {
	cli *sql.DB
}

func (repo *User) Get(ctx context.Context, id int) (entity.User, error) {
	query := `
	SELECT
		id, username, password, first_name,
		last_name, email, validated_email, cellphone, validated_cellphone,
		is_admin, status , created_at, updated_at, deleted_at
	FROM
		users
	WHERE
		id = $1 AND deleted_at IS NULL
   `

	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return entity.User{}, fmt.Errorf("repository.User.Get.PrepareContext: %w", err)
	}

	row := stmt.QueryRowContext(ctx, id)

	var user entity.User
	if err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.ValidatedEmail,
		&user.Cellphone,
		&user.ValidatedCellphone,
		&user.Admin,
		&user.Status,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, derror.NewNotFoundError("User not found")
		}
		return entity.User{}, fmt.Errorf("repository.User.Get.Scan: %w", err)
	}

	return user, nil
}

func (repo *User) GetByUsername(ctx context.Context, username string) (entity.User, error) {
	query := `
		SELECT
			id, username, password, first_name,
			last_name, email, validated_email, cellphone, validated_cellphone,
			is_admin, status , created_at, updated_at, deleted_at
		FROM
			users
		WHERE
			username = $1 AND deleted_at IS NULL
	`

	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return entity.User{}, fmt.Errorf("repository.User.GetByUsername.PrepareContext: %w", err)
	}

	row := stmt.QueryRowContext(ctx, username)

	var user entity.User
	if err := row.Scan(&user.ID, &user.Username, &user.Password,
		&user.FirstName, &user.LastName, &user.Email, &user.ValidatedEmail, &user.Cellphone, &user.ValidatedCellphone,
		&user.Admin, &user.Status, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return entity.User{}, derror.NewNotFoundError("User not found")
		}
		return entity.User{}, fmt.Errorf("repository.User.GetByUsername.Scan: %w", err)
	}

	return user, nil
}

func (repo *User) Insert(ctx context.Context, user entity.User) error {
	query := `
		INSERT INTO
			users
			(
				username, password, first_name,
				last_name, email, validated_email, cellphone, validated_cellphone,
				is_admin, status , created_at, updated_at, deleted_at
			)
		VALUES
			(
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
				NOW(), NOW(), NULL
			)
	`

	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.User.Insert.PrepareContext: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Username, user.Password,
		user.FirstName, user.LastName, user.Email, user.ValidatedEmail, user.Cellphone, user.ValidatedCellphone,
		user.Admin, user.Status)
	if err != nil {
		return fmt.Errorf("repository.User.Insert.ExecContext: %w", err)
	}

	return nil
}

func (repo *User) Update(ctx context.Context, user entity.User) error {
	query := `
	UPDATE 
		users
	SET
		password = $1,
		first_name = $2,
		last_name = $3,
		email = $4,
		validated_email = $5, 
		cellphone = $6,
		validated_cellphone = $7,
		status = $8,
		updated_at = NOW()
	WHERE
		id = $9
	`

	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("repository.User.Update.PrepareContext: %w", err)
	}

	res, err := stmt.ExecContext(ctx, user.Password, user.FirstName, user.LastName,
		user.Email, user.ValidatedEmail, user.Cellphone, user.ValidatedCellphone, user.ID)
	if err != nil {
		return fmt.Errorf("repository.User.Update.PrepareContext: %w", err)
	}

	ra, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("repository.User.Update.PrepareContext: %w", err)
	}

	if ra == 0 {
		return derror.NewNotFoundError("User not found")
	}

	return nil
}

func (repo *User) IsUsernameExist(ctx context.Context, username string) (bool, error) {
	query := `
	SELECT
		COUNT(*)
	FROM
		users
	WHERE
		username = $1 AND deleted_at IS NULL
`
	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return false, fmt.Errorf("repository.User.IsUsernameExist.PrepareContext: %w", err)
	}

	row := stmt.QueryRowContext(ctx, username)

	var count uint
	if err := row.Scan(&count); err != nil {
		return false, fmt.Errorf("repository.User.IsUsernameExist.Scan: %w", err)
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}

func (repo *User) IsExist(ctx context.Context, id int) (bool, error) {
	query := `
	SELECT
		COUNT(*)
	FROM
		users
	WHERE
		id = $1 AND deleted_at IS NULL
`
	stmt, err := repo.cli.PrepareContext(ctx, query)
	if err != nil {
		return false, fmt.Errorf("repository.User.IsExist.PrepareContext: %w", err)
	}

	row := stmt.QueryRowContext(ctx, id)

	var count uint
	if err := row.Scan(&count); err != nil {
		return false, fmt.Errorf("repository.User.is_exist.scan: %w", err)
	}

	if count == 0 {
		return false, nil
	}

	return true, nil
}
