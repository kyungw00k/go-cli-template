package i18n

import (
	"fmt"
	"os"
	"strings"
)

type Key string

var current map[Key]string
var isKo bool

func init() {
	lang := os.Getenv("LC_ALL")
	if lang == "" {
		lang = os.Getenv("LC_MESSAGES")
	}
	if lang == "" {
		lang = os.Getenv("LANG")
	}
	if strings.HasPrefix(lang, "ko") {
		current = ko
		isKo = true
	} else {
		current = en
		isKo = false
	}
}

func T(key Key) string {
	if msg, ok := current[key]; ok {
		return msg
	}
	return string(key)
}

func Tf(key Key, args ...interface{}) string {
	return fmt.Sprintf(T(key), args...)
}

func IsKorean() bool {
	return isKo
}
