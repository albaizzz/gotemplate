package db

import (
	"context"
	"gotemplate/internal/models"

	"github.com/jmoiron/sqlx"
)

func NewUserDb(master, slave *sqlx.DB) IUserDB {
	return &userdb{
		master: master,
		slave:  slave,
	}
}

type userdb struct {
	master *sqlx.DB
	slave  *sqlx.DB
}

type IUserDB interface {
	SQLDb
	Save(ctx context.Context, tx *sqlx.Tx, user models.User) error
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	GetByPhone(ctx context.Context, phone string) (*models.User, error)
}

//BeginMasterx begins the transaction for master db
func (d *userdb) BeginMasterx() *sqlx.Tx {
	tx, _ := d.master.Beginx()
	return tx
}

func (u *userdb) Save(ctx context.Context, tx *sqlx.Tx, user models.User) error {
	qry := `insert into users (name, username, email, phone) values(?, ?,?)`
	_, err := execute(ctx, u.master, tx, qry, user.Name, user.Username, user.Email, user.Phone)
	return err
}

func (u *userdb) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	qry := `select id, name, username, email, phone, created_at, updated_at from users where username = ?`
	err := u.slave.GetContext(ctx, &user, qry, username)
	return &user, err
}

func (u *userdb) GetByPhone(ctx context.Context, phone string) (*models.User, error) {
	var user models.User
	qry := `select id, created_at, updated_at, name, password, username,  email, phone from users where phone = ?`
	err := u.master.GetContext(ctx, &user, qry, phone)
	if err != nil {
		return nil, err
	}
	return &user, err
}
