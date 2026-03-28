#!/bin/bash
set -euo pipefail

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

usage() {
  echo "Usage: $0 --name <name> --module <module> --description <desc> --description-en <desc-en>"
  echo ""
  echo "Options:"
  echo "  --name            Project/binary name (e.g., juso)"
  echo "  --module          Go module path (e.g., github.com/kyungw00k/juso)"
  echo "  --description     Korean description"
  echo "  --description-en  English description"
  exit 1
}

NAME=""
MODULE=""
DESC=""
DESC_EN=""

while [[ $# -gt 0 ]]; do
  case $1 in
    --name) NAME="$2"; shift 2 ;;
    --module) MODULE="$2"; shift 2 ;;
    --description) DESC="$2"; shift 2 ;;
    --description-en) DESC_EN="$2"; shift 2 ;;
    *) usage ;;
  esac
done

[[ -z "$NAME" || -z "$MODULE" || -z "$DESC" || -z "$DESC_EN" ]] && usage

# Infer GitHub owner and repo from module path
# e.g., github.com/kyungw00k/juso → owner=kyungw00k, repo=juso
OWNER=$(echo "$MODULE" | cut -d'/' -f2)
REPO=$(echo "$MODULE" | cut -d'/' -f3)

echo -e "${GREEN}Initializing project: ${NAME}${NC}"
echo "  Module:      $MODULE"
echo "  Description: $DESC"
echo "  English:     $DESC_EN"
echo "  GitHub:      $OWNER/$REPO"
echo ""

# Replace placeholders in all files (exclude .git/ and binary files)
echo -e "${YELLOW}Replacing placeholders...${NC}"
find . -type f \
  -not -path './.git/*' \
  -not -path './build/*' \
  -not -name '*.gif' \
  -not -name '*.png' \
  -not -name '*.jpg' \
  -not -name 'init.sh' \
  | while read -r file; do
    if file "$file" | grep -q text; then
      sed -i '' \
        -e "s|{{PROJECT_NAME}}|${NAME}|g" \
        -e "s|{{MODULE_PATH}}|${MODULE}|g" \
        -e "s|{{DESCRIPTION}}|${DESC}|g" \
        -e "s|{{DESCRIPTION_EN}}|${DESC_EN}|g" \
        -e "s|{{GITHUB_OWNER}}|${OWNER}|g" \
        -e "s|{{GITHUB_REPO}}|${REPO}|g" \
        "$file" 2>/dev/null || true
    fi
  done

# Rename cmd directory
if [ -d "cmd/{{PROJECT_NAME}}" ]; then
  echo -e "${YELLOW}Renaming cmd/{{PROJECT_NAME}} → cmd/${NAME}${NC}"
  mv "cmd/{{PROJECT_NAME}}" "cmd/${NAME}"
fi

# Rename plugin skill directory if exists
if [ -d ".claude/skills/{{PROJECT_NAME}}" ]; then
  mv ".claude/skills/{{PROJECT_NAME}}" ".claude/skills/${NAME}"
fi

# Set Go module
echo -e "${YELLOW}Setting Go module...${NC}"
cat > go.mod << EOF
module ${MODULE}

go 1.25
EOF

# Install dependencies
echo -e "${YELLOW}Installing dependencies...${NC}"
go get github.com/spf13/cobra
go get github.com/charmbracelet/lipgloss
go get github.com/mattn/go-runewidth
go get golang.org/x/term
go get golang.org/x/text
go get github.com/jmoiron/sqlx
go get modernc.org/sqlite
go mod tidy

# Verify build
echo -e "${YELLOW}Building...${NC}"
make build

# Clean up init script
echo -e "${YELLOW}Cleaning up...${NC}"
rm -rf scripts/
rm -f docs/superpowers/specs/2026-03-28-go-cli-template-design.md
rmdir docs/superpowers/specs 2>/dev/null || true
rmdir docs/superpowers 2>/dev/null || true

echo ""
echo -e "${GREEN}✅ Project initialized: ${NAME}${NC}"
echo ""
echo -e "${YELLOW}⚠️  GitHub Secrets 설정 필요:${NC}"
echo "  gh secret set HOMEBREW_TAP_TOKEN --body \"your-pat-token\""
echo ""
echo -e "${YELLOW}⚠️  GitHub Pages 설정 필요:${NC}"
echo "  Settings → Pages → Source: \"GitHub Actions\""
echo ""
echo -e "${YELLOW}다음 단계:${NC}"
echo "  1. API 클라이언트 수정: api/client.go"
echo "  2. 데이터 모델 수정: api/client.go (Result struct)"
echo "  3. 테이블 컬럼 수정: internal/cli/root.go, internal/output/formatter.go"
echo "  4. i18n 메시지 수정: internal/i18n/messages.go"
echo "  5. tool-schema 수정: internal/cli/tool_schema.go"
echo "  6. GitHub Pages 수정: docs/ko.html, docs/en.html"
echo "  7. 데모 GIF 생성: docs/demo.cast → docs/demo.gif"
