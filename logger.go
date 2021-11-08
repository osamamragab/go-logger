// Package logger is a simple logging helper.
package logger

import (
	"fmt"
	"io"
	"os"
)

// Logger contains logging settings.
type Logger struct {
	// ASCII colors.
	Colors Colors
	// formatting function.
	Formatter Formatter

	DebugMode bool

	OutputSuccess io.Writer
	OutputError   io.Writer
	OutputWarn    io.Writer
	OutputInfo    io.Writer
	OutputDebug   io.Writer
}

// NewLogger creates new Logger instance with defaults set.
func NewLogger() *Logger {
	return &Logger{
		Colors:        DefaultColors,
		Formatter:     DefaultFormatter,
		DebugMode:     false,
		OutputSuccess: os.Stdout,
		OutputError:   os.Stderr,
		OutputWarn:    os.Stdout,
		OutputInfo:    os.Stdout,
		OutputDebug:   os.Stdout,
	}
}

// Printf prints formatted log message to standard output.
func (l *Logger) Printf(format string, a ...interface{}) (int, error) {
	return fmt.Printf("%s", l.Formatter(nil, "", format, a...))
}

// Fsuccess prints formatted success message to w.
func (l *Logger) Fsuccess(w io.Writer, msg string, a ...interface{}) {
	fmt.Fprintln(w, l.Formatter(l.Colors, "SUCCESS", msg, a...))
}

// Ferror prints formatted error message to w.
func (l *Logger) Ferror(w io.Writer, msg string, a ...interface{}) {
	fmt.Fprintln(w, l.Formatter(l.Colors, "ERROR", msg, a...))
}

// Fwarn prints formatted warning message to w.
func (l *Logger) Fwarn(w io.Writer, msg string, a ...interface{}) {
	fmt.Fprintln(w, l.Formatter(l.Colors, "WARNING", msg, a...))
}

// Finfo prints formatted info message to w.
func (l *Logger) Finfo(w io.Writer, msg string, a ...interface{}) {
	fmt.Fprintln(w, l.Formatter(l.Colors, "INFO", msg, a...))
}

// Fdebug prints formatted debug message to w,
// if DebugMode is set to true.
func (l *Logger) Fdebug(w io.Writer, msg string, a ...interface{}) {
	if l.DebugMode {
		fmt.Fprintln(w, l.Formatter(l.Colors, "DEBUG", msg, a...))
	}
}

// Success prints formatted success message to OutputSuccess.
func (l *Logger) Success(msg string, a ...interface{}) {
	l.Fsuccess(l.OutputSuccess, msg, a...)
}

// Error prints formatted error message to OutputError.
func (l *Logger) Error(msg string, a ...interface{}) {
	l.Ferror(l.OutputError, msg, a...)
}

// Warn prints formatted warning message to OutputWarn.
func (l *Logger) Warn(msg string, a ...interface{}) {
	l.Fwarn(l.OutputWarn, msg, a...)
}

// Info prints formatted info message to OutputInfo.
func (l *Logger) Info(msg string, a ...interface{}) {
	l.Finfo(l.OutputInfo, msg, a...)
}

// Debug prints formatted debug message to OutputDebug,
// if DebugMode is set to true.
func (l *Logger) Debug(msg string, a ...interface{}) {
	l.Fdebug(l.OutputDebug, msg, a...)
}

// DefaultLogger is the default logger used by direct functions.
var DefaultLogger = NewLogger()

// Fsuccess prints formatted success message to w.
func Fsuccess(w io.Writer, msg string, a ...interface{}) {
	DefaultLogger.Fsuccess(w, msg, a...)
}

// Ferror prints formatted error message to w.
func Ferror(w io.Writer, msg string, a ...interface{}) {
	DefaultLogger.Ferror(w, msg, a...)
}

// Fwarn prints formatted warning message to w.
func Fwarn(w io.Writer, msg string, a ...interface{}) {
	DefaultLogger.Fwarn(w, msg, a...)
}

// Finfo prints formatted info message to w.
func Finfo(w io.Writer, msg string, a ...interface{}) {
	DefaultLogger.Finfo(w, msg, a...)
}

// Success prints formatted success message to standard output.
func Success(msg string, a ...interface{}) {
	DefaultLogger.Success(msg, a...)
}

// Error prints formatted error message to standard error.
func Error(msg string, a ...interface{}) {
	DefaultLogger.Error(msg, a...)
}

// Warn prints formatted warning message to standard output.
func Warn(msg string, a ...interface{}) {
	DefaultLogger.Warn(msg, a...)
}

// Info prints formatted info message to standard output.
func Info(msg string, a ...interface{}) {
	DefaultLogger.Info(msg, a...)
}

// Debug prints formatted debug message to standard output,
// if DebugMode is set to true.
func Debug(msg string, a ...interface{}) {
	DefaultLogger.Debug(msg, a...)
}
