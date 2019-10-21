package models

import "database/sql"

type Base struct {
	ID        int64          `json:"id" db:"id"`
	CreatedAt sql.NullString `json:"created_at,omitempty" db:"created_at" `
	UpdatedAt sql.NullString `json:"updated_at,omitempty" db:"updated_at"`
}

type User struct {
	*Base
	Name     string `json:"name" db:"name"`
	Password string `json:"password, omitempty"`
	Username string `json:"username" `
	Email    string `json:"email" `
	Phone    string `json:"phone"`
}
