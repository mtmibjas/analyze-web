package server

import (
	"analyze-web/app/config"
	"analyze-web/app/http/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"analyze-web/app/resolver"
)

func Run(cfg *config.Config, ctr *resolver.Resolver) *http.Server {
	container := ctr.Resolve()
	route := router.Init(cfg, container)
	route.Logger.Fatal(route.Start(":" + strconv.Itoa(cfg.Service.Port)))

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(cfg.Service.Port),
		Handler: route,
		// good practice to set timeouts to avoid Slowloris attacks
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	// run our server in a goroutine so that it doesn't block
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println(err)
			panic("Service shutting down unexpectedly...")
		}
	}()

	return srv
}

func Stop(ctx context.Context, srv *http.Server) {
	fmt.Println("Service shutting down...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}
}
