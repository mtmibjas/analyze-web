package main

import (
	"analyze-web/app/config"
	"analyze-web/app/resolver"
	"analyze-web/app/server"
	"analyze-web/pkg/logger/zap"
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
	loggger := zap.NewLogger(cfg)
	loggger.Init()
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

	<-ctx.Done()
	log.Println("server exiting...")
	cancel()
	os.Exit(0)
}
