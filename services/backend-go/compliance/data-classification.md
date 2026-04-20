---
service: "backend-go"
tier: 1
classification: "internal-pii-adjacent"
last_reviewed: "2026-02-28"
reviewer: "brandon@pantalasa.org"
---

# Data Classification — backend-go

## Summary

`backend-go` does not store customer PII directly, but emits request IDs that can be correlated with user sessions. Logs are treated as **PII-adjacent** and retained under the standard 90-day log retention window.

## Data Inventory

| Field | Source | Classification | Retention |
|-------|--------|----------------|-----------|
| `request_id` | Generated | Internal | 90 days |
| `path` | HTTP request | Internal | 90 days |
| `user_agent` | HTTP request | Internal | 90 days |

## Controls

- Logs redact any `Authorization` headers before emission.
- Request IDs are UUIDv4 and are not derivable from user identity.

## Review Cadence

Reviewed every six months or when the service's data flow changes materially.
