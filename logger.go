package logger

import (
	"fmt"
	"io"
	"time"
)

// Colors describes console ASCII colors
type Colors struct {
	Success, Error string
	Warn, Info     string
	Time           string
}

var (
	colorReset = "\x1b[0m"

	colors = Colors{
		Success: "\x1b[32m",
		Error:   "\x1b[31m",
		Warn:    "\x1b[33m",
		Info:    "\x1b[34m",
		Time:    "\x1b[36m",
	}
)

func format(title, color, msg string, vars ...interface{}) string {
	date := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	text := "(" + title + ")"

	if color != "" {
		date = colors.Time + date + colorReset
		text = color + text + colorReset
	}

	return date + " " + text + ": " + fmt.Sprintf(msg, vars...)
}

func setColor(old, new string) string {
	if new == "" {
		return old
	}

	return new
}

// SetColors override default ASCII colors
func SetColors(c *Colors) {
	colors.Success = setColor(colors.Success, c.Success)
	colors.Error = setColor(colors.Error, c.Error)
	colors.Warn = setColor(colors.Warn, c.Warn)
	colors.Info = setColor(colors.Info, c.Info)
	colors.Time = setColor(colors.Time, c.Time)
}

// Success prints formatted success message to standard output
func Success(msg string, a ...interface{}) {
	fmt.Println(format("SUCCESS", colors.Success, msg, a...))
}

// Error prints formatted error message to standard output
func Error(msg string, a ...interface{}) {
	fmt.Println(format("ERROR", colors.Error, msg, a...))
}

// Warn prints formatted warning message to standard output
func Warn(msg string, a ...interface{}) {
	fmt.Println(format("WARNING", colors.Warn, msg, a...))
}

// Info prints formatted info message to standard output
func Info(msg string, a ...interface{}) {
	fmt.Println(format("INFO", colors.Info, msg, a...))
}

// Fsuccess prints formatted success message to w
func Fsuccess(w io.Writer, msg string, a ...interface{}) {
	fmt.Fprintln(w, format("SUCCESS", "", msg, a...))
}

// Ferror prints formatted error message to w
func Ferror(w io.Writer, msg string, a ...interface{}) {
	fmt.Fprintln(w, format("ERROR", "", msg, a...))
}

// Fwarn prints formatted warning message to w
func Fwarn(w io.Writer, msg string, a ...interface{}) {
	fmt.Fprintln(w, format("WARNING", "", msg, a...))
}

// Finfo prints formatted info message to w
func Finfo(w io.Writer, msg string, a ...interface{}) {
	fmt.Fprintln(w, format("INFO", "", msg, a...))
}
