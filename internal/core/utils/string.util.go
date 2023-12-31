package utils

import (
	"net/mail"
	"os"
	"strconv"
	"strings"

	commonError "github.com/sijanstha/electronic-voting-system/internal/core/error"
	"golang.org/x/crypto/bcrypt"
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LoadResourceAsString(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", &commonError.ErrInternalServer{Message: "file not found on path: " + path}
	}

	return strings.TrimSpace(string(data)), nil
}

func ParseInteger(str string) int64 {
	data, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0
	}
	return data
}

func ParseBoolean(str string) *bool {
	b := false
	if str == "true" {
		b = true
	}

	return &b
}