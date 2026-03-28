# Build & Release

## Commands

```bash
make build          # Build to ./build/{{PROJECT_NAME}}
make test           # Run all tests
make install        # Install to ~/.local/bin/
make lint           # golangci-lint
make clean          # Remove build artifacts
```

## Release Pipeline

- Conventional Commits → go-semantic-release → GoReleaser
- `feat:` → minor, `fix:` → patch, `BREAKING CHANGE:` → major
- Homebrew: `brew install {{GITHUB_OWNER}}/cli/{{PROJECT_NAME}}`
- Binaries: linux/darwin/windows × amd64/arm64

## Demo GIF

터미널 데모는 asciinema cast 파일을 작성 후 agg로 GIF 변환합니다.

### cast 파일 형식 (asciinema v2)

`docs/demo.cast`:
```json
{"version": 2, "width": 100, "height": 30, "timestamp": 1711612800, "env": {"SHELL": "/bin/zsh", "TERM": "xterm-256color"}}
[시간(초), "o", "출력 텍스트\r\n"]
```

- 첫 줄: 헤더 (width/height 필수)
- 이후: `[timestamp_seconds, "o", "output_text"]` 배열
- `\r\n`: 줄바꿈, `\u001b[...m`: ANSI 색상 코드
- 타이밍: 이전 이벤트 대비 0.05s 간격이면 즉시 출력, 1~2s 간격이면 사람이 타이핑하는 느낌

### 권장 데모 구성

1. 상단 헤더: 프로젝트명 + 설명 + GitHub URL
2. `--help` 출력
3. 기본 검색 (테이블)
4. 영문 출력 (`--lang en`)
5. JSON 출력 (`-o json`)
6. CSV 내보내기 (`-o csv`)
7. 캐시 관리 (`cache stats`)

### ANSI 색상 참고

| 코드 | 색상 |
|------|------|
| `\u001b[1;36m...\u001b[0m` | Bold Cyan (섹션 제목) |
| `\u001b[1;33m...\u001b[0m` | Bold Yellow (코멘트) |
| `\u001b[32m...\u001b[0m` | Green (프롬프트 ❯) |
| `\u001b[1m...\u001b[0m` | Bold (테이블 헤더) |
| `\u001b[2m...\u001b[0m` | Dim (설치 안내) |

### 변환 명령

```bash
agg docs/demo.cast docs/demo.gif --font-size 16 --cols 100 --rows 30
```

### 적용 위치

- `README.md`: `![demo](docs/demo.gif)`
- `docs/ko.html`, `docs/en.html`: `<img src="demo.gif">` (이미 placeholder 있음)
