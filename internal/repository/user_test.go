package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUser_Get(t *testing.T) {
	// Test cases
	tests := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "successful case",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"id", "username", "password", "first_name",
					"last_name", "email", "validated_email", "cellphone", "validated_cellphone",
					"is_admin", "status", "created_at", "updated_at", "deleted_at",
				}).AddRow(1, "testuser", "testpass", "test", "user", "test@user.com", true, "1234567890", true, false, 1, time.Now(), time.Now(), nil)
				mock.ExpectPrepare("^SELECT (.+) FROM users WHERE id = \\$1 AND deleted_at IS NULL").ExpectQuery().WithArgs(1).WillReturnRows(rows)
			},
			wantErr: false,
		},
		{
			name: "user not found",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("^SELECT (.+) FROM users WHERE id = \\$1 AND deleted_at IS NULL").ExpectQuery().WithArgs(1).WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
		},
		{
			name: "prepare context error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("^SELECT (.+) FROM users WHERE id = \\$1 AND deleted_at IS NULL").WillReturnError(sql.ErrConnDone) // Simulating a connection error for demonstration
			},
			wantErr: true,
		},
		{
			name: "scan error",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"id", "username", // Intentionally leaving out other columns to cause a scan error
				}).AddRow(1, "testuser")
				mock.ExpectPrepare("^SELECT (.+) FROM users WHERE id = \\$1 AND deleted_at IS NULL").ExpectQuery().WithArgs(1).WillReturnRows(rows)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.mock(mock)

			repo := &User{cli: db}
			_, err = repo.Get(context.TODO(), 1)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

// go test github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/repository -coverprofile=coverage.out
// go tool cover -func=coverage.out
// go tool cover -html=coverage.out

func TestUser_GetByUsername(t *testing.T) {
	// Test cases
	tests := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		wantErr bool
		want    entity.User
	}{
		{
			name: "successful case",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"id", "username", "password", "first_name",
					"last_name", "email", "validated_email", "cellphone", "validated_cellphone",
					"is_admin", "status", "created_at", "updated_at", "deleted_at",
				}).AddRow(1, "testuser", "testpass", "test", "user", "test@user.com", true, "1234567890", true, false, 1, time.Now(), time.Now(), nil)
				mock.ExpectPrepare("^SELECT (.+) FROM users WHERE username = \\$1 AND deleted_at IS NULL").ExpectQuery().WithArgs("testuser").WillReturnRows(rows)
			},
			wantErr: false,
			want: entity.User{
				ID:                 1,
				Username:           "testuser",
				Password:           "testpass",
				FirstName:          "test",
				LastName:           "user",
				Email:              "test@user.com",
				ValidatedEmail:     true,
				Cellphone:          "1234567890",
				ValidatedCellphone: true,
				Admin:              false,
				Status:             1,
			},
		},
		{
			name: "user not found",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("^SELECT (.+) FROM users WHERE username = \\$1 AND deleted_at IS NULL").ExpectQuery().WithArgs("testuser").WillReturnError(sql.ErrNoRows)
			},
			wantErr: true,
			want:    entity.User{},
		},
		{
			name: "prepare context error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("^SELECT (.+) FROM users WHERE username = \\$1 AND deleted_at IS NULL").WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
			want:    entity.User{},
		},
		{
			name: "scan error",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{
					"id", "username", // Intentionally leaving out other columns to cause a scan error
				}).AddRow(1, "testuser")
				mock.ExpectPrepare("^SELECT (.+) FROM users WHERE username = \\$1 AND deleted_at IS NULL").ExpectQuery().WithArgs("testuser").WillReturnRows(rows)
			},
			wantErr: true,
			want:    entity.User{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.mock(mock)

			repo := &User{cli: db}
			user, err := repo.GetByUsername(context.TODO(), "testuser")
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				// require.Equal(t, tt.want, user)

				require.Equal(t, tt.want.ID, user.ID)
				require.Equal(t, tt.want.Username, user.Username)
				require.Equal(t, tt.want.Password, user.Password)
				require.Equal(t, tt.want.FirstName, user.FirstName)
				require.Equal(t, tt.want.LastName, user.LastName)
				require.Equal(t, tt.want.Email, user.Email)
				require.Equal(t, tt.want.ValidatedEmail, user.ValidatedEmail)
				require.Equal(t, tt.want.Cellphone, user.Cellphone)
				require.Equal(t, tt.want.ValidatedCellphone, user.ValidatedCellphone)
				require.Equal(t, tt.want.Admin, user.Admin)
				require.Equal(t, tt.want.Status, user.Status)

			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

func TestUser_Insert(t *testing.T) {
	// Test cases
	tests := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		wantErr bool
		input   entity.User
	}{
		{
			name: "successful insert",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare(`^INSERT INTO
			users
			\(
				username, password, first_name,
				last_name, email, validated_email, cellphone, validated_cellphone,
				is_admin, status , created_at, updated_at, deleted_at
			\)
		VALUES
			\(
				\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8, \$9, \$10,
				NOW\(\), NOW\(\), NULL
			\)`).ExpectExec().WithArgs(
					"testuser", "testpass", "test", "user", "test@user.com", true, "1234567890", true, false, 1,
				).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
			input: entity.User{
				Username:           "testuser",
				Password:           "testpass",
				FirstName:          "test",
				LastName:           "user",
				Email:              "test@user.com",
				ValidatedEmail:     true,
				Cellphone:          "1234567890",
				ValidatedCellphone: true,
				Admin:              false,
				Status:             1,
			},
		},
		{
			name: "prepare context error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("^INSERT INTO users").WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
			input: entity.User{
				Username: "testuser",
			},
		},
		{
			name: "exec context error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("^INSERT INTO users").ExpectExec().WithArgs(
					"testuser", "testpass", "test", "user", "test@user.com", true, "1234567890", true, false, 1,
				).WillReturnError(errors.New("execution error"))
			},
			wantErr: true,
			input: entity.User{
				Username:           "testuser",
				Password:           "testpass",
				FirstName:          "test",
				LastName:           "user",
				Email:              "test@user.com",
				ValidatedEmail:     true,
				Cellphone:          "1234567890",
				ValidatedCellphone: true,
				Admin:              false,
				Status:             1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.mock(mock)

			repo := &User{cli: db}
			err = repo.Insert(context.TODO(), tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

func TestUser_IsExist(t *testing.T) {
	tests := []struct {
		name     string
		mock     func(mock sqlmock.Sqlmock)
		wantErr  bool
		input    int
		expected bool
	}{
		{
			name: "user ID exists",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(1)
				mock.ExpectPrepare("^SELECT COUNT\\(\\*\\) FROM users WHERE").ExpectQuery().WithArgs(1).WillReturnRows(rows)
			},
			wantErr:  false,
			input:    1,
			expected: true,
		},
		{
			name: "user ID doesn't exist",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(0)
				mock.ExpectPrepare("^SELECT COUNT\\(\\*\\) FROM users WHERE").ExpectQuery().WithArgs(1).WillReturnRows(rows)
			},
			wantErr:  false,
			input:    1,
			expected: false,
		},
		{
			name: "prepare context error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("^SELECT COUNT\\(\\*\\) FROM users WHERE").WillReturnError(sql.ErrConnDone)
			},
			wantErr:  true,
			input:    1,
			expected: false,
		},
		{
			name: "query error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("^SELECT COUNT\\(\\*\\) FROM users WHERE").ExpectQuery().WithArgs(1).WillReturnError(sql.ErrNoRows)
			},
			wantErr:  true,
			input:    1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.mock(mock)

			repo := &User{cli: db}
			exists, err := repo.IsExist(context.TODO(), tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, exists)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

func TestUser_IsUsernameExist(t *testing.T) {
	tests := []struct {
		name     string
		mock     func(mock sqlmock.Sqlmock)
		wantErr  bool
		input    string
		expected bool
	}{
		{
			name: "username exists",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(1)
				mock.ExpectPrepare("^SELECT COUNT\\(\\*\\) FROM users WHERE").ExpectQuery().WithArgs("testuser").WillReturnRows(rows)
			},
			wantErr:  false,
			input:    "testuser",
			expected: true,
		},
		{
			name: "username doesn't exist",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"COUNT(*)"}).AddRow(0)
				mock.ExpectPrepare("^SELECT COUNT\\(\\*\\) FROM users WHERE").ExpectQuery().WithArgs("testuser").WillReturnRows(rows)
			},
			wantErr:  false,
			input:    "testuser",
			expected: false,
		},
		{
			name: "prepare context error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("^SELECT COUNT\\(\\*\\) FROM users WHERE").WillReturnError(sql.ErrConnDone)
			},
			wantErr:  true,
			input:    "testuser",
			expected: false,
		},
		{
			name: "query error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("^SELECT COUNT\\(\\*\\) FROM users WHERE").ExpectQuery().WithArgs("testuser").WillReturnError(sql.ErrNoRows)
			},
			wantErr:  true,
			input:    "testuser",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.mock(mock)

			repo := &User{cli: db}
			exists, err := repo.IsUsernameExist(context.TODO(), tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expected, exists)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

func TestUser_Update(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		wantErr bool
		input   entity.User
	}{
		{
			name: "successful update",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("^UPDATE users SET").ExpectExec().WithArgs(
					"testpass", "test", "user", "test@user.com", true, "1234567890", true, 1,
				).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantErr: false,
			input: entity.User{
				ID:                 1,
				Username:           "testuser",
				Password:           "testpass",
				FirstName:          "test",
				LastName:           "user",
				Email:              "test@user.com",
				ValidatedEmail:     true,
				Cellphone:          "1234567890",
				ValidatedCellphone: true,
				Status:             1,
			},
		},
		{
			name: "no rows affected",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("^UPDATE users SET").ExpectExec().WithArgs(
					"testpass", "test", "user", "test@user.com", true, "1234567890", true, 1,
				).WillReturnResult(sqlmock.NewResult(0, 0))
			},
			wantErr: true,
			input: entity.User{
				ID:                 1,
				Username:           "testuser",
				Password:           "testpass",
				FirstName:          "test",
				LastName:           "user",
				Email:              "test@user.com",
				ValidatedEmail:     true,
				Cellphone:          "1234567890",
				ValidatedCellphone: true,
				Status:             1,
			},
		},
		{
			name: "rows affected error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("^UPDATE users SET").ExpectExec().WithArgs(
					"testpass", "test", "user", "test@user.com", true, "1234567890", true, 1,
				).WillReturnResult(sqlmock.NewErrorResult(errors.New("rows affected error")))
			},
			wantErr: true,
			input: entity.User{
				ID:                 1,
				Username:           "testuser",
				Password:           "testpass",
				FirstName:          "test",
				LastName:           "user",
				Email:              "test@user.com",
				ValidatedEmail:     true,
				Cellphone:          "1234567890",
				ValidatedCellphone: true,
				Status:             1,
			},
		},
		{
			name: "prepare context error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("^UPDATE users SET").WillReturnError(errors.New("prepare context error"))
			},
			wantErr: true,
			input: entity.User{
				ID:                 1,
				Username:           "testuser",
				Password:           "testpass",
				FirstName:          "test",
				LastName:           "user",
				Email:              "test@user.com",
				ValidatedEmail:     true,
				Cellphone:          "1234567890",
				ValidatedCellphone: true,
				Status:             1,
			},
		},
		{
			name: "exec context error",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectPrepare("^UPDATE users SET").ExpectExec().WithArgs(
					"testpass", "test", "user", "test@user.com", true, "1234567890", true, 1,
				).WillReturnError(errors.New("exec context error"))
			},
			wantErr: true,
			input: entity.User{
				ID:                 1,
				Username:           "testuser",
				Password:           "testpass",
				FirstName:          "test",
				LastName:           "user",
				Email:              "test@user.com",
				ValidatedEmail:     true,
				Cellphone:          "1234567890",
				ValidatedCellphone: true,
				Status:             1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			tt.mock(mock)

			repo := &User{
				cli: db,
			}

			err = repo.Update(context.Background(), tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("User.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

type mockDatabase struct {
	db *sql.DB
}

func (m *mockDatabase) Close() error {
	return m.db.Close()
}

func (m *mockDatabase) DB() *sql.DB {
	return m.db
}
func TestNewUser(t *testing.T) {
	// Create an instance of the sqlmock database.
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock: %s", err)
	}
	defer db.Close()

	// Using the mockDatabase type to implement the Database interface.
	mockDb := &mockDatabase{
		db: db,
	}

	// Now, call your NewUser function.
	userRepo := NewUser(mockDb)

	// Assert that the resulting type is *User and that its cli is the same as the mock db.
	assert.IsType(t, &User{}, userRepo)
	assert.Equal(t, db, userRepo.cli)
}
