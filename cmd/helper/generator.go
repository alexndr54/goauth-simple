package helper

import (
	"crypto/rand"
	"errors"
	"github.com/google/uuid"
	"strings"
)

func GenerateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	for i, v := range b {
		b[i] = charset[v%byte(len(charset))]
	}
	return string(b), nil
}

func GenerateApiKey() (error, string) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return errors.New("Ada kesalahan ketika membuat UUID"), ""
	}

	add, err := GenerateRandomString(10)
	if err != nil {
		return errors.New("Ada kesalahan ketika membuat opsional apikey"), ""
	}

	apikey := strings.Replace(newUUID.String(), "-", "", -1) + add
	return nil, apikey
}

func GenerateApiUsername() (string, error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return "", errors.New("Ada kesalahan ketika membuat UUID")
	}
	return strings.Replace(newUUID.String(), "-", "", -1), nil
}
