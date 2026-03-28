# go-cli-template

kyungw00k Go CLI 프로젝트 템플릿.

Cobra CLI, 멀티포맷 출력(table/json/jsonl/csv), i18n(ko/en), SQLite 캐시, 셀프 업데이트, AI tool-schema, GitHub Pages, 릴리스 파이프라인을 포함합니다.

## 사용법

### 1. 템플릿에서 새 리포 생성

```bash
gh repo create myproject --template kyungw00k/go-cli-template --public
cd myproject
```

### 2. 프로젝트 초기화

```bash
./scripts/init.sh \
  --name myproject \
  --module github.com/kyungw00k/myproject \
  --description "내 프로젝트 설명" \
  --description-en "My project description"
```

### 3. 커스터마이징

| 파일 | 수정 내용 |
|------|----------|
| `api/client.go` | API URL, 요청/응답 구조체 |
| `internal/cli/root.go` | 테이블 컬럼, 검색 로직 |
| `internal/output/formatter.go` | CSV AllColumns |
| `internal/i18n/messages.go` | 메시지 추가/수정 |
| `internal/cli/tool_schema.go` | AI Schema 수정 |
| `docs/ko.html`, `docs/en.html` | GitHub Pages 내용 |

### 4. GitHub 설정

```bash
# Homebrew tap 토큰 (kyungw00k/homebrew-cli 리포 push용)
gh secret set HOMEBREW_TAP_TOKEN --body "your-pat-token"

# GitHub Pages
# Settings → Pages → Source: "GitHub Actions"
```

## 포함된 기능

- **Cobra CLI** — 커맨드 그룹, 서브커맨드, 글로벌 플래그
- **멀티포맷 출력** — table (lipgloss + CJK), json, jsonl, csv
- **i18n** — POSIX 로케일 감지, ko/en 번역
- **SQLite 캐시** — 24시간 TTL, XDG 경로 지원
- **셀프 업데이트** — GitHub Release 기반 + rollback
- **AI 연동** — tool-schema (JSON Schema), Claude Code 플러그인/스킬
- **릴리스** — go-semantic-release + GoReleaser + Homebrew
- **GitHub Pages** — 다국어 랜딩 (ko/en), 다크 테마

## 프로젝트 구조

```
cmd/{{PROJECT_NAME}}/main.go     # 엔트리포인트
api/                             # HTTP API 클라이언트 (public)
cache/                           # SQLite 캐시 (public)
internal/cli/                    # Cobra 커맨드
internal/i18n/                   # 다국어 메시지
internal/output/                 # 출력 포매터
.claude/skills/{{PROJECT_NAME}}/ # Claude Code 스킬
docs/                            # GitHub Pages + install.sh
scripts/init.sh                  # 프로젝트 초기화 스크립트
```

## 이 템플릿으로 만든 프로젝트

| 프로젝트 | 설명 |
|---------|------|
| [juso](https://github.com/kyungw00k/juso) | 한국 우편번호 검색 CLI |
| [upbit](https://github.com/kyungw00k/upbit) | Upbit 거래소 CLI |
