package repository

import (
	"autentikasi1/cmd/model/internal_db"
	"context"
)

type repositoryUsers interface {
	CreateUsers(ctx context.Context, Users *internal_db.ModelUsers) error
	FindUsersByEmail(ctx context.Context, email string) (error, *internal_db.ModelUsers)
	ChangePasswordByEmail(ctx context.Context, email, passwordHash string) error
	DecrementBalanceByEmail(ctx context.Context, email string, amount int) error
	IncrementBalanceByEmail(ctx context.Context, email string, amount int) error
	SuspendUserByEmail(ctx context.Context, email string) error
	FindUsersOrCreateUsers(ctx context.Context, Password string, Users *internal_db.ModelUsers) (error, *internal_db.ModelUsers)
}
