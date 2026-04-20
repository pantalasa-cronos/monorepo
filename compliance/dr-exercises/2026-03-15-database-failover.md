---
date: "2026-03-15"
exercise_type: "failover"
facilitator: "brandon@pantalasa.org"
participants: 6
outcome: "successful"
---

# DR Exercise: Database Failover Drill

## Scenario

Simulated loss of the primary PostgreSQL instance for the `transactions` service at 14:00 UTC. On-call engineer paged via PagerDuty and asked to execute the database failover procedure from `dr-plan.md`.

## Recovery Steps Tested

1. Confirmed primary was unreachable via `cronosctl db ping --primary`.
2. Promoted replica `replica-1` using `cronosctl db failover --target replica-1`.
3. Updated DNS records and waited for propagation (took ~2 minutes).
4. Verified application health with `cronosctl health --service transactions`.

## Participants

- Incident Commander: brandon@pantalasa.org
- DBA on-call: shane@pantalasa.org
- SRE observer: mick@pantalasa.org
- Service engineer: jose@pantalasa.org
- Security observer: fadi@pantalasa.org
- Scribe: alma@pantalasa.org

## Timing

- **Detection:** 14:00 UTC (synthetic trigger)
- **First response:** 14:02 UTC
- **Failover initiated:** 14:08 UTC
- **DNS updated:** 14:10 UTC
- **Recovery confirmed:** 14:17 UTC
- **Total downtime:** ~17 minutes (within 60-minute RTO)

## Action Items

- Automate DNS update step to reduce manual toil (owner: mick@pantalasa.org, due 2026-05-01)
- Add synthetic failover alert to on-call playbook (owner: shane@pantalasa.org, due 2026-04-15)

## Lessons Learned

- The DNS propagation step took longer than expected; worth investing in automation.
- Escalation path worked correctly but the runbook link in the pager alert was stale; fixed in a follow-up PR.
