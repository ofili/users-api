package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		OutputPaths: 	[]string{"stdout"},
		Level: 			zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: 		"json",
		EncoderConfig: 	zapcore.EncoderConfig{
			LevelKey: 	"level",
			TimeKey: 	"time",
			MessageKey: "msg",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func GetLogger () {
	return
}

func Info(msg string, tag ...zap.Field) {
	log.Info(msg, tag...)
	log.Sync()
}

func Error(msg string, err error, tag ...zap.Field) {
	if err != nil {
		tag = append(tag, zap.NamedError("error", err))
	}
	log.Error(msg, tag...)
	log.Sync()
}