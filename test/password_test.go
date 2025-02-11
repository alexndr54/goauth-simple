package test

import (
	"autentikasi1/cmd/helper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPassword(t *testing.T) {
	password := "arman123"
	var hashPassword string

	t.Run("Enkripsi", func(t *testing.T) {
		pass, err := helper.HashPassword(password)
		assert.Nil(t, err)
		assert.NotEmpty(t, pass)
		hashPassword = pass
		t.Logf("Password hash: %s", pass)
	})

	t.Run("VerifikasiPassword", func(t *testing.T) {
		result := helper.VerifyPassword(password, hashPassword)
		assert.True(t, result)
		t.Log("Password berhasil diverifikasi: ", result)
	})
}
