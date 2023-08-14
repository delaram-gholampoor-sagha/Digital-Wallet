package entity

import (
	"time"

	"github.com/delaram-gholampoor-sagha/Digital-Wallet/internal/entity/enum"
)

type User struct {
	ID                 int
	Username           string
	Password           string
	FirstName          string
	LastName           string
	Email              string
	ValidatedEmail     bool
	Cellphone          string
	ValidatedCellphone bool
	Admin              bool
	Status             enum.UserStatus
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          *time.Time
}
