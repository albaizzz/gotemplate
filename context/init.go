package context

import (
	"errors"
	"fmt"
	config "gotemplate/cfg"
	"gotemplate/internal/consts"
	"gotemplate/pkg/file"
)

//NewAppCtx initialize application context
//args:f
//	env: environment variable
//returns:
//	error operation
func NewAppCtx(env string) (app *AppContext, ferr error) {
	app = new(AppContext)
	var fpath []string
	switch env {
	case consts.EnvProduction:
		fpath = []string{consts.FilesStagProdPath}
	case consts.EnvStaging:
		fpath = []string{consts.FilesStagProdPath}
	default:
		fpath = consts.FilesDevelPaths
	}
	//init logging
	// slog.Init(&slog.Config{LogLevel: slog.LevelDebug, Output: os.Stdout})
	//init config
	cfg, err := readCfg("app.yaml", fpath...)
	if err != nil {
		ferr = err
		return
	}
	app.Config = cfg
	//init messages
	// err = msg.InitFromPaths("msgs.yaml", fpath...)
	// if err != nil {
	// 	ferr = err
	// 	return
	// }
	//init errors
	// err = oerrors.InitFromPaths("errors.yaml", fpath...)
	if err != nil {
		ferr = err
		return
	}
	//init datadog
	// addr := fmt.Sprintf("%s:%d", app.Config.Monitor.Address, app.Config.Monitor.Port)
	// ddog.Init(ddog.Config{DatadogConfig: &ddog.DatadogConfig{
	// 	AgentAddress: addr,
	// 	AppName:      app.Config.Common.AppName,
	// }})
	return
}

//readCfg reads the configuration from file
//args:
//	fname: filename
//	ps: full path of possible configuration files
//returns:
//	*config.Configuration: configuration ptr object
//	error: error operation
func readCfg(fname string, ps ...string) (*config.Configuration, error) {
	var cfg *config.Configuration
	for _, p := range ps {
		f := fmt.Sprint(p, fname)
		err := file.ReadFromYAML(f, &cfg)
		if err != nil {
			continue
		}
		break
	}
	if cfg == nil {
		return nil, errors.New("no confif file found")
	}
	return cfg, nil
}
