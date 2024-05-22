package main

import (
	"analyze-web/app/config"
	"analyze-web/app/resolver"
	"analyze-web/app/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Init Config
	cfg := config.Parse("/config")

	// tp := tracer.ExposeTracer(cfg)
	ctr := resolver.NewAdapter(cfg)
	srv := server.Run(cfg, ctr)

	// Wait for interrupt signal to gracefully shutdown the server
	sigterm := make(chan os.Signal, 1)

	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	<-sigterm
	log.Println("received interrupt signal...")

	var wait time.Duration = 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), wait)

	// gracefully stop the server
	server.Stop(ctx, srv)
	// tracer.Stop(ctx, tp)

	<-ctx.Done()
	log.Println("server exiting...")
	cancel()
	os.Exit(0)
}
