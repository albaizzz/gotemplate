package db

import (
	"context"
	"gotemplate/internal/models"

	"github.com/jmoiron/sqlx"
)

func NewUserDb(master, slave *sqlx.DB) User {
	return &user{
		master: master,
		slave:  slave,
	}
}

type user struct {
	master *sqlx.DB
	slave  *sqlx.DB
}

type User interface {
	Save(ctx context.Context, tx *sqlx.Tx, user models.User) error
	GetByUsername(ctx context.Context, username string) (*models.User, error)
}

func (u *user) Save(ctx context.Context, tx *sqlx.Tx, user models.User) error {
	qry := `insert into users (username, email, phone) values(?, ?,?)`
	_, err := execute(ctx, u.master, tx, qry, user.Username, user.Email, user.Phone)
	return err
}

func (u *user) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	qry := `select id, username, email, phone from users where username = ?`
	err := u.slave.GetContext(ctx, &user, qry, username)
	return &user, err
}
