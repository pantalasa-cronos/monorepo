---
service: "api-python"
tier: 2
classification: "internal"
last_reviewed: "2026-03-05"
reviewer: "brandon@pantalasa.org"
---

# Data Classification — api-python

## Summary

`api-python` is an internal Tier 2 service that serves aggregate metadata. It does not access customer data or PII directly; all handled data is considered **internal-only**.

## Data Inventory

| Field | Source | Classification | Retention |
|-------|--------|----------------|-----------|
| `request_id` | Generated | Internal | 90 days |
| `path` | HTTP request | Internal | 90 days |

## Controls

- No `Authorization` headers are accepted or logged.
- No persistent storage; stateless service.

## Review Cadence

Reviewed annually or on material change.
