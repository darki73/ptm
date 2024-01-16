package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"strings"
)

// FieldsMap is the type of the map that is passed to the log functions.
type FieldsMap map[string]interface{}

// Level is the type of the log level.
type Level uint32

const (
	// LevelTrace is the log level for tracing.
	LevelTrace = iota
	// LevelDebug is the log level for debugging.
	LevelDebug
	// LevelInfo is the log level for informational messages.
	LevelInfo
	// LevelWarn is the log level for warnings.
	LevelWarn
	// LevelError is the log level for errors.
	LevelError
	// LevelFatal is the log level for fatal errors.
	LevelFatal
	// LevelPanic is the log level for panics.
	LevelPanic
)

// init initializes the logger.
func init() {
	logrus.SetLevel(logrus.TraceLevel)
}

// SetOutput sets the output writer for the logger.
func SetOutput(writer io.Writer) {
	logrus.SetOutput(writer)
}

// SetLevel sets the log level for the logger.
func SetLevel(newLevel Level) {
	var level logrus.Level

	switch newLevel {
	case LevelTrace:
		level = logrus.TraceLevel
	case LevelDebug:
		level = logrus.DebugLevel
	case LevelInfo:
		level = logrus.InfoLevel
	case LevelWarn:
		level = logrus.WarnLevel
	case LevelError:
		level = logrus.ErrorLevel
	case LevelFatal:
		level = logrus.FatalLevel
	case LevelPanic:
		level = logrus.PanicLevel
	}

	logrus.SetLevel(level)
}

// ParseLevel parses the given string and returns the corresponding log level.
func ParseLevel(level string) (Level, error) {
	switch strings.ToLower(level) {
	case "t", "trace":
		return LevelTrace, nil
	case "d", "debug":
		return LevelDebug, nil
	case "i", "info":
		return LevelInfo, nil
	case "w", "warn", "warning":
		return LevelWarn, nil
	case "e", "err", "error":
		return LevelError, nil
	case "f", "fatal":
		return LevelFatal, nil
	case "p", "panic":
		return LevelPanic, nil
	default:
		return 0, fmt.Errorf("unknown level: %s", level)
	}
}

// Trace logs a message at the trace level.
func Trace(message string) {
	logrus.Trace(message)
}

// Tracef logs a formatted message at the trace level.
func Tracef(format string, args ...interface{}) {
	logrus.Tracef(format, args...)
}

// TraceWithFields logs a message at the trace level with additional fields.
func TraceWithFields(message string, fields FieldsMap) {
	logrus.WithFields(logrus.Fields(fields)).Trace(message)
}

// TracefWithFields logs a formatted message at the trace level with additional fields.
func TracefWithFields(format string, fields FieldsMap, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Tracef(format, args...)
}

// Debug logs a message at the debug level.
func Debug(message string) {
	logrus.Debug(message)
}

// Debugf logs a formatted message at the debug level.
func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

// DebugWithFields logs a message at the debug level with additional fields.
func DebugWithFields(message string, fields FieldsMap) {
	logrus.WithFields(logrus.Fields(fields)).Debug(message)
}

// DebugfWithFields logs a formatted message at the debug level with additional fields.
func DebugfWithFields(format string, fields FieldsMap, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Debugf(format, args...)
}

// Info logs a message at the info level.
func Info(message string) {
	logrus.Info(message)
}

// Infof logs a formatted message at the info level.
func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

// InfoWithFields logs a message at the info level with additional fields.
func InfoWithFields(message string, fields FieldsMap) {
	logrus.WithFields(logrus.Fields(fields)).Info(message)
}

// InfofWithFields logs a formatted message at the info level with additional fields.
func InfofWithFields(format string, fields FieldsMap, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Infof(format, args...)
}

// Warn logs a message at the warn level.
func Warn(message string) {
	logrus.Warn(message)
}

// Warnf logs a formatted message at the warn level.
func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

// WarnWithFields logs a message at the warn level with additional fields.
func WarnWithFields(message string, fields FieldsMap) {
	logrus.WithFields(logrus.Fields(fields)).Warn(message)
}

// WarnfWithFields logs a formatted message at the warn level with additional fields.
func WarnfWithFields(format string, fields FieldsMap, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Warnf(format, args...)
}

// Error logs a message at the error level.
func Error(message string) {
	logrus.Error(message)
}

// Errorf logs a formatted message at the error level.
func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

// ErrorWithFields logs a message at the error level with additional fields.
func ErrorWithFields(message string, fields FieldsMap) {
	logrus.WithFields(logrus.Fields(fields)).Error(message)
}

// ErrorfWithFields logs a formatted message at the error level with additional fields.
func ErrorfWithFields(format string, fields FieldsMap, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Errorf(format, args...)
}

// Fatal logs a message at the fatal level.
func Fatal(message string) {
	logrus.Fatal(message)
}

// Fatalf logs a formatted message at the fatal level.
func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

// FatalWithFields logs a message at the fatal level with additional fields.
func FatalWithFields(message string, fields FieldsMap) {
	logrus.WithFields(logrus.Fields(fields)).Fatal(message)
}

// FatalfWithFields logs a formatted message at the fatal level with additional fields.
func FatalfWithFields(format string, fields FieldsMap, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Fatalf(format, args...)
}

// Panic logs a message at the panic level.
func Panic(message string) {
	logrus.Panic(message)
}

// Panicf logs a formatted message at the panic level.
func Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

// PanicWithFields logs a message at the panic level with additional fields.
func PanicWithFields(message string, fields FieldsMap) {
	logrus.WithFields(logrus.Fields(fields)).Panic(message)
}

// PanicfWithFields logs a formatted message at the panic level with additional fields.
func PanicfWithFields(format string, fields FieldsMap, args ...interface{}) {
	logrus.WithFields(logrus.Fields(fields)).Panicf(format, args...)
}
