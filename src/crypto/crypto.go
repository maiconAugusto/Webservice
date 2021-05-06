package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

func CreateHash(hash string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(hash), bcrypt.DefaultCost)
}
func VerifyHash(dbHash string, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(dbHash), []byte(hash))
}
