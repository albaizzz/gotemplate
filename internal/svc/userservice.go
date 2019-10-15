package svc

import (
	"context"
	"database/sql"
	"gotemplate/internal/consts"
	"gotemplate/internal/db"
	"gotemplate/internal/helpers"
	"gotemplate/internal/models"
)

type usersvc struct {
	Userdb db.IUserDB
}

type IUserSvc interface {
	Save(ctx context.Context, user models.User) (resp models.APIResponse)
	GetByUsername(ctx context.Context, username string) (resp models.APIResponse)
}

func NewUserSvc(userdb db.IUserDB) IUserSvc {
	return &usersvc{Userdb: userdb}
}

func (u *usersvc) Save(ctx context.Context, user models.User) models.APIResponse {
	trx := u.Userdb.BeginMasterx()
	err := u.Userdb.Save(ctx, trx, user)
	if err != nil {
		return helpers.ServiceResp(consts.APIUnknownCode, nil)
	}
	return helpers.ServiceResp(consts.APISuccessCode, nil)
}

func (u *usersvc) GetByUsername(ctx context.Context, username string) models.APIResponse {
	data, err := u.Userdb.GetByUsername(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			return helpers.ServiceResp(consts.APINotFoundCode, nil)
		}
		return helpers.ServiceResp(consts.APIUnknownCode, nil)
	}
	return helpers.ServiceResp(consts.APISuccessCode, data)
}
