package util_test

import (
	"github.com/basslove/daradara/internal/api/pkg/util"
	"testing"
)

func TestHashFromString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: "e3b0c44298fc",
		},
		{
			name:     "hello world",
			input:    "hello world",
			expected: "2ef7bde608ce",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := util.HashFromString(tt.input)
			if got != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, got)
			}
		})
	}
}
