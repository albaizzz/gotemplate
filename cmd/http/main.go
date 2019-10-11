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

	"github.com/spf13/viper"
)

func RegistryHttpServer() {

	srv := http.Server{
		Addr:         fmt.Sprintf(":%s", viper.GetString("common.app_http")),
		Handler:      containers.RegistryAppServer(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	fmt.Println(fmt.Sprintf("starting http server port: %s", viper.GetString("common.app_port")))

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	select {
	case <-ctx.Done():
		fmt.Println("server shutdown of 5 second")

	}

	fmt.Println("server exiting")

}
