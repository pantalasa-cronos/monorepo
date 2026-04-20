---
policy_id: "POL-IR-001"
owner: "security@pantalasa.org"
last_reviewed: "2026-02-01"
effective_date: "2025-08-15"
---

# Incident Response Policy

## Purpose

Describe how Pantalasa Cronos detects, responds to, and learns from production incidents and security events.

## Scope

Applies to all production services, customer-facing systems, and any security event affecting Pantalasa Cronos infrastructure or data.

## Policy

### Detection

Production incidents are detected through:

- Automated alerting via `cronos-observability`
- Customer reports via support channels
- Security tool alerts (IDS, EDR, SIEM)

### Classification

Incidents are classified as:

- **Sev 1** — customer-facing outage, data breach, or regulatory event. 15-minute response target.
- **Sev 2** — partial degradation, non-customer-facing security event. 60-minute response target.
- **Sev 3** — minor issue, no customer impact. Best-effort response.

### Response

1. First responder acknowledges the incident and updates the incident channel.
2. Incident Commander is assigned for Sev 1 and Sev 2 incidents.
3. Communications Lead owns customer-facing and internal updates.
4. Resolution steps are documented in real time in the incident channel.
5. Post-incident review is conducted within 5 business days for Sev 1 and Sev 2 incidents.

### Learning

Post-incident reviews are blameless and produce action items tracked to completion. Patterns across incidents are reported quarterly to engineering leadership.

## Enforcement

Failure to follow this policy may be reviewed by the Security team and engineering leadership.
