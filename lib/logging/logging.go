// Package logging is a second-level shared library: lib/common depends on it,
// so services that depend on lib/common reach it only transitively. That is
// the edge a hand-maintained `paths` list misses (ENG-1878).
package logging

import "fmt"

// Prefix renders a log line prefix.
func Prefix(component string) string {
	return fmt.Sprintf("[%s] ", component)
}
