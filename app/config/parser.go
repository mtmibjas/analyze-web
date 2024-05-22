package config

func Parse(cfgDir string) *Config {
	dir := getConfigDir(cfgDir)
	return &Config{
		Service: parseServiceConfig(dir),
		Logs:    parseLoggerConfig(dir),
	}
}

func parseServiceConfig(dir string) ServiceConfig {
	cfg := ServiceConfig{}
	parseConfig(dir+"service.yaml", &cfg)
	validateAppConfig(&cfg)
	return cfg
}

func validateAppConfig(cfg *ServiceConfig) {
	if cfg.ServiceName == "" {
		panic("Name is empty in service.yaml")
	}
	if cfg.Port == 0 {
		panic("Port is empty in service.yaml")
	}
}
func parseLoggerConfig(dir string) LoggerConfig {
	log := LoggerConfig{}
	parseConfig(dir+"logger.yaml", &log)
	return log
}
