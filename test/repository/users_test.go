package repository

import (
	"autentikasi1/cmd/helper"
	"autentikasi1/cmd/model/internal_db"
	repositoryimplementation "autentikasi1/cmd/repository/repository_impl"
	"autentikasi1/configs"
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	repoUsers = repositoryimplementation.NewRepositoryUsers(configs.GetConnectionDB())
	Ctx       = context.Background()
)

func TestUsers(t *testing.T) {
	PasswordHash, err := helper.HashPassword("arman123")
	email := "demo@demo.com"
	assert.Nil(t, err)

	err, username := helper.GenerateApiKey()
	assert.Nil(t, err)

	key, err := helper.GenerateApiUsername()
	assert.Nil(t, err)

	au := sql.NullString{
		String: username,
		Valid:  true,
	}

	ak := sql.NullString{
		String: key,
		Valid:  true,
	}

	Users := &internal_db.ModelUsers{
		Fullname:     "Alex",
		Email:        email,
		Balance:      1000,
		PasswordHash: PasswordHash,
		ApiUsername:  au,
		ApiKey:       ak,
	}

	err = repoUsers.CreateUsers(Ctx, Users)
	assert.Nil(t, err)

	t.Run("Ambil Data Pengguna", func(t *testing.T) {
		err, user := repoUsers.FindUsersByEmail(Ctx, email)
		assert.Nil(t, err)
		t.Log("Pengguna: ", user)
	})

	t.Run("Ubah password", func(t *testing.T) {
		PH, err := helper.HashPassword("Indoksia")
		assert.Nil(t, err)
		err = repoUsers.ChangePasswordByEmail(Ctx, email, PH)
		assert.Nil(t, err)
	})

	t.Run("Tambah saldo", func(t *testing.T) {
		err = repoUsers.IncrementBalanceByEmail(Ctx, email, 5000)
		assert.Nil(t, err)

		err, user := repoUsers.FindUsersByEmail(Ctx, email)
		assert.Nil(t, err)
		t.Log("Saldo Setelah ditambah: ", user.Balance)
	})

	t.Run("Kurangi saldo", func(t *testing.T) {
		err = repoUsers.DecrementBalanceByEmail(Ctx, email, 1000)
		assert.Nil(t, err)

		err, user := repoUsers.FindUsersByEmail(Ctx, email)
		assert.Nil(t, err)
		t.Log("Saldo setelah dikurangi: ", user.Balance)
	})

	t.Run("Suspend Pengguna", func(t *testing.T) {
		err = repoUsers.SuspendUserByEmail(Ctx, email)
		assert.Nil(t, err)

		err, user := repoUsers.FindUsersByEmail(Ctx, email)
		assert.Nil(t, err)
		t.Log("Status Pengguna: ", user.Suspend)
	})

}
