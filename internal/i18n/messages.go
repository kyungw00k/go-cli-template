package i18n

const (
	// Root
	MsgRootShort Key = "MsgRootShort"
	MsgRootLong  Key = "MsgRootLong"

	// Flags
	FlagOutputUsage Key = "FlagOutputUsage"

	// Groups
	GroupCache Key = "GroupCache"
	GroupUtil  Key = "GroupUtil"

	// Table headers — customize per project
	HdrID   Key = "HdrID"
	HdrName Key = "HdrName"

	// Cache
	MsgCacheShort   Key = "MsgCacheShort"
	MsgCacheClear   Key = "MsgCacheClear"
	MsgCacheStats   Key = "MsgCacheStats"
	MsgCacheCleared Key = "MsgCacheCleared"
	MsgCacheEntries Key = "MsgCacheEntries"
	MsgCacheSize    Key = "MsgCacheSize"

	// Errors
	ErrNoKeyword Key = "ErrNoKeyword"
	ErrNoResults Key = "ErrNoResults"

	// Tool Schema
	MsgToolSchemaShort Key = "MsgToolSchemaShort"

	// Update
	MsgUpdateShort Key = "MsgUpdateShort"
)

var ko = map[Key]string{
	MsgRootShort:       "{{DESCRIPTION}}",
	MsgRootLong:        "{{DESCRIPTION}}",
	FlagOutputUsage:    "출력 형식: auto, table, json, jsonl, csv",
	GroupCache:         "캐시:",
	GroupUtil:          "유틸리티:",
	HdrID:              "ID",
	HdrName:            "이름",
	MsgCacheShort:      "캐시 관리",
	MsgCacheClear:      "캐시 전체 삭제",
	MsgCacheStats:      "캐시 통계",
	MsgCacheCleared:    "캐시가 삭제되었습니다.",
	MsgCacheEntries:    "캐시 항목: %d건",
	MsgCacheSize:       "캐시 크기: %s",
	ErrNoKeyword:       "검색어를 입력하세요.",
	ErrNoResults:       "검색 결과가 없습니다.",
	MsgToolSchemaShort: "AI Agent용 JSON Schema 출력",
	MsgUpdateShort:     "최신 버전으로 업데이트",
}

var en = map[Key]string{
	MsgRootShort:       "{{DESCRIPTION_EN}}",
	MsgRootLong:        "{{DESCRIPTION_EN}}",
	FlagOutputUsage:    "Output format: auto, table, json, jsonl, csv",
	GroupCache:         "Cache:",
	GroupUtil:          "Utility:",
	HdrID:              "ID",
	HdrName:            "NAME",
	MsgCacheShort:      "Manage cache",
	MsgCacheClear:      "Clear all cached data",
	MsgCacheStats:      "Show cache statistics",
	MsgCacheCleared:    "Cache cleared.",
	MsgCacheEntries:    "Cache entries: %d",
	MsgCacheSize:       "Cache size: %s",
	ErrNoKeyword:       "Please enter a search keyword.",
	ErrNoResults:       "No results found.",
	MsgToolSchemaShort: "Export JSON Schema for AI agents",
	MsgUpdateShort:     "Update to the latest version",
}
