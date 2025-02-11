package repositoryimplementation

import (
	"autentikasi1/cmd/helper"
	"autentikasi1/cmd/model/internal_db"
	"context"
	"database/sql"
	"errors"
)

type RepositoryUsersImpl struct {
	DB *sql.DB
}

func NewRepositoryUsers(db *sql.DB) *RepositoryUsersImpl {
	return &RepositoryUsersImpl{DB: db}
}

func (r RepositoryUsersImpl) CreateUsers(ctx context.Context, Users *internal_db.ModelUsers) error {
	query := `INSERT INTO users (fullname,email,password_hash,api_username,api_key`
	values := `VALUES (?,?,?,?,?`
	args := []interface{}{Users.Fullname, Users.Email, Users.PasswordHash, Users.ApiUsername, Users.ApiKey}

	if Users.LoginProvider != "" {
		query += `,login_provider`
		values += `,?`
		args = append(args, Users.LoginProvider)
	}

	query += `) ` + values + `)`

	ex, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return errors.New("Gagal menambahkan pengguna")
	}

	id, err := ex.LastInsertId()
	if err != nil {
		return errors.New("Gagal menambahkan pengguna, ID tidak ditemukan")
	}

	if id > 0 {
		return nil
	}

	return errors.New("Gagal menambahkan pengguna, error tidak diketahui")
}

func (r RepositoryUsersImpl) FindUsersByEmail(ctx context.Context, email string) (error, *internal_db.ModelUsers) {
	query := "SELECT id,fullname,email,balance,api_username,api_key,suspend,is_admin,password_hash,created_at,login_provider FROM users WHERE email = ?"
	result, err := r.DB.QueryContext(ctx, query, email)
	if err != nil {
		return errors.New("Gagal mengambil data user"), nil
	}

	if result.Next() {
		var user internal_db.ModelUsers
		err = result.Scan(&user.Id, &user.Fullname, &user.Email, &user.Balance, &user.ApiUsername, &user.ApiKey, &user.Suspend, &user.IsAdmin, &user.PasswordHash, &user.CreatedAt, &user.LoginProvider)
		if err != nil {
			return errors.New("Gagal scan data pengguna"), nil
		}

		return nil, &user
	}

	return errors.New("Pengguna tidak ditemukan"), nil
}

func (r RepositoryUsersImpl) ChangePasswordByEmail(ctx context.Context, email, passwordHash string) error {
	err, getUser := r.FindUsersByEmail(ctx, email)
	if err != nil {
		return err
	}

	query := "UPDATE users SET password_hash = ? WHERE email = ?"
	result, err := r.DB.ExecContext(ctx, query, passwordHash, getUser.Email)
	if err != nil {
		return errors.New("Gagal mengubah password")
	}

	countAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("Password gagal diubah")
	}
	if countAffected > 0 {
		return nil
	}

	return errors.New("Ada kesalahan tidak diketahui ketika update password")
}

func (r RepositoryUsersImpl) DecrementBalanceByEmail(ctx context.Context, email string, amount int) error {
	err, getUser := r.FindUsersByEmail(ctx, email)
	if err != nil {
		return err
	}

	query := "UPDATE users SET balance = balance - ? WHERE email = ?"
	result, err := r.DB.ExecContext(ctx, query, amount, getUser.Email)
	if err != nil {
		return err
	}

	countAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("Saldo gagal untuk di tambahkan")
	}

	if countAffected > 0 {
		return nil
	}

	return errors.New("Ada kesalahan ketika menambahkan saldo")
}

func (r RepositoryUsersImpl) IncrementBalanceByEmail(ctx context.Context, email string, amount int) error {
	err, getUser := r.FindUsersByEmail(ctx, email)
	if err != nil {
		return err
	}

	query := "UPDATE users SET balance = balance + ? WHERE email = ?"
	result, err := r.DB.ExecContext(ctx, query, amount, getUser.Email)
	if err != nil {
		return err
	}

	countAffected, err := result.RowsAffected()
	if err != nil {
		return errors.New("Saldo gagal dikurangi")
	}

	if countAffected > 0 {
		return nil
	}

	return errors.New("Ada kesalahan ketika mengurangi saldo")
}

func (r RepositoryUsersImpl) SuspendUserByEmail(ctx context.Context, email string) error {
	err, getUser := r.FindUsersByEmail(ctx, email)
	if err != nil {
		return err
	}

	query := "UPDATE users SET suspend = ? WHERE email = ?"
	_, err = r.DB.ExecContext(ctx, query, true, getUser.Email)
	if err != nil {
		return errors.New("Ada kesalahan ketika melakukan penangguhan pengguna")
	}

	return nil
}

func (r RepositoryUsersImpl) FindUsersOrCreateUsers(ctx context.Context, Password string, Users *internal_db.ModelUsers) (error, *internal_db.ModelUsers) {
	err, getUser := r.FindUsersByEmail(context.Background(), Users.Email)
	if err != nil {
		// CREATE USER
		// TIDAK ADA USER

		HashPassword, err := helper.HashPassword(Password)
		if err != nil {
			return err, nil
		}
		ApiUsername, err := helper.GenerateApiUsername()
		if err != nil {
			return err, nil
		}
		err, ApiKey := helper.GenerateApiKey()
		if err != nil {
			return err, nil
		}

		Users.ApiUsername = sql.NullString{
			String: ApiUsername,
			Valid:  true,
		}
		Users.ApiKey = sql.NullString{
			String: ApiKey,
			Valid:  true,
		}
		Users.PasswordHash = HashPassword
		err = r.CreateUsers(context.Background(), Users)

		if err != nil {
			return err, nil
		}

		Users.PasswordHash = "true"
		return nil, Users
	} else {
		// USER SUDAH ADA
		return nil, getUser
	}

}
