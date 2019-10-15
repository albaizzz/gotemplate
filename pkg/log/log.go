package log

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	// LoggerOpts is the standard logger options
	LoggerOpts = []zap.Option{zap.AddCaller(), zap.AddCallerSkip(2), zap.AddStacktrace(zap.ErrorLevel)}

	encoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     RFC3339NanoEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})

	output   zapcore.WriteSyncer
	logLevel zapcore.Level
	log      *zap.Logger
	logMu    sync.Mutex
)

func init() {
	defaultLog()
}

func createLogger(o zapcore.WriteSyncer, lvl zapcore.Level) {
	core := zapcore.NewCore(encoder, o, lvl)
	log = zap.New(core).WithOptions(LoggerOpts...)
}

func defaultLog() {
	logMu.Lock()
	defer logMu.Unlock()

	output = os.Stderr
	logLevel = zap.InfoLevel
	createLogger(output, logLevel)
}

// Reset will reset the log to the original setup
func Reset() {
	defaultLog()
}

// Core returns core of the logger
func Core() zapcore.Core {
	return log.Core()
}

// DebugMode sets the log level to debug
func DebugMode() {
	logMu.Lock()
	defer logMu.Unlock()

	logLevel = zap.DebugLevel
	createLogger(output, logLevel)
}

// SetOutput sets the log output to whatever that implement the zapcore.WriteSyncer interface.
// This function exist only for testing purpose.
func SetOutput(o zapcore.WriteSyncer) {
	logMu.Lock()
	defer logMu.Unlock()

	output = o
	createLogger(output, logLevel)
}

// SetComponent sets the log component field to the given component param.
func SetComponent(component string) {
	logMu.Lock()
	defer logMu.Unlock()

	log = log.With(zap.String("component", component))
}

// Debug add log entry with or without fields to debug level
func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

// Info add log entry with or without fields to info level
func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

// Warn add log entry with or without fields to warn level
func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
}

// Error add log entry with or without fields to error level
func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

// Fatal add log entry with or without fields to fatal level
func Fatal(msg string, fields ...zap.Field) {
	log.Fatal(msg, fields...)
}

// Panic add log entry with or without fields to panic level
func Panic(msg string, fields ...zap.Field) {
	log.Panic(msg, fields...)
}
