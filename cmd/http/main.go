package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gotemplate/containers"

	"github.com/arkeus/core/log"
	"github.com/spf13/viper"
)

func RegistryHttpServer() {

	srv := http.Server{
		Addr:         fmt.Sprintf(":%s", viper.GetString("app.http")),
		Handler:      containers.RegistryAppServer(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Info(fmt.Sprintf("starting http server port: %s", viper.GetString("app.port")))

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	select {
	case <-ctx.Done():
		log.Info("server shutdown of 5 second")

	}

	log.Info("server exiting")

}
