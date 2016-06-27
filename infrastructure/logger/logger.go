package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// Level defines the logging levels.
type Level uint8

// Level enum to define the different logging levels.
const (
	Critical Level = iota
	Error
	Warning
	Notice
	Info
	Debug
)

// Logger manages several standard library loggers with different levels.
type Logger struct {
	level  Level
	logger *log.Logger
}

// New creates a new logger at the given level.
func New(level Level) *Logger {
	logger := Logger{
		level: level,
	}
	logger.logger = log.New(&logger, "", log.Lshortfile)
	return &logger
}

// ParseLevel parses a level from string to log level.
func ParseLevel(level string) Level {
	switch level {
	case "CRITICAL":
		return Critical
	case "ERROR":
		return Error
	case "WARNING":
		return Warning
	case "NOTICE":
		return Notice
	case "INFO":
		return Info
	case "DEBUG":
		return Debug
	default:
		return Debug
	}
}

// Write provides the interface for the standard library logger.
func (l *Logger) Write(bytes []byte) (int, error) {
	return fmt.Fprintf(os.Stderr, "%v %v", time.Now().UTC().Format(time.RFC3339), string(bytes))
}

// Criticalf logs a critical message.
func (l *Logger) Criticalf(format string, v ...interface{}) {
	if l.level >= Critical {
		format = "[CRITICAL] " + format
		l.logger.Output(2, fmt.Sprintf(format, v...))
	}
}

// Errorf logs an error message.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if l.level >= Error {
		format = "[ERROR] " + format
		l.logger.Output(2, fmt.Sprintf(format, v...))
	}
}

// Warningf logs a warning message.
func (l *Logger) Warningf(format string, v ...interface{}) {
	if l.level >= Warning {
		format = "[WARNING] " + format
		l.logger.Output(2, fmt.Sprintf(format, v...))
	}
}

// Noticef logs a notice message.
func (l *Logger) Noticef(format string, v ...interface{}) {
	if l.level >= Notice {
		format = "[NOTICE] " + format
		l.logger.Output(2, fmt.Sprintf(format, v...))
	}
}

// Infof logs an info message.
func (l *Logger) Infof(format string, v ...interface{}) {
	if l.level >= Info {
		format = "[INFO] " + format
		l.logger.Output(2, fmt.Sprintf(format, v...))
	}
}

// Debugf logs a debug message.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if l.level >= Debug {
		format = "[DEBUG] " + format
		l.logger.Output(2, fmt.Sprintf(format, v...))
	}
}
