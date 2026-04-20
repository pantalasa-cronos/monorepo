---
rto_minutes: 60
rpo_minutes: 15
last_reviewed: "2026-02-10"
approver: "brandon@pantalasa.org"
owner: "security@pantalasa.org"
---

# Disaster Recovery Plan

## Overview

This document describes the disaster recovery plan for Pantalasa Cronos production systems. It covers recovery objectives, ownership, recovery procedures, and contact information.

## Recovery Objectives

- **Recovery Time Objective (RTO):** 60 minutes for tier-1 services.
- **Recovery Point Objective (RPO):** 15 minutes of data loss tolerated.

## Scope

This plan covers all production services running on the `cronos` infrastructure. Services classify themselves under one of:

- **Tier 1** — customer-facing, 24x7, revenue-critical. Covered by 15-minute RPO.
- **Tier 2** — internal or batch workloads, 60-minute RPO acceptable.
- **Tier 3** — experimental, no RPO guarantee.

## Recovery Procedures

### Database Failover

1. Confirm primary is unreachable via `cronosctl db ping --primary`.
2. Promote replica with `cronosctl db failover --target replica-1`.
3. Update DNS via the internal DNS automation tool.
4. Verify application health with `cronosctl health --service <name>`.

### Application Restart

1. Identify the affected service from on-call dashboard.
2. Trigger a rolling restart via the service's runbook.
3. Monitor error rates for 30 minutes post-restart.

## Contact List

- **Incident Commander on-call:** see PagerDuty schedule `cronos-incident`.
- **Database on-call:** see PagerDuty schedule `cronos-db`.
- **Security on-call:** see PagerDuty schedule `cronos-security`.

## Dependencies

- Cloud provider status page
- Internal DNS automation
- PagerDuty
- Status page at `status.pantalasa.org`

## Review Cadence

This plan is reviewed quarterly. The last review was conducted on 2026-02-10 by `brandon@pantalasa.org`.
