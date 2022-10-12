package middleware

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(hashedPassword), err
}

func CheckPasswordHash(password, hash string) bool {
	sameHashAndPassword := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return sameHashAndPassword == nil
}
