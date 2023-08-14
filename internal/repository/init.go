package repository

import "github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/protocol"

func NewUser(database protocol.Database) *User {
	return &User{cli: database.DB()}
}


