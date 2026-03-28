# Output & Cache

## Output System

| Format | Flag | Behavior |
|--------|------|----------|
| auto | `-o auto` (default) | TTY → table, pipe → json |
| table | `-o table` | lipgloss styled, CJK-aware columns |
| json | `-o json` | TTY → pretty, pipe → compact |
| jsonl | `-o jsonl` | One JSON object per line |
| csv | `-o csv` | Full fields (AllColumns), nil → empty string |

## Cache

- Engine: modernc.org/sqlite (pure Go, no CGO)
- TTL: 24 hours
- Key: NFC-normalized + lowercased search keyword
- Path: $XDG_CACHE_HOME/{{PROJECT_NAME}}/ → ~/.cache/{{PROJECT_NAME}}/ → ~/.{{PROJECT_NAME}}/

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 3 | Rate limit (HTTP 429) |
| 4 | Validation error (HTTP 4xx) |
