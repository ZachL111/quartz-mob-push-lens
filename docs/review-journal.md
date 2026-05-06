# Review Journal

The repository goal stays the same: create a Go reference implementation for push workflows, centered on protocol validation, framed sample traffic, and bounds and ordering tests. This note explains the added review angle.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 119, lane `watch`
- `stress`: `sync drift`, score 204, lane `ship`
- `edge`: `local state`, score 144, lane `ship`
- `recovery`: `conflict cost`, score 186, lane `ship`
- `stale`: `form pressure`, score 174, lane `ship`

## Note

The repository should be understandable without pretending it is larger than it is.
