package core

import "golang.org/x/crypto/bcrypt"

func HashGenerator(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CheckPassword(passwordHash, passwordPlain string) (err error) {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordPlain))
}
