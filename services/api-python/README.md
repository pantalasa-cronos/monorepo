# api-python

A tiny Python HTTP service built with Flask. Returns a greeting and some metadata; exists primarily as a realistic component shape for Lunar monorepo experiments.

## Installation

```bash
cd services/api-python
pip install -r requirements.txt
```

## Usage

```bash
python main.py
curl http://localhost:8081/
```

The service listens on `:8081` and exposes `/` and `/healthz`.

## Tier

Tier 2, internal. See `compliance/data-classification.md`.

## Contributing

Changes follow the repo-level CODEOWNERS and the review process described in `../../compliance/policies/`.
