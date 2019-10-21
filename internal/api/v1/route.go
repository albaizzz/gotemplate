package v1

import (
	"gotemplate/context"
	"gotemplate/internal/api/v1/handlers"
	"gotemplate/internal/db"
	"gotemplate/internal/middleware"
	"gotemplate/internal/svc"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/syariatifaris/redisc"
)

// RegistryRoute old api version support
func RegistryRoute(router gin.IRouter, dbMaster *sqlx.DB, dbSlave *sqlx.DB, redis redisc.Redis, app *context.AppContext) {

	userDB := db.NewUserDb(dbMaster, dbSlave)
	usersvc := svc.NewUserSvc(userDB)
	userHandler := handlers.NewUserHandler(usersvc)

	v1 := router.Group("scmlite/v1")

	v1.GET("/users/:user", userHandler.Get)
	v1.POST("/users", userHandler.Save)
	v1.POST("/users/auth", userHandler.Auth)

	health := handlers.NewHealthCheck()
	test := router.Group("/")
	test.Use(middleware.AuthProcess())
	test.GET("/health/check", health.Info)
}
