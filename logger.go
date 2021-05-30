// logger is a simple logging helper.
package logger

import (
	"fmt"
	"io"
	"os"
)

// Logger describes logger's colors and formatter function.
type Logger struct {
	Colors    map[string]string
	Formatter Formatter
}

// NewLogger creates new Logger instance with DefaultColors and
// DefaultFormatter.
func NewLogger() *Logger {
	return &Logger{
		Colors:    DefaultColors,
		Formatter: DefaultFormatter,
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

// Success prints formatted success message to standard output.
func (l *Logger) Success(msg string, a ...interface{}) {
	l.Fsuccess(os.Stdout, msg, a...)
}

// Error prints formatted error message to standard error.
func (l *Logger) Error(msg string, a ...interface{}) {
	l.Ferror(os.Stderr, msg, a...)
}

// Warn prints formatted warning message to standard output.
func (l *Logger) Warn(msg string, a ...interface{}) {
	l.Fwarn(os.Stdout, msg, a...)
}

// Info prints formatted info message to standard output.
func (l *Logger) Info(msg string, a ...interface{}) {
	l.Finfo(os.Stdout, msg, a...)
}

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
