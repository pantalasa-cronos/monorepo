---
service: "backend-go"
last_review: "2026-02-20"
next_review: "2026-08-20"
reviewer: "dane@pantalasa.org"
outcome: "pass"
---

# Security Review — backend-go

## Scope

Threat modeling and design review covering the HTTP surface, dependency posture, container image hygiene, and logging practices of the `backend-go` service.

## Findings

- **Dependencies:** No critical CVEs in the build graph as of review date.
- **Base image:** Uses `cgr.dev/chainguard/static:latest`. Base image is ephemeral; no shell present.
- **Logging:** Request logs redact sensitive headers. Compliant with `compliance/policies/access-control.md`.
- **Secrets:** No secrets in source. Environment-injected at deploy time.

## Open Items

None at review close.

## Next Review

Scheduled for 2026-08-20 or sooner if the service introduces authentication flows or external data stores.
