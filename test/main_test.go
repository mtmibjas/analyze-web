package test

import (
	"analyze-web/app/config"
	"analyze-web/pkg/logger/zap"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	cfg := config.Parse("/config")
	loggger := zap.NewLogger(cfg)
	loggger.Init()
	code := m.Run()
	os.Exit(code)
}
