package svc

import (
	"context"
	"database/sql"
	"gotemplate/internal/consts"
	"gotemplate/internal/db"
	"gotemplate/internal/helpers"
	"gotemplate/internal/models"
	"gotemplate/pkg/log"
	"sync"
)

type usersvc struct {
	Userdb db.IUserDB
}

type IUserSvc interface {
	Save(ctx context.Context, user models.User) (resp models.APIResponse)
	GetByUsername(ctx context.Context, username string) (resp models.APIResponse)
	AuthUser(ctx context.Context, user models.User) models.APIResponse
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

func (u *usersvc) AuthUser(ctx context.Context, user models.User) models.APIResponse {
	var userData *models.User
	userChan := make(chan *models.User, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	go u.validateByPhone(ctx, &wg, userChan, user.Phone)
	go u.validateByUsername(ctx, &wg, userChan, user.Username)
	wg.Wait()
	var idx int = 0
	for elem := range userChan {
		idx++
		if elem != nil {
			userData = elem
			break
		}
		if idx == 2 {
			break
		}
	}
	if userData == nil {
		return helpers.ServiceResp(consts.APINotFoundCode, nil)
	}
	return helpers.ServiceResp(consts.APISuccessCode, "OK")
}

func (u *usersvc) validateByPhone(ctx context.Context, wg *sync.WaitGroup, userChan chan *models.User, phone string) {
	user, err := u.Userdb.GetByPhone(ctx, phone)
	if err != nil {
		if err == sql.ErrNoRows {
			userChan <- nil
		} else {
			log.Error(err.Error())
		}
		wg.Done()
		return
	}
	if user != nil {
		userChan <- user
	}
	wg.Done()
}

//get user by username or phone
func (u *usersvc) validateByUsername(ctx context.Context, wg *sync.WaitGroup, userChan chan *models.User, username string) {
	user, err := u.Userdb.GetByUsername(ctx, username)
	if err != nil {
		if err == sql.ErrNoRows {
			userChan <- nil
		} else {
			log.Error(err.Error())
		}
		wg.Done()
		return
	}
	if user != nil {
		userChan <- user
	}
	wg.Done()
}
