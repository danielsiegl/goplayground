package main

import "testing"

func TestBuildURL(t *testing.T) {
	tests := []struct {
		project string
		branch  string
		want    string
	}{
		{"my-project", "main", "https://chatgpt.com/codex/my-project?branch=main"},
		{"danielsiegl/goplayground", "main", "https://chatgpt.com/codex/danielsiegl/goplayground?branch=main"},
	}
	for _, tt := range tests {
		if got := buildURL(tt.project, tt.branch); got != tt.want {
			t.Errorf("buildURL(%q, %q) = %q, want %q", tt.project, tt.branch, got, tt.want)
		}
	}
}
