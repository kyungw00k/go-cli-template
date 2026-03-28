# go-cli-template Design Spec

## Goal

kyungw00k Go CLI 프로젝트의 GitHub Template Repository. 새 CLI를 만들 때 `gh repo create --template`으로 복제 후 `scripts/init.sh`로 플레이스홀더를 치환하면 빌드 가능한 프로젝트가 완성된다.

## Architecture

풀 스캐폴딩 — Cobra CLI, 멀티포맷 출력(table/json/jsonl/csv), i18n(ko/en), SQLite 캐시, 셀프 업데이트, AI tool-schema, GitHub Pages, 릴리스 파이프라인을 모두 포함.

## File Structure

```
go-cli-template/
├── cmd/{{PROJECT_NAME}}/main.go
├── internal/
│   ├── cli/
│   │   ├── root.go
│   │   ├── cache_root.go
│   │   ├── cache_clear.go
│   │   ├── cache_stats.go
│   │   ├── tool_schema.go
│   │   └── update.go
│   ├── i18n/
│   │   ├── i18n.go
│   │   └── messages.go
│   └── output/
│       ├── formatter.go
│       ├── table.go
│       ├── json.go
│       ├── jsonl.go
│       └── csv.go
├── cache/
│   ├── cache.go
│   └── cache_test.go
├── Makefile
├── .goreleaser.yml
├── .github/workflows/
│   ├── release.yml
│   └── pages.yml
├── docs/
│   ├── index.html          # 언어 감지 리다이렉트
│   ├── ko.html              # 한국어 랜딩
│   ├── en.html              # 영문 랜딩
│   ├── install.sh
│   └── skill.md
├── scripts/
│   └── init.sh
├── CLAUDE.md
├── README.md
├── .gitignore
└── go.mod
```

## Placeholders

| Placeholder | Purpose | Example |
|-------------|---------|---------|
| `{{PROJECT_NAME}}` | Binary name, directory | `juso` |
| `{{MODULE_PATH}}` | Go module path | `github.com/kyungw00k/juso` |
| `{{DESCRIPTION}}` | Korean description | `한국 우편번호 검색 CLI` |
| `{{DESCRIPTION_EN}}` | English description | `Korean postal code lookup CLI` |
| `{{GITHUB_OWNER}}` | GitHub owner | `kyungw00k` |
| `{{GITHUB_REPO}}` | GitHub repo name | `juso` |

## init.sh Behavior

1. Parse flags (--name, --module, --description, --description-en)
2. Infer GITHUB_OWNER and GITHUB_REPO from MODULE_PATH
3. sed replace all placeholders in all files
4. Rename cmd/{{PROJECT_NAME}}/ directory
5. go mod edit -module to set module path
6. go mod tidy
7. Remove scripts/ directory
8. make build to verify
9. git add -A && git commit

## CLAUDE.md

LLM이 프로젝트를 즉시 파악할 수 있는 컨텍스트 파일:
- 아키텍처 개요 + 패키지 의존 관계
- 새 커맨드/i18n 키/테이블 컬럼 추가 절차
- 빌드/테스트/릴리스 명령어
- kyungw00k-cli-guide 코딩 컨벤션 요약

## GitHub Pages

upbit 스타일: index.html이 브라우저 언어 감지 후 ko.html/en.html로 리다이렉트. 다크 테마(#0d1117), 배지, 카드 그리드, 데모 GIF 섹션.
