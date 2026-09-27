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

The service listens on `:8080` (override with the `ADDR` env var) and exposes the following endpoints:

| Method | Path       | Description                          |
| ------ | ---------- | ------------------------------------ |
| GET    | `/`        | Returns a JSON greeting and metadata |
| GET    | `/healthz` | Liveness probe, returns `ok`         |

## Tier

Tier 1, PII-adjacent. See `compliance/data-classification.md`.

## Contributing

Changes follow the repo-level CODEOWNERS and the review process described in `../../compliance/policies/`.

Chained-run test A: 2026-09-26T21:29:38Z
Chained-run test B: 2026-09-26T21:30:09Z
Chained-run test A2: 2026-09-26T23:24:25Z
Chained-run test B2: 2026-09-26T23:24:25Z
Main-control test C: 2026-09-27T00:22:36Z
Chained-run test A3: 2026-09-27T12:44:24Z
Chained-run test B3: 2026-09-27T12:44:24Z
Chained-run test A4: 2026-09-27T16:44:37Z
Chained-run test B4: 2026-09-27T16:44:37Z
