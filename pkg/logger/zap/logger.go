package zap

import (
	"analyze-web/app/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/natefinch/lumberjack"
)

var Logger *zap.Logger

type ZapLogger struct {
	Config *config.Config
}

func NewLogger(cfg *config.Config) *ZapLogger {
	return &ZapLogger{
		Config: cfg,
	}
}
func (l *ZapLogger) Init() {
	logConfigs := l.Config.Logs.Logs
	var cores []zapcore.Core
	for _, env := range logConfigs {
		writeSyncer := getLogWriter(env)
		encoder := getEncoder()
		level := getLevelEnabler(env.Level)
		core := zapcore.NewCore(encoder, writeSyncer, level)
		cores = append(cores, core)
	}
	Logger = zap.New(zapcore.NewTee(cores...), zap.AddStacktrace(zapcore.ErrorLevel))
}

func getLogWriter(l config.LogConfig) zapcore.WriteSyncer {

	lumberJackLogger := &lumberjack.Logger{
		Filename:   l.File.Path + l.File.Name, // Log name
		MaxSize:    l.Rotation.MaxSize,        // File content size, MB
		MaxBackups: l.Rotation.MaxBackups,     // Maximum number of old files retained
		MaxAge:     l.Rotation.MaxAge,         // Maximum number of days to keep old files
		Compress:   l.Rotation.Compress,       // Is the file compressed
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	// The format time can be customized
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}

func getLevelEnabler(level string) zapcore.LevelEnabler {
	switch level {
	case "debug":
		return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.DebugLevel
		})
	case "info":
		return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.InfoLevel
		})
	case "warn":
		return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.WarnLevel
		})
	case "error":
		return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.ErrorLevel
		})
	case "panic":
		return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.PanicLevel
		})
	case "fatal":
		return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.FatalLevel
		})
	default:
		return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl == zapcore.InfoLevel
		})
	}
}

func Debug(serviceName string, message ...interface{}) {
	Logger.Debug(serviceName, zap.Any(serviceName, message))
}

func Error(serviceName string, err error) {
	Logger.Error(serviceName, zap.Error(err))
}

func Info(serviceName string, message ...interface{}) {
	Logger.Info(serviceName, zap.Any(serviceName, message))
}

func Warn(serviceName string, message ...interface{}) {
	Logger.Warn(serviceName, zap.Any(serviceName, message))
}

func Panic(serviceName string, message ...interface{}) {
	Logger.Panic(serviceName, zap.Any(serviceName, message))
}

func Fatal(serviceName string, message ...interface{}) {
	Logger.Fatal(serviceName, zap.Any(serviceName, message))
}
