package crypto

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

type crypto struct {
}

func New() Crypto {
	return crypto{}
}

func (c crypto) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 1)
	return string(bytes), err
}

func (c crypto) CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (c crypto) GenerateOTP(length int) (string, error) {
	otpChars := "0123456789"
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

func (c crypto) GenerateToken(length int) (string, error) {
	tokenChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	tokenCharsLength := len(tokenChars)
	for i := 0; i < length; i++ {
		buffer[i] = tokenChars[int(buffer[i])%tokenCharsLength]
	}

	return string(buffer), nil
}
