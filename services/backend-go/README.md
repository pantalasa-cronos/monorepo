# backend-go

A tiny Go HTTP service. Returns a greeting and some metadata; exists primarily as a realistic component shape for Lunar monorepo experiments.

## Installation

```bash
cd services/backend-go
go build -o backend-go
```

## Usage

```bash
./backend-go
curl http://localhost:8080/
```

The service listens on `:8080` and exposes `/` and `/healthz`.

## Tier

Tier 1, PII-adjacent. See `compliance/data-classification.md`.

## Contributing

Changes follow the repo-level CODEOWNERS and the review process described in `../../compliance/policies/`.
