package internal_db

import (
	"database/sql"
	"time"
)

type ModelUsers struct {
	Id            int64          `json:"id"`
	Fullname      string         `json:"fullname" `
	Email         string         `json:"email"`
	Balance       float64        `json:"balance"`
	ApiUsername   sql.NullString `json:"api_username"`
	ApiKey        sql.NullString `json:"api_key"`
	Suspend       bool           `json:"suspend"`
	IsAdmin       bool           `json:"is_admin"`
	PasswordHash  string         `json:"password_hash"`
	LoginProvider string         `json:"login_provider" enum:"GOOGLE,REGISTER"`
	CreatedAt     time.Time      `json:"created_at"`
}
