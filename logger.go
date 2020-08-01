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

// SetColors override defualt ASCII colors for conosle
func SetColors(c *Colors) {
	if c.Success != "" {
		colors.Success = c.Success
	}

	if c.Error != "" {
		colors.Error = c.Error
	}

	if c.Warn != "" {
		colors.Warn = c.Warn
	}

	if c.Info != "" {
		colors.Info = c.Info
	}

	if c.Time != "" {
		colors.Time = c.Time
	}
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
