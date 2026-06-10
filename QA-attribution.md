# Monorepo CI → component attribution QA

This repo is a test bed for Lunar's CI-agent attribution: how a CI job's
collected facts get pinned to the right monorepo sub-component
(`ci/observe.go:selectComponents`). It exercises all **five** attribution
methods, one job each, and asserts the result with the `qa-attribution` policy
in `pantalasa-cronos/lunar`.

## The five methods

| # | Method | Where | Resolves to |
|---|--------|-------|-------------|
| 1 | explicit single | `ci.yml` job `explicit-single` (`LUNAR_COMPONENT=services/backend-go`) | `services/backend-go` |
| 2 | explicit comma-multi | `ci.yml` job `explicit-multi` (`LUNAR_COMPONENT="services/backend-go,compliance"`) | `services/backend-go` + `compliance` |
| 3 | ciPipelines | `compliance.yml` (workflow name `compliance-ci` ↔ component `ciPipelines`) | `compliance` |
| 4 | INFER changed-paths | `ci.yml` job `infer-changed-paths` (`INFER=true`, repo-root cwd) | the subdir the commit **touched** |
| 5 | INFER cwd | `ci.yml` job `infer-cwd` (`INFER=true`, `working-directory: services/api-python`) | `services/api-python` |

Each job emits a unique marker command `echo "lunar-qa::<method>"`. The `ci-meta`
collector records every traced command under the resolved component's
`.ci.commands.*`, so the marker lands on whichever component attribution chose.

## Why two probe commits

In `selectComponents`, **changed-paths is global and tried before cwd**. So a
job's cwd inference only fires when the commit touches **no** component path.
That makes #4 and #5 mutually exclusive on a single push:

- **Probe A — touch a subdir** (e.g. edit `services/api-python/probe.txt`):
  exercises #1, #2, #3, and **#4** (changed-paths → `api-python`). The
  `infer-cwd` job also resolves `api-python` here, but via changed-paths, not
  cwd — so this probe does **not** cleanly test #5.
- **Probe B — touch only a repo-root file** (e.g. edit `./qa-probe.txt`):
  changed-paths is empty, so **#5** resolves `api-python` via **cwd**. The
  `infer-changed-paths` job (repo-root) falls to the whole-repo fallback
  (the top-level `…/monorepo` component) — documenting that path too.

Run both probes to cover all five cleanly.

## Expected attribution (asserted by the policy)

Markers that fire on **every** push: `explicit-single` → backend-go;
`explicit-multi` → backend-go + compliance; `ci-pipelines` → compliance.

Isolation invariants (must hold on every push):

| Component | Must have | Must NOT have |
|---|---|---|
| `services/backend-go` | explicit-single, explicit-multi | ci-pipelines, infer-cwd |
| `compliance` | ci-pipelines, explicit-multi | explicit-single, infer-cwd |
| `services/api-python` | — (markers only via probes) | explicit-single, explicit-multi, ci-pipelines |

`infer-changed-paths` / `infer-cwd` markers are probe-dependent and land on
`api-python` (or the top-level component for changed-paths on probe B); the
policy treats them as allowed-where-expected, never as cross-attribution.
