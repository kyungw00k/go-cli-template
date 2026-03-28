---
name: {{PROJECT_NAME}}
description: {{DESCRIPTION_EN}}
version: 1.0.0
---

# {{PROJECT_NAME}}

{{DESCRIPTION}}

## Commands

### Search (default)

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| keyword | string | Yes | - | Search keyword |
| -o | string | No | auto | Output: auto, table, json, jsonl, csv |

Examples:
- `{{PROJECT_NAME}} hello`
- `{{PROJECT_NAME}} hello -o json`

### Cache

- `{{PROJECT_NAME}} cache stats` — Show cache statistics
- `{{PROJECT_NAME}} cache clear` — Clear all cached data

### JSON Schema

- `{{PROJECT_NAME}} tool-schema` — Export all command schemas
- `{{PROJECT_NAME}} tool-schema search` — Export search schema only
