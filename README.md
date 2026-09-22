# Pantalasa Cronos Monorepo

A polyglot monorepo used as a test environment for Earthly Lunar product exploration around monorepo support.

## Layout

- `services/backend-go/` — a tiny Go HTTP service
- `services/api-python/` — a tiny Python HTTP service
- `compliance/` — organization-wide compliance documentation (DR plan, DR exercises, policy docs)

Each service also maintains its own `compliance/` subdirectory for service-scoped compliance artifacts (data classification, security review, etc.).

## Purpose

This repo exists so we can exercise how Lunar handles:

- Multiple components living in one repository (subdirectory components)
- Shared compliance data referenced by multiple services via `lunar component get-json`
- Per-service compliance artifacts stamped into each service's release bundle
- The code collector and CI collector flows on a realistic monorepo shape

See [services/backend-go/README.md](services/backend-go/README.md) and [services/api-python/README.md](services/api-python/README.md) for per-service details.

## Installation

```bash
git clone https://github.com/pantalasa-cronos/monorepo.git
cd monorepo
```

## Usage

Each service has its own README with run instructions. The CI workflow is a placeholder while Lunar monorepo support is being designed.

## Contributing

Changes to shared compliance go through a documented review process described in `compliance/policies/`. Service changes follow the service's own README.

<!-- Trigger collectors with the new lunar-config -->
