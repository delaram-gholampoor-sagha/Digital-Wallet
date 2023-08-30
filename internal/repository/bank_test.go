package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity"
	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBank_GetByID(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		wantErr bool
		want    entity.Bank
	}{
		{
			name: "successful case",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "code", "status"}).
					AddRow(1, "Test Bank", "TST", 1)
				mock.ExpectQuery("^SELECT (.+) FROM banks WHERE id = \\$1").
					WithArgs(1).
					WillReturnRows(rows)
			},
			wantErr: false,
			want: entity.Bank{
				BankID:   1,
				Name:     "Test Bank",
				BankCode: "TST",
				Status:   1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.mock(mock)

			repo := &Bank{cli: db}
			bank, err := repo.GetByID(context.TODO(), 1)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, *bank)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

func TestBank_GetByCode(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		wantErr bool
		want    entity.Bank
	}{
		{
			name: "successful case",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "code", "status"}).
					AddRow(1, "Test Bank", "TST", 1)
				mock.ExpectQuery("^SELECT (.+) FROM banks WHERE code = \\$1").
					WithArgs("TST").
					WillReturnRows(rows)
			},
			wantErr: false,
			want: entity.Bank{
				BankID:   1,
				Name:     "Test Bank",
				BankCode: "TST",
				Status:   1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.mock(mock)

			repo := &Bank{cli: db}
			bank, err := repo.GetByCode(context.TODO(), "TST")
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, *bank)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

func TestBank_GetByName(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		wantErr bool
		want    entity.Bank
	}{
		{
			name: "successful case",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "code", "status"}).
					AddRow(1, "Test Bank", "TST", 1)
				mock.ExpectQuery("^SELECT (.+) FROM banks WHERE name = \\$1").
					WithArgs("Test Bank").
					WillReturnRows(rows)
			},
			wantErr: false,
			want: entity.Bank{
				BankID:   1,
				Name:     "Test Bank",
				BankCode: "TST",
				Status:   1,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.mock(mock)

			repo := &Bank{cli: db}
			bank, err := repo.GetByName(context.TODO(), "Meli")
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, *bank)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

func TestBank_Insert(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "successful case",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("^INSERT INTO banks").
					WithArgs("Test Bank", "TST", 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.mock(mock)

			repo := &Bank{cli: db}
			bank := &entity.Bank{
				Name:     "Test Bank",
				BankCode: "TST",
				Status:   1,
			}
			err = repo.Insert(context.TODO(), bank)
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

func TestBank_Update(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "successful case",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("^UPDATE banks SET").
					WithArgs("Updated Bank", "UPD", 1, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.mock(mock)

			repo := &Bank{cli: db}
			bank := &entity.Bank{
				BankID:   1,
				Name:     "Updated Bank",
				BankCode: "UPD",
				Status:   1,
			}
			err = repo.Update(context.TODO(), bank)
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

func TestBank_Delete(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "successful case",
			mock: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec("^DELETE FROM banks WHERE id = \\$1").
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.mock(mock)

			repo := &Bank{cli: db}
			err = repo.Delete(context.TODO(), 1)
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

func TestBank_ListAll(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		wantErr bool
		want    []*entity.Bank
	}{
		{
			name: "successful case",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "code", "status"}).
					AddRow(1, "Bank One", "B1", 1).
					AddRow(2, "Bank Two", "B2", 1)
				mock.ExpectQuery("^SELECT (.+) FROM banks$").
					WillReturnRows(rows)
			},
			wantErr: false,
			want: []*entity.Bank{
				{BankID: 1, Name: "Bank One", BankCode: "B1", Status: 1},
				{BankID: 2, Name: "Bank Two", BankCode: "B2", Status: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.mock(mock)

			repo := &Bank{cli: db}
			banks, err := repo.ListAll(context.TODO())
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, banks)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

// You'll need to import or define your enum.BankStatus

func TestBank_ListByStatus(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		wantErr bool
		want    []*entity.Bank
	}{
		{
			name: "successful case",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "code", "status"}).
					AddRow(1, "Bank One", "B1", 1)
				mock.ExpectQuery("^SELECT (.+) FROM banks WHERE status = \\$1").
					WithArgs(1).
					WillReturnRows(rows)
			},
			wantErr: false,
			want: []*entity.Bank{
				{BankID: 1, Name: "Bank One", BankCode: "B1", Status: 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.mock(mock)

			repo := &Bank{cli: db}
			banks, err := repo.ListByStatus(context.TODO(), enum.BankActive) // Replace `enum.Active` with the actual status enum you're using.
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, banks)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

func TestBank_IsBankCodeExist(t *testing.T) {
	tests := []struct {
		name    string
		mock    func(mock sqlmock.Sqlmock)
		wantErr bool
		want    bool
	}{
		{
			name: "successful case",
			mock: func(mock sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"count"}).
					AddRow(1)
				mock.ExpectQuery("^SELECT COUNT(.+) FROM banks WHERE code = \\$1").
					WithArgs("B1").
					WillReturnRows(rows)
			},
			wantErr: false,
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			defer db.Close()

			tt.mock(mock)

			repo := &Bank{cli: db}
			exist, err := repo.IsBankCodeExist(context.TODO(), "B1")
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.want, exist)
			}

			err = mock.ExpectationsWereMet()
			require.NoError(t, err)
		})
	}
}

type mockBankDatabase struct {
	db *sql.DB
}

func (m *mockBankDatabase) Close() error {
	return m.db.Close()
}

func (m *mockBankDatabase) DB() *sql.DB {
	return m.db
}
func TestNewBank(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error creating sqlmock: %s", err)
	}
	defer db.Close()

	mockDb := &mockBankDatabase{
		db: db,
	}

	bankRepo := NewBank(mockDb)

	assert.IsType(t, &Bank{}, bankRepo)
	assert.Equal(t, db, bankRepo.cli)
}
