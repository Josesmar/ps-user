package security

import "golang.org/x/crypto/bcrypt"

// Hash takes a string and puts a hash on it
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// VerifyPassword compares a password and hash and returns if they match
func VerifyPassword(passwordWithHash string, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordWithHash), []byte(passwordString))
}
