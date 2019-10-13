package models

import "database/sql"

type Base struct {
	ID        int64          `json:"id" db:"id"`
	CreatedAt sql.NullString `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at" db:"updated_at"`
}

type User struct {
	*Base
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}
