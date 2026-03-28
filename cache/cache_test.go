package cache

import (
	"os"
	"testing"
)

func setupTestCache(t *testing.T) *Cache {
	t.Helper()
	dir := t.TempDir()
	os.Setenv("XDG_CACHE_HOME", dir)
	t.Cleanup(func() { os.Unsetenv("XDG_CACHE_HOME") })

	c, err := Open()
	if err != nil {
		t.Fatalf("Open() error: %v", err)
	}
	t.Cleanup(func() { c.Close() })
	return c
}

func TestCache_SetAndGet(t *testing.T) {
	c := setupTestCache(t)

	err := c.Set("test", []string{"a", "b"})
	if err != nil {
		t.Fatalf("Set() error: %v", err)
	}

	data, ok := c.Get("test")
	if !ok {
		t.Fatal("Get() returned false")
	}
	if len(data) == 0 {
		t.Fatal("Get() returned empty data")
	}
}

func TestCache_GetMiss(t *testing.T) {
	c := setupTestCache(t)

	_, ok := c.Get("missing")
	if ok {
		t.Fatal("Get() returned true for missing key")
	}
}

func TestCache_NormalizeKey(t *testing.T) {
	c := setupTestCache(t)

	c.Set("Hello", []string{"a"})
	_, ok := c.Get("hello")
	if !ok {
		t.Fatal("Get() should match case-insensitively")
	}
}

func TestCache_ClearAndStats(t *testing.T) {
	c := setupTestCache(t)

	c.Set("k1", []string{"a"})
	c.Set("k2", []string{"b"})

	stats, _ := c.Stats()
	if stats.Entries != 2 {
		t.Errorf("Entries = %d, want 2", stats.Entries)
	}

	c.Clear()

	stats, _ = c.Stats()
	if stats.Entries != 0 {
		t.Errorf("Entries after clear = %d, want 0", stats.Entries)
	}
}

func TestFormatSize(t *testing.T) {
	tests := []struct {
		bytes int64
		want  string
	}{
		{500, "500 B"},
		{1024, "1.0 KB"},
		{1048576, "1.0 MB"},
	}
	for _, tt := range tests {
		got := FormatSize(tt.bytes)
		if got != tt.want {
			t.Errorf("FormatSize(%d) = %q, want %q", tt.bytes, got, tt.want)
		}
	}
}
