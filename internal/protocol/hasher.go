package protocol

// Hasher provides an interface to hash passwords, abstracting the bcrypt package.
type Hasher interface {
	GenerateFromPassword(password []byte, cost int) ([]byte, error)
	CompareHashAndPassword(hashedPassword, password []byte) error
}
