# quartz-mob-push-lens

`quartz-mob-push-lens` is a focused Go codebase around create a Go reference implementation for push workflows, centered on protocol validation, framed sample traffic, and bounds and ordering tests. It is meant to be easy to inspect, run, and extend without a hosted service.

## Quartz Mob Push Lens Walkthrough

I would read the project from the outside in: command, fixture, model, then roadmap. That keeps the mobile workflows idea grounded in files that can be checked locally.

## Reason For The Project

This is not a wrapper around a service. It is a self-contained project that shows how the model behaves when demand, capacity, latency, risk, and weight move in different directions.

## Capabilities

- Models local state with deterministic scoring and explicit review decisions.
- Uses fixture data to keep sync pressure changes visible in code review.
- Includes extended examples for form constraints, including `surge` and `degraded`.
- Documents offline paths tradeoffs in `docs/operations.md`.
- Runs locally with a single verification command and no external credentials.

## How It Is Put Together

The interesting part is the boundary between accepted and reviewed scenarios. Extended examples sit near that boundary so future edits can show whether the model became more permissive or more cautious. The Go layout uses small packages and table-oriented tests so the behavior stays easy to follow.

## Where Things Live

- `policy`: Go package with the core model
- `cmd`: small command entry point
- `fixtures`: compact golden scenarios
- `examples`: expanded scenario set
- `metadata`: project constants and verification metadata
- `docs`: operations and extension notes
- `scripts`: local verification and audit commands
- `go.mod`: Go module metadata

## Getting It Running

Clone the repository, enter the directory, and run the verifier. No database server, cloud account, or token is required.

## Command Examples

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

This runs the language-level build or test path against the compact fixture set.

## Check The Work

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/audit.ps1
```

The audit command checks repository structure and README constraints before it delegates to the verifier.

## Data Notes

`baseline` is the first example I would inspect because it lands on the `review` path with a score of 51. The broader file also keeps `degraded` at -116 and `surge` at 151, which gives the model a useful low-to-high spread.

## Tradeoffs

This code is local-first. It makes no claim about deployed usage and avoids credentials, hosted state, and environment-specific setup.

## Possible Extensions

- Add a loader for `examples/extended_cases.csv` and promote selected cases into the language test suite.
- Add a short report command that prints the score breakdown for a single scenario.
- Add malformed input fixtures so the failure path is as visible as the happy path.
- Add one more mobile workflows fixture that focuses on a malformed or borderline input.
