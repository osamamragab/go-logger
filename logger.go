package logger

import (
	"fmt"
	"time"
)

// Colors represents ASCII colors for console
type Colors struct {
	Success, Error string
	Warn, Info     string
	Time           string
}

var (
	colorEmpty = ""
	colorReset = "\x1b[0m"

	colors = Colors{
		Success: "\x1b[32m",
		Error:   "\x1b[31m",
		Warn:    "\x1b[33m",
		Info:    "\x1b[34m",
		Time:    "\x1b[35m",
	}
)

func format(text, color string) string {
	now := "[" + time.Now().Format("2006-01-02 15:04:05") + "]"
	msg := "(" + text + ")"

	if color != colorEmpty {
		now = colors.Time + now + colorReset
		msg = color + msg + colorReset
	}

	return now + " " + msg + ":"
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

// Success prints a success message
func Success(msg string, a ...interface{}) {
	fmt.Println(format("SUCCESS", colors.Success), fmt.Sprintf(msg, a...))
}

// Error prints a error message
func Error(msg string, a ...interface{}) {
	fmt.Println(format("ERROR", colors.Error), fmt.Sprintf(msg, a...))
}

// Warn prints a warning message
func Warn(msg string, a ...interface{}) {
	fmt.Println(format("WARNING", colors.Warn), fmt.Sprintf(msg, a...))
}

// Info prints an info message
func Info(msg string, a ...interface{}) {
	fmt.Println(format("INFO", colors.Info), fmt.Sprintf(msg, a...))
}
