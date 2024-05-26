package config

type Config struct {
	Service ServiceConfig
	Logs    LoggerConfig
}

type ServiceConfig struct {
	ServiceName     string `yaml:"name"`
	BaseURL         string `yaml:"base_url"`
	WebPath         string `yaml:"web_path"`
	TracingExporter string `yaml:"tracing_exporter"`
	AppEnv          string `yaml:"app_env"`
	Port            int    `yaml:"port"`
	Timeout         int    `yaml:"timeout"`
}

type LogRotation struct {
	MaxSize    int  `yaml:"max_size"`
	MaxBackups int  `yaml:"max_backups"`
	MaxAge     int  `yaml:"max_age"`
	Compress   bool `yaml:"compress"`
}

type LogFile struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
}

type LogConfig struct {
	Level    string      `yaml:"level"`
	File     LogFile     `yaml:"file"`
	Rotation LogRotation `yaml:"rotation"`
}

type LoggerConfig struct {
	Logs []LogConfig `yaml:"logs"`
}
