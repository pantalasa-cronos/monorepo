# Organization Compliance

This directory holds compliance artifacts that apply to the entire organization, not just a single service.

## Contents

- `dr-plan.md` — disaster recovery plan, recovery objectives (RTO/RPO), and procedures
- `dr-exercises/` — historical DR exercise records (tabletops, failover drills, etc.)
- `policies/` — organization-wide policies (access control, incident response)

## How this is consumed

Services in this monorepo reference the shared compliance data via the `release-bundle` collector, which snapshots the relevant slice into each service's Component JSON at collection time. This way, each service carries a point-in-time copy of the organization's compliance posture as part of its release record.

Per-service compliance artifacts (data classification, security review) live under each service's own `compliance/` directory and are merged into the service's release bundle alongside the shared org-level data.
