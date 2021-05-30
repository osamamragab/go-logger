package logger

const colorReset string = "\x1b[0m"

// Colors describes console ASCII colors.
type Colors map[string]string

// Get returns the value to the given key.
func (c Colors) Get(key string) string {
	if c == nil {
		return ""
	}
	return c[key]
}

var DefaultColors = Colors{
	"SUCCESS": "\x1b[32m",
	"ERROR":   "\x1b[31m",
	"WARNING": "\x1b[33m",
	"INFO":    "\x1b[34m",
}
