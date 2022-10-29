package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	config.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder

	log, err := config.Build()
	if err != nil {
		panic(err)
	}

	logger = log
}

func Info(msg string, args ...any) {
	logger.Info(msg, fields(args...)...)
}

func Warn(msg string, args ...any) {
	logger.Warn(msg, fields(args...)...)
}

func Error(msg string, args ...any) {
	logger.Error(msg, fields(args...)...)
}

func Fatal(msg string, args ...any) {
	logger.Fatal(msg, fields(args...)...)
}

func fields(args ...any) []zapcore.Field {
	var fields []zapcore.Field
	for i := 0; i < len(args); i += 2 {
		fields = append(fields, field(args[i].(string), args[i+1]))
	}

	return fields
}

func field(key string, value any) zapcore.Field {
	var field zapcore.Field

	if v, ok := value.(string); ok && len(v) > 0 {
		field = zap.String(key, value.(string))
		return field
	}
	if v, ok := value.(int64); ok && v != 0 {
		field = zap.Int64(key, value.(int64))
		return field
	}
	if v, ok := value.(int); ok && v != 0 {
		field = zap.Int(key, value.(int))
		return field
	}

	return field
}
