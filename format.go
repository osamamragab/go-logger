package logger

import (
	"fmt"
	"time"
)

const (
	timeColor  string = "\x1b[36m"
	timeFormat string = "2006-01-02 15:04:05"
)

// Formatter is the formatting function used to format log messages.
type Formatter func(colors Colors, label, msg string, a ...interface{}) string

// DefaultFormatter is the default formatting function, implements the Formatter type.
func DefaultFormatter(colors Colors, label, msg string, a ...interface{}) string {
	if label == "" {
		return fmt.Sprintf("%s[%s]%s: %s", timeColor, time.Now().Format(timeFormat), colorReset, fmt.Sprintf(msg, a...))
	}
	return fmt.Sprintf("%s[%s]%s %s(%s)%s: %s", timeColor, time.Now().Format(timeFormat), colorReset, colors.Get(label), label, colorReset, fmt.Sprintf(msg, a...))
}
