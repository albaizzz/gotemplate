package v1

import (
	"gotemplate/context"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

// RegistryRoute old api version support
func RegistryRoute(router gin.IRouter, dbMaster *sqlx.DB, dbSlave *sqlx.DB, app *context.AppContext) {

}
