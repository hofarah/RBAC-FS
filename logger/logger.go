package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func init() {
	logAddr := "/tmp/app.log"
	_, err := os.ReadFile(logAddr)
	if err != nil {
		os.Create(logAddr)
	}
	var level = zapcore.DebugLevel
	logger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(level),
		OutputPaths: []string{"/tmp/app.log"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "message",
			LevelKey:     "level",
			EncodeLevel:  zapcore.CapitalLevelEncoder,
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}.Build()
	zap.ReplaceGlobals(logger)
}
