// Package common is the shared library backend-go depends on. It in turn
// depends on lib/logging, giving the monorepo a two-level dependency chain:
// services/backend-go -> lib/common -> lib/logging.
package common

import "github.com/pantalasa-cronos/monorepo/lib/logging"

// Greeting returns the greeting served by backend-go.
func Greeting() string {
	return Prefixed("hello from lib/common")
}

// Prefixed tags a message with the library's log prefix.
func Prefixed(msg string) string {
	return logging.Prefix("common") + msg
}
