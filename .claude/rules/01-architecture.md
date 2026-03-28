# Architecture

> Built from [kyungw00k/go-cli-template](https://github.com/kyungw00k/go-cli-template).

## Package Layout

```
cmd/{{PROJECT_NAME}}/main.go    → Entrypoint (cli.Execute + exit codes)
internal/cli/                   → Cobra commands (CLI-only, not importable)
internal/i18n/                  → ko/en message translations (CLI-only)
internal/output/                → Multi-format output: table/json/jsonl/csv (CLI-only)
api/                            → HTTP API client (public, importable)
cache/                          → SQLite cache with TTL (public, importable)
```

## Dependency Flow

```
cmd/{{PROJECT_NAME}} → internal/cli → api/
                                    → cache/
                                    → internal/i18n
                                    → internal/output
```

- `api/` and `cache/` are public — external Go projects can import them
- `internal/` packages are CLI-only — not importable by others
