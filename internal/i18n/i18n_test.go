package i18n

import "testing"

func TestT_ReturnsMessage(t *testing.T) {
	current = en
	got := T(MsgCacheCleared)
	want := "Cache cleared."
	if got != want {
		t.Errorf("T() = %q, want %q", got, want)
	}
}

func TestT_UnknownKey(t *testing.T) {
	current = en
	got := T("UnknownKey")
	if got != "UnknownKey" {
		t.Errorf("T(unknown) = %q, want key back", got)
	}
}

func TestTf_WithArgs(t *testing.T) {
	current = en
	got := Tf(MsgCacheEntries, 42)
	want := "Cache entries: 42"
	if got != want {
		t.Errorf("Tf() = %q, want %q", got, want)
	}
}
