package logger

import (
	"fmt"
	"strings"
)

// Colors represents ASCII colors for console
type Colors struct {
	Success, Error, Warn, Info string
}

var (
	colors = Colors{
		Success: "\x1b[32m",
		Error:   "\x1b[31m",
		Warn:    "\x1b[33m",
		Info:    "\x1b[34m",
	}

	colorReset = "\x1b[0m"
)

func formatWithColor(text, color string) string {
	return color + "[" + text + "]" + colorReset + ":"
}

// SetColors override defualt ASCII colors for conosle
func SetColors(c *Colors) {
	c.Success = strings.TrimSpace(c.Success)
	c.Error = strings.TrimSpace(c.Error)
	c.Warn = strings.TrimSpace(c.Warn)
	c.Info = strings.TrimSpace(c.Info)

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
}

// Success prints a success message
func Success(msg string, a ...interface{}) {
	fmt.Println(formatWithColor("SUCCESS", colors.Success), fmt.Sprintf(msg, a...))
}

// Error prints a error message
func Error(msg string, a ...interface{}) {
	fmt.Println(formatWithColor("ERROR", colors.Error), fmt.Sprintf(msg, a...))
}

// Warn prints a warning message
func Warn(msg string, a ...interface{}) {
	fmt.Println(formatWithColor("WARNING", colors.Warn), fmt.Sprintf(msg, a...))
}

// Info prints an info message
func Info(msg string, a ...interface{}) {
	fmt.Println(formatWithColor("INFO", colors.Info), fmt.Sprintf(msg, a...))
}
