package internal_db

import "time"

type ModelResetPassword struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}
