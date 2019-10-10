package containers

import (
	"gotemplate/pkg/sqldb"
)

func RegistryAppServer() *gin.Engine {

	db := sqldb.NewMaria("default")
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
	// old.RegistryRoute(router, db)
	// v1.RegistryRoute(router, db)

	return router

}
