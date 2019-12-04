package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

//var sugar *zap.SugaredLogger

func  newZapEncoderConfig()  zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "lineNum",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    CustomLevelEncoder,
		EncodeTime:     TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
}

func CustomLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + level.CapitalString() + "]")
}

func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func init() {
}

func getLogLevel(logLevel string) zapcore.Level {
	var zapLevel zapcore.Level
	switch logLevel {
	case "debug":
		zapLevel = zap.DebugLevel
	case "info":
		zapLevel = zap.InfoLevel
	case "warn":
		zapLevel = zap.WarnLevel
	case "error":
		zapLevel = zap.ErrorLevel
	default:
		zapLevel = zap.InfoLevel
	}
	return zapLevel
}

type Logger struct {
	sugar *zap.SugaredLogger
}

func NewLogger(logFile string, logLevel string) *Logger{
	if logFile == "" {
		logFile = "./logs/output.log"
	}
	// 日志文件切割
	lumberjackLogger := lumberjack.Logger{
		Filename: logFile,
		MaxSize: 64, // MB
		MaxBackups: 3,
		MaxAge: 7,
		Compress: true,
	}

	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(newZapEncoderConfig()), // zapcore.NewJSONEncoder()
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&lumberjackLogger)),
		getLogLevel(logLevel),
	)

	logger := zap.New(core, zap.AddCaller(), zap.Development())
	sugar := logger.Sugar()

	log := &Logger{sugar}
	log.sugar.Infof("init Logger success logger file path %s ", logFile)
	return log
}

func (log *Logger)Debug(args ...interface{}) {
	log.sugar.Debug(args)
}

func (log *Logger)Info(args ...interface{}) {
	log.sugar.Info(args)
}

func (log *Logger)Warn(args ...interface{}) {
	log.sugar.Warn(args)
}

func (log *Logger)Error(args ...interface{}) {
	log.sugar.Error(args)
}

func (log *Logger)Panic(args ...interface{}) {
	log.sugar.Panic(args)
}

func (log *Logger)Fatal(args ...interface{}) {
	log.sugar.Fatal(args)
}

func (log *Logger)Debugf(template string, args ...interface{}) {
	log.sugar.Debugf(template, args)
}

func (log *Logger)Infof(template string, args ...interface{}) {
	log.sugar.Infof(template, args)
}

func (log *Logger)Warnf(template string, args ...interface{}) {
	log.sugar.Warnf(template, args)
}

func (log *Logger)Errorf(template string, args ...interface{}) {
	log.sugar.Error(template, args)
}

func (log *Logger)Panicf(template string, args ...interface{}) {
	log.sugar.Panicf(template, args)
}

func (log *Logger)Fatalf(template string, args ...interface{}) {
	log.sugar.Fatalf(template, args)
}