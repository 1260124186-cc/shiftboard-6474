# Shiftboard

Shiftboard is a local command-line service for small operations teams to track work
orders, assign them to shifts, and produce a compact workload report. It is designed
for supervisors who need predictable planning without a database or external service.

## Project layout

- `cmd/shiftboard`: example executable that creates and reports a work board.
- `internal/model`: work-order and stage domain types.
- `internal/store`: in-memory persistence with snapshot support.
- `internal/plan`: shift assignment and workload calculations.
- `internal/report`: human-readable report formatting.
- `internal/service`: business workflow that coordinates the packages above.

## Commands

```bash
go build ./...
go test ./...
go run ./cmd/shiftboard
```

The sample executable prints a workload report for a few in-memory work orders. The
application has no required environment variables and does not call online services.
