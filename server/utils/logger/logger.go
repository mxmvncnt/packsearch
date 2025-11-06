package logger

import (
	"log"
	"os"

	"github.com/mxmvncnt/packsearch/server/config"
)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var blue = "\033[34m"
var magenta = "\033[35m"
var cyan = "\033[36m"
var gray = "\033[37m"
var white = "\033[97m"

var (
	stdoutLog = log.New(os.Stdout, "", log.LstdFlags)
	stderrLog = log.New(os.Stderr, "", log.LstdFlags)
)

func color(color string, message string) string {
	return color + message + reset
}

func Debug(message string) {
	if config.LogLevel <= config.LogLevelDebug {
		Debugf(message)
	}
}

func Debugf(message string, args ...any) {
	if config.LogLevel <= config.LogLevelDebug {
		stdoutLog.Printf("["+green+"DEBUG"+reset+"] "+message, args...)
	}
}

func Info(message string) {
	if config.LogLevel <= config.LogLevelInfo {
		Infof(message)
	}
}

func Infof(message string, args ...any) {
	if config.LogLevel <= config.LogLevelInfo {
		stdoutLog.Printf("["+cyan+"INFO"+reset+" ] "+message, args...)
	}
}

func Warn(message string) {
	if config.LogLevel <= config.LogLevelWarn {
		Warnf(message)
	}
}

func Warnf(message string, args ...any) {
	if config.LogLevel <= config.LogLevelWarn {
		stdoutLog.Printf("["+yellow+"WARN"+reset+" ] "+message, args...)
	}
}

func Error(message string) {
	if config.LogLevel <= config.LogLevelError {
		Errorf(message)
	}
}

func Errorf(message string, args ...any) {
	if config.LogLevel <= config.LogLevelError {
		stderrLog.Printf("["+red+"ERROR"+reset+"] "+message, args...)
	}
}

func Fatal(message string) {
	if config.LogLevel <= config.LogLevelError {
		Fatalf(message)
	}
}

func Fatalf(message string, args ...any) {
	if config.LogLevel <= config.LogLevelError {
		stderrLog.Fatalf("["+red+"FATAL"+reset+"] "+message, args...)
	}
}
