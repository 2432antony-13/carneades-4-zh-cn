package main

import (
	"path/filepath"
	"testing"
)

func TestTemplatesDirForLocale(t *testing.T) {
	base := filepath.Join("src", "web", "templates")

	if got := templatesDirForLocale(base, "en"); got != base {
		t.Fatalf("English templates directory = %q, want %q", got, base)
	}

	wantChinese := filepath.Join(base, "zh-CN")
	if got := templatesDirForLocale(base, "zh-CN"); got != wantChinese {
		t.Fatalf("Chinese templates directory = %q, want %q", got, wantChinese)
	}
}
