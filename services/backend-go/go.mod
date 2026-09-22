module github.com/pantalasa-cronos/monorepo/services/backend-go

go 1.22

require github.com/pantalasa-cronos/monorepo/lib/common v0.0.0

// Same-repo shared library — resolved from the sibling subdirectory, so a
// change under lib/common/ rebuilds this service even though nothing under
// services/backend-go/ changed.
replace github.com/pantalasa-cronos/monorepo/lib/common => ../../lib/common

// lib/common imports lib/logging, and Go only honours `replace` from the main
// module — so the transitive dependency needs its own replace here too.
require github.com/pantalasa-cronos/monorepo/lib/logging v0.0.0

replace github.com/pantalasa-cronos/monorepo/lib/logging => ../../lib/logging
