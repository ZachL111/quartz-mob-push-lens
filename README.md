# quartz-mob-push-lens

`quartz-mob-push-lens` explores mobile workflows with a small Go codebase and local fixtures. The technical goal is to create a Go reference implementation for push workflows, centered on protocol validation, framed sample traffic, and bounds and ordering tests.

## Reason For The Project

The point is to make a small domain rule concrete enough that a reader can change it and immediately see what broke.

## Quartz Mob Push Lens Review Notes

The first comparison I would make is `sync drift` against `form pressure` because it shows where the rule is most opinionated.

## What It Does

- `fixtures/domain_review.csv` adds cases for form pressure and sync drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/quartz-mob-push-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `sync drift` and `form pressure`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## How It Is Put Together

The core code exposes a scoring path and the added review layer uses `signal`, `slack`, `drag`, and `confidence`. The domain terms are `form pressure`, `sync drift`, `local state`, and `conflict cost`.

The added Go path is deliberately direct, with fixtures doing most of the explaining.

## Run It

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Check It

The verifier is intentionally local. It should fail if the fixture score math, lane assignment, or language-specific test drifts.

## Boundaries

This remains a local project with deterministic fixtures. It does not depend on credentials, hosted services, or live data. Future work should add richer malformed inputs before widening the public API.
