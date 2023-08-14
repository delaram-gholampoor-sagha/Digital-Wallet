package enum

type UserStatus uint

const (
	UserInactive UserStatus = iota
	UserActive
	UserBanned
)
