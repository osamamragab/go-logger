package logger

import "fmt"

var colors = map[string]string{
	"reset":  "\x1b[0m",
	"green":  "\x1b[32m",
	"red":    "\x1b[31m",
	"yellow": "\x1b[33m",
	"blue":   "\x1b[34m",
}

// Success prints a success message
func Success(msg string, a ...interface{}) {
	fmt.Println(colors["green"]+"[SUCCESS]"+colors["reset"]+":", fmt.Sprintf(msg, a...))
}

// Error prints a error message
func Error(msg string, a ...interface{}) {
	fmt.Println(colors["red"]+"[ERROR]"+colors["reset"]+":", fmt.Sprintf(msg, a...))
}

// Warn prints a warning message
func Warn(msg string, a ...interface{}) {
	fmt.Println(colors["yellow"]+"[WARNING]"+colors["reset"]+":", fmt.Sprintf(msg, a...))
}

// Info prints an info message
func Info(msg string, a ...interface{}) {
	fmt.Println(colors["blue"]+"[INFO]"+colors["reset"]+":", fmt.Sprintf(msg, a...))
}
