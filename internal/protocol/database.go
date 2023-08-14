package protocol

import "database/sql"

type Database interface {
	Close() error
	DB() *sql.DB
}
