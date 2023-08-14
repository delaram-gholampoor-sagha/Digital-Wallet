package postgres

import (
	"database/sql"
	"fmt"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/config"
	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func New(cfg config.Postgres) (*Repository, error) {
	db, err := sql.Open("postgres", dsn(cfg))
	if err != nil {
		return nil, fmt.Errorf("open postgres error: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping postgres error: %w", err)
	}

	return &Repository{
		db: db,
	}, nil
}

func (repo *Repository) Close() error {
	if err := repo.db.Close(); err != nil {
		return fmt.Errorf("close postgres connections error: %w", err)
	}
	return nil
}

func (repo *Repository) DB() *sql.DB {
	return repo.db
}

func dsn(cfg config.Postgres) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
}
