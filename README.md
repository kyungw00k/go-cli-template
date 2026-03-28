# go-cli-template

kyungw00k Go CLI 프로젝트 템플릿.

Cobra CLI, 멀티포맷 출력(table/json/jsonl/csv), i18n(ko/en), SQLite 캐시, 셀프 업데이트, AI tool-schema, GitHub Pages, 릴리스 파이프라인을 포함합니다.

<!-- 초기화 후 아래 주석을 해제하고 데모 GIF를 추가하세요
![demo](docs/demo.gif)
-->

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
| `docs/demo.cast` → `docs/demo.gif` | 터미널 데모 (아래 가이드 참고) |

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

## 데모 GIF 만들기

[asciinema](https://asciinema.org/) cast 파일을 작성하고 [agg](https://github.com/asciinema/agg)로 GIF 변환합니다.

### 1. cast 파일 작성

`docs/demo.cast` — asciinema v2 형식으로 직접 작성하거나 녹화:

```json
{"version": 2, "width": 100, "height": 30, "timestamp": 1711612800, "env": {"SHELL": "/bin/zsh", "TERM": "xterm-256color"}}
[0.5, "o", "\u001b[1;36myourcli\u001b[0m — 프로젝트 설명\r\n"]
[1.0, "o", "\r\n"]
[1.5, "o", "\u001b[32m❯\u001b[0m yourcli search keyword\r\n"]
[2.3, "o", "결과 출력...\r\n"]
```

- 녹화: `asciinema rec docs/demo.cast --cols 100 --rows 30`
- 직접 작성: `[시간(초), "o", "출력 텍스트"]` 형식

### 2. GIF 변환

```bash
# agg 설치: cargo install agg
agg docs/demo.cast docs/demo.gif --font-size 16 --cols 100 --rows 30
```

### 3. README에 추가

```markdown
![demo](docs/demo.gif)
```

GitHub Pages의 `ko.html`/`en.html`에는 이미 `<img src="demo.gif">` placeholder가 포함되어 있습니다.

## 이 템플릿으로 만든 프로젝트

| 프로젝트 | 설명 |
|---------|------|
| [juso](https://github.com/kyungw00k/juso) | 한국 우편번호 검색 CLI |
| [upbit](https://github.com/kyungw00k/upbit) | Upbit 거래소 CLI |
