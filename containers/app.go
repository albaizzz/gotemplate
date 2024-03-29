package containers

import (
	"fmt"
	"gotemplate/context"
	"gotemplate/internal/consts"
	"gotemplate/pkg/log"
	"gotemplate/pkg/sqldb"

	v1 "gotemplate/internal/api/v1"

	"gotemplate/pkg/env"

	"github.com/syariatifaris/redisc"

	"github.com/gin-gonic/gin"
)

func RegistryAppServer() *gin.Engine {
	envi := env.Get(consts.EnvKey)
	//Init Application configuration Context
	App, err := context.NewAppCtx(envi)
	//init DB Master
	dbMaster, err := sqldb.NewMaria(&sqldb.MariaConfig{Host: App.Config.DB.Maria.MasterDB.Host,
		DBName: App.Config.DB.Maria.MasterDB.Name, Pass: App.Config.DB.Maria.MasterDB.Pass,
		User: App.Config.DB.Maria.MasterDB.User, Port: App.Config.DB.Maria.MasterDB.Port})
	if err != nil {
		log.Fatal("asdas")
	}

	//Init DB Slave
	dbSlave, err := sqldb.NewMaria(&sqldb.MariaConfig{Host: App.Config.DB.Maria.SlaveDB.Host,
		DBName: App.Config.DB.Maria.SlaveDB.Name, Pass: App.Config.DB.Maria.SlaveDB.Pass,
		User: App.Config.DB.Maria.SlaveDB.User, Port: App.Config.DB.Maria.SlaveDB.Port})

	if err != nil {
		log.Fatal("Unable to connect slave db")
	}

	//redis cluster
	redconf := App.Config.DB.Redis
	rconf := &redisc.Config{
		Host:               redconf.Host,
		RetryCount:         redconf.RetryCount,
		RetryDuration:      redconf.RetryDuration,
		MaxActive:          redconf.MaxActive,
		IdleTimeout:        redconf.IdleTimeout,
		MaxIdle:            redconf.IdleTimeout,
		DialConnectTimeout: redconf.DialConnectionTimeout,
	}

	redisc, err := redisc.New(rconf)
	if err != nil {
		log.Fatal(fmt.Sprintf("unable to connect to redis cluster, err =", err.Error()))
	}
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	///logger
	log.SetComponent("http")
	v1.RegistryRoute(router, dbMaster, dbSlave, redisc, App)

	return router

}
