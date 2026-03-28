---
name: {{PROJECT_NAME}}
description: "{{DESCRIPTION_EN}}. Search, cache management, and AI tool-schema export."
---

# {{PROJECT_NAME}}

{{DESCRIPTION}}

## CLI Usage

### Search (default command)

`{{PROJECT_NAME}} <keyword>` — Search by keyword.

| Flag | Default | Description |
|------|---------|-------------|
| `-o, --output` | `auto` | Output: auto, table, json, jsonl, csv |

### Cache Management

- `{{PROJECT_NAME}} cache stats` — Show cache statistics (entries, size)
- `{{PROJECT_NAME}} cache clear` — Clear all cached data

### AI Integration

- `{{PROJECT_NAME}} tool-schema` — Export JSON Schema for all commands
- `{{PROJECT_NAME}} tool-schema search` — Export search command schema only

### Self-Update

- `{{PROJECT_NAME}} update` — Update to latest version
- `{{PROJECT_NAME}} update --check` — Check for updates only

## Library Usage (Go)

```go
import "{{MODULE_PATH}}/api"

client := api.NewClient()
results, err := client.Search(ctx, "keyword")
```

## Exit Codes

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | General error |
| 3 | Rate limit |
| 4 | Validation error |
