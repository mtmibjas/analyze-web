package config

type Config struct {
	Service ServiceConfig
	Logs    LoggerConfig
}

type ServiceConfig struct {
	ServiceName     string `yaml:"name"`
	BaseURL         string `yaml:"base_url"`
	TracingExporter string `yaml:"tracing_exporter"`
	AppEnv          string `yaml:"app_env"`
	Port            int    `yaml:"port"`
	Timeout         int    `yaml:"time_out"`
}

type LoggerConfig struct {
	Log any `yaml:"log"`
}
