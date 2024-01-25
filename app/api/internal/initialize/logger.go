package initialize

import (
	"dianshang/app/api/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

func SetupLogger() {
	level := zap.NewAtomicLevel()
	level.SetLevel(zapcore.DebugLevel)
	encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:       "message",
		LevelKey:         "level",
		TimeKey:          "time",
		NameKey:          "logger",
		CallerKey:        "caller",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeTime:       CustomTimeEncoder,
		EncodeDuration:   zapcore.StringDurationEncoder,
		EncodeCaller:     zapcore.FullCallerEncoder,
		ConsoleSeparator: "",
	})
	cores := [...]zapcore.Core{
		zapcore.NewCore(encoder, os.Stdout, level),
		zapcore.NewCore(
			encoder,
			zapcore.AddSync(getwritesync()),
			level,
		),
	}

	global.Logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(global.Logger)
	global.Logger.Info("initialize logger success")
}
func getwritesync() zapcore.WriteSyncer {
	lumberjacklogger := &lumberjack.Logger{
		Filename:   global.Config.ZapConfig.Filename,
		MaxSize:    global.Config.ZapConfig.MaxSize,
		MaxAge:     global.Config.ZapConfig.MaxAge,
		MaxBackups: global.Config.ZapConfig.MaxBackups,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberjacklogger)
}
func CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	encoder.AppendString(t.Format("[2006-01-02 15:04:05.000]"))
}
