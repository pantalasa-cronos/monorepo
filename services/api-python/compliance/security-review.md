---
service: "api-python"
last_review: "2026-03-10"
next_review: "2027-03-10"
reviewer: "dane@pantalasa.org"
outcome: "pass"
---

# Security Review — api-python

## Scope

Lightweight security review focused on dependency posture and container image hygiene for the `api-python` service.

## Findings

- **Dependencies:** Flask 3.0.3 pinned. No critical CVEs in the build graph as of review date.
- **Base image:** Uses `cgr.dev/chainguard/python:latest` (minimal surface).
- **Secrets:** No secrets in source. No external data stores.

## Open Items

None at review close.

## Next Review

Scheduled for 2027-03-10 or sooner on material change.
