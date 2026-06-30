// Package common is a shared library used by services in this monorepo.
//
// It models a monorepo dependency scenario: a component defined as a
// subdirectory (services/backend-go) depends on a library in a *sibling*
// subdirectory (lib/common). When this library changes, nothing in
// services/backend-go/ changes — yet backend-go must be rebuilt and
// re-evaluated. Lunar handles that because backend-go declares lib/common/* in
// its component `paths` (see pantalasa-cronos/lunar lunar-config.d/10-components.yml).
package common

// Greeting returns the message served by every service that depends on this
// shared library. Changing this string is a library-only change: it forces a
// rebuild of backend-go without touching anything under services/backend-go/.
func Greeting() string {
	return "hello from lib/common v3 (fan-out: lib + backend-go)"
}
