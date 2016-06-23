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
func ParseLevel(level string) (Level, error) {
	switch level {
	case "CRITICAL":
		return Critical, nil
	case "ERROR":
		return Error, nil
	case "WARNING":
		return Warning, nil
	case "NOTICE":
		return Notice, nil
	case "INFO":
		return Info, nil
	case "DEBUG":
		return Debug, nil
	default:
		return Debug, fmt.Errorf("invalid logging level: %v", level)
	}
}

// Write provides the interface for the standard library logger.
func (l *Logger) Write(bytes []byte) (int, error) {
	return fmt.Fprintf(os.Stderr, "%v %v", time.Now().UTC().Format(time.RFC3339), string(bytes))
}

// output will print a message with a given level.
func (l *Logger) output(level Level, prefix string, format string, v ...interface{}) {
	if level <= l.level {
		format = fmt.Sprintf("%v %v", prefix, format)
		l.logger.Printf(format, v...)
	}
}

// Criticalf logs a critical message.
func (l *Logger) Criticalf(format string, v ...interface{}) {
	l.output(Critical, "[CRITICAL]", format, v...)
}

// Errorf logs an error message.
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.output(Error, "[ERROR]", format, v...)
}

// Warningf logs a warning message.
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.output(Warning, "[WARNING]", format, v...)
}

// Noticef logs a notice message.
func (l *Logger) Noticef(format string, v ...interface{}) {
	l.output(Notice, "[NOTICE]", format, v...)
}

// Infof logs an info message.
func (l *Logger) Infof(format string, v ...interface{}) {
	l.output(Info, "[INFO]", format, v...)
}

// Debugf logs a debug message.
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.output(Debug, "[DEBUG]", format, v...)
}
