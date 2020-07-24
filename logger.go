package logger

import "fmt"

var colors = map[string]string{
	"green":  "\x1b[32m",
	"red":    "\x1b[31m",
	"yellow": "\x1b[33m",
	"blue":   "\x1b[34m",
}

func format(msg, color string) string {
	return color + "[" + msg + "]" + "\x1b[0m" + ":"
}

// Success prints a success message
func Success(msg string, a ...interface{}) {
	fmt.Println(format("SUCCESS", colors["green"]), fmt.Sprintf(msg, a...))
}

// Error prints a error message
func Error(msg string, a ...interface{}) {
	fmt.Println(format("ERROR", colors["red"]), fmt.Sprintf(msg, a...))
}

// Warn prints a warning message
func Warn(msg string, a ...interface{}) {
	fmt.Println(format("WARNING", colors["yellow"]), fmt.Sprintf(msg, a...))
}

// Info prints an info message
func Info(msg string, a ...interface{}) {
	fmt.Println(format("INFO", colors["blue"]), fmt.Sprintf(msg, a...))
}
