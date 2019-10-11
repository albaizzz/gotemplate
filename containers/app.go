package containers

import (
	"fmt"
	"gotemplate/context"
	"gotemplate/internal/consts"
	"gotemplate/pkg/sqldb"

	"log"

	v1 "gotemplate/internal/api/v1"

	"github.com/syariatifaris/redisc"

	"github.com/gin-gonic/gin"
)

func RegistryAppServer() *gin.Engine {

	// environ := env.Get(consts.EnvKey)
	App, err := context.NewAppCtx(consts.EnvDevelopment)
	dbMaster, err := sqldb.NewMaria(&sqldb.MariaConfig{Host: App.Config.DB.Maria.MasterDB.Host,
		DBName: App.Config.DB.Maria.MasterDB.Name, Pass: App.Config.DB.Maria.MasterDB.Pass,
		User: App.Config.DB.Maria.MasterDB.User, Port: App.Config.DB.Maria.MasterDB.Port})

	if err != nil {
		log.Fatalln("Unable to connect master db")
	}
	fmt.Print(dbMaster)

	dbSlave, err := sqldb.NewMaria(&sqldb.MariaConfig{Host: App.Config.DB.Maria.SlaveDB.Host,
		DBName: App.Config.DB.Maria.SlaveDB.Name, Pass: App.Config.DB.Maria.SlaveDB.Pass,
		User: App.Config.DB.Maria.SlaveDB.User, Port: App.Config.DB.Maria.SlaveDB.Port})

	if err != nil {
		log.Fatalln("Unable to connect slave db")
	}

	fmt.Print(dbSlave)

	//redis
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

	redis, err := redisc.New(rconf)
	if err != nil {
		log.Fatalln("unable to connect to redis cluster, err =", err.Error())
	}

	fmt.Print(redis)

	//cache := bootstrap.RegistryRedis()

	// apm := bootstrap.RegistryTracer("ms_test")

	// kc := koolkit.Config{
	// 	DebugMode:            viper.GetBool("app.debug_mode"),        // true | false
	// 	ServiceName:          viper.GetString("app.name"),            // repo name
	// 	CommandName:          "http",                                 // command name
	// 	ServiceCheckEndpoint: viper.GetString("monitor.apm.address"), // currently we use dogstatsd,
	// }

	// env := viper.GetString("app.env")
	// // selected transform env to koolkit env
	// switch env {
	// case "prod", "production":
	// 	kc.Env = koolkit.ProductionEnv
	// case "stg", "staging":
	// 	kc.Env = koolkit.StagingEnv
	// case "dev", "development":
	// 	kc.Env = koolkit.DevelopmentEnv
	// }

	// // registered koolkit monitor
	// kc.Dependencies = append(kc.Dependencies, db)
	// koolkit.Play(kc)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	// router.Use(middlewares.MiddlewareTracer(apm), middlewares.PanicRecovery())
	v1.RegistryRoute(router, dbMaster, dbSlave, App)

	return router

}
