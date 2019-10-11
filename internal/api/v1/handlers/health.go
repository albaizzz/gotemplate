package handlers

import (
	"go-es/helpers"
	"gotemplate/internal/consts"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

func (h *HealthCheck) Info(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, helpers.GetAPIResponse(consts.APISuccessCode, "OK"))
	return
}
