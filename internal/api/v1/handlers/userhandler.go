package handlers

import (
	"gotemplate/internal/helpers"
	"gotemplate/internal/models"
	"gotemplate/internal/svc"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userSvc svc.IUserSvc
}

func NewUserHandler(userSvc svc.IUserSvc) *UserHandler {
	return &UserHandler{
		userSvc: userSvc,
	}
}

func (u *UserHandler) Get(ctx *gin.Context) {
	username := ctx.Param("user")

	res := u.userSvc.GetByUsername(ctx.Request.Context(), username)
	helpers.APIResponse(ctx, res)
	return
}

func (u *UserHandler) Save(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)

	res := u.userSvc.Save(ctx.Request.Context(), user)
	helpers.APIResponse(ctx, res)
	return
}

func (u *UserHandler) Auth(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)
	res := u.userSvc.AuthUser(ctx.Request.Context(), user)
	helpers.APIResponse(ctx, res)
	return
}
