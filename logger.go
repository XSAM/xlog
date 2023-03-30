package xlog

import (
	"log"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/stdr"
)

var globalLogger = stdr.New(log.New(os.Stderr, "", log.LstdFlags))

func Logger() logr.Logger {
	return globalLogger
}

// SetLogger is not safe for concurrent use.
func SetLogger(l logr.Logger) {
	globalLogger = l
}

// Info prints info messages
func Info(msg string, keysAndValues ...interface{}) {
	globalLogger.Info(msg, keysAndValues...)
}

// Error prints messages about exceptional states.
func Error(err error, msg string, keysAndValues ...interface{}) {
	globalLogger.Error(err, msg, keysAndValues...)
}

// Debug prints debug messages.
func Debug(msg string, keysAndValues ...interface{}) {
	globalLogger.V(5).Info(msg, keysAndValues...)
}

// Fatal prints info messages and exits
func Fatal(msg string, keysAndValues ...interface{}) {
	globalLogger.Info(msg, keysAndValues...)
	os.Exit(1)
}
