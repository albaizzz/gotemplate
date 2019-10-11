package v1

import (
	"gotemplate/context"
	"gotemplate/internal/api/v1/handlers"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/syariatifaris/redisc"
)

// RegistryRoute old api version support
func RegistryRoute(router gin.IRouter, dbMaster *sqlx.DB, dbSlave *sqlx.DB, redis redisc.Redis, app *context.AppContext) {
	health := handlers.NewHealthCheck()
	router.GET("/health/check", health.Info)
}
