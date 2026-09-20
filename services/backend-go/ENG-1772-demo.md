# ENG-1772: current and referenced component results

This fixture demonstrates component processing provenance in Cronos.

After the baseline collection, this service receives a change while `lib/common`
remains unchanged. The backend should show results processed at the latest repo
commit. The library should show the same requested repo commit while explicitly
identifying the earlier commit that supplied its JSON and checks.

The components are:

- `github.com/pantalasa-cronos/monorepo/services/backend-go`
- `github.com/pantalasa-cronos/monorepo/lib/common`

Only this service's directory changes in the demonstration commit. No library
source, shared dependency, or repository-wide build input changes.
