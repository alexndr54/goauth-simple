package repository

import (
	"autentikasi1/cmd/model/internal_db"
	"context"
)

type repositoryReset interface {
	CreateResetPassword(ctx context.Context, ResetPassword *internal_db.ModelResetPassword) error
	FindResetPasswordByEmail(ctx context.Context, email string) (error, *internal_db.ModelResetPassword)
	DeleteResetPasswordByEmail(ctx context.Context, email string) error
	FindResetPasswordByToken(ctx context.Context, token string) (error, *internal_db.ModelResetPassword)
}
