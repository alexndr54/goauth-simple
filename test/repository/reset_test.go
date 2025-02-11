package repository

import (
	"autentikasi1/cmd/model/internal_db"
	repositoryimplementation "autentikasi1/cmd/repository/repository_impl"
	"autentikasi1/configs"
	_ "autentikasi1/init"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResetPassword(t *testing.T) {
	ctx := context.Background()
	repo := repositoryimplementation.NewRepositoryReset(configs.GetConnectionDB())

	t.Run("Tambah", func(t *testing.T) {
		err := repo.CreateResetPassword(ctx, &internal_db.ModelResetPassword{Email: "demo@demo.com", Token: "123456789"})
		assert.Nil(t, err)
	})

	t.Run("Find Reset Password By Email", func(t *testing.T) {
		err, resetPassword := repo.FindResetPasswordByEmail(ctx, "demo@demo.com")
		assert.Nil(t, err)
		t.Log(resetPassword)
	})

	//t.Run("Delete Reset Password By Email", func(t *testing.T) {
	//	err := repo.DeleteResetPasswordByEmail(ctx, "demo@demo.com")
	//	assert.Nil(t, err)
	//})

	t.Log("Test reset password berhasil dilakukan")
}
