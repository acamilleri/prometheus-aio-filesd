package formatter

import "strings"

var registered = make(map[string]func() Formatter)

// Register Register a formatter
func Register(name string, loadFn func() Formatter) {
	name = strings.ToLower(name)
	registered[name] = loadFn
}
