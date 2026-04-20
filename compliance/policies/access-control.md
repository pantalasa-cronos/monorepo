---
policy_id: "POL-AC-001"
owner: "security@pantalasa.org"
last_reviewed: "2026-01-15"
effective_date: "2025-10-01"
---

# Access Control Policy

## Purpose

Define how access to Pantalasa Cronos production systems is requested, granted, reviewed, and revoked.

## Scope

Applies to all production environments, source code, and engineering tooling for services running on the `cronos` infrastructure.

## Policy

### Principle of Least Privilege

All access grants must follow the principle of least privilege. Default permissions are none; elevation requires written business justification.

### Access Request Process

1. Employee submits an access request ticket.
2. Manager approves the business justification.
3. Security reviews and issues time-bound credentials.
4. Access is automatically revoked after 90 days unless renewed.

### Review Cadence

Quarterly access reviews are conducted by the Security team. Access grants without documented justification are revoked during the review.

## Enforcement

Violations are reviewed by the Security team and may result in revocation of access, discipline, or termination.
