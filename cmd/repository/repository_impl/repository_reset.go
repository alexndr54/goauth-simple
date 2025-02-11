package repositoryimplementation

import (
	"autentikasi1/cmd/helper"
	"autentikasi1/cmd/model/internal_db"
	"context"
	"database/sql"
	"errors"
	"time"
)

type RepositoryResetImpl struct {
	DB *sql.DB
}

func NewRepositoryReset(db *sql.DB) *RepositoryResetImpl {
	return &RepositoryResetImpl{DB: db}
}

func (r RepositoryResetImpl) CreateResetPassword(ctx context.Context, ResetPassword *internal_db.ModelResetPassword) error {
	query := `INSERT INTO reset_password (email,token) VALUES (?,?)`
	ex, err := r.DB.ExecContext(ctx, query, ResetPassword.Email, ResetPassword.Token)
	if err != nil {
		err, resetPassword := r.FindResetPasswordByEmail(ctx, ResetPassword.Email)
		if err != nil {
			return err
		}

		jakarta, created := helper.GetDateTime(resetPassword.CreatedAt)
		if time.Now().In(jakarta).After(created.Add(10 * time.Minute)) {
			err := r.DeleteResetPasswordByEmail(ctx, ResetPassword.Email)
			if err != nil {
				return err
			}

			err = r.CreateResetPassword(ctx, ResetPassword)
			if err != nil {
				return err
			}

			return nil
		} else {
			return errors.New("Kami telah mengirimkan link reset password ke email anda cek spam/inbox")
		}
	}

	id, err := ex.LastInsertId()
	if err != nil {
		return errors.New("Gagal menambahkan reset password, ID tidak ditemukan")
	}

	if id > 0 {
		return nil
	}

	return errors.New("Gagal menambahkan reset password, error tidak diketahui")
}

func (r RepositoryResetImpl) FindResetPasswordByEmail(ctx context.Context, email string) (error, *internal_db.ModelResetPassword) {
	query := "SELECT id,email,token,created_at FROM reset_password WHERE email = ?"
	result, err := r.DB.QueryContext(ctx, query, email)
	if err != nil {
		return errors.New("Gagal mengambil data reset password"), nil
	}

	if result.Next() {
		var resetPassword internal_db.ModelResetPassword
		err = result.Scan(&resetPassword.Id, &resetPassword.Email, &resetPassword.Token, &resetPassword.CreatedAt)
		if err != nil {
			return errors.New("Gagal scan data reset password"), nil
		}

		return nil, &resetPassword
	}

	return errors.New("Anda belum melakukan permintaan reset password"), nil
}

func (r RepositoryResetImpl) DeleteResetPasswordByEmail(ctx context.Context, email string) error {
	query := "DELETE FROM reset_password WHERE email = ?"
	_, err := r.DB.ExecContext(ctx, query, email)
	if err != nil {
		return errors.New("Gagal menghapus reset password")
	}

	return nil
}

func (r RepositoryResetImpl) FindResetPasswordByToken(ctx context.Context, token string) (error, *internal_db.ModelResetPassword) {
	query := "SELECT id,email,token,created_at FROM reset_password WHERE token = ?"
	result, err := r.DB.QueryContext(ctx, query, token)
	if err != nil {
		return errors.New("Kesalahan, Token anda tidak valid"), nil
	}

	if result.Next() {
		var resetPassword internal_db.ModelResetPassword
		err = result.Scan(&resetPassword.Id, &resetPassword.Email, &resetPassword.Token, &resetPassword.CreatedAt)
		if err != nil {
			return errors.New("Gagal scan data reset password"), nil
		}

		jakarta, created := helper.GetDateTime(resetPassword.CreatedAt)
		if time.Now().In(jakarta).After(created.Add(10 * time.Minute)) {
			err = r.DeleteResetPasswordByEmail(ctx, resetPassword.Email)
			if err != nil {
				return err, nil
			}

			return errors.New("Link anda sudah kadaluwarsa"), nil
		}

		return nil, &resetPassword
	}

	return errors.New("Token anda tidak valid"), nil
}
