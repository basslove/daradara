package util_test

import (
	"github.com/basslove/daradara/internal/api/pkg/util"
	"testing"
)

func TestBuildUniqueStringSlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		l        []string
		expected []string
	}{
		{
			name:     "empty slice",
			l:        []string{},
			expected: []string{},
		},
		{
			name:     "single element",
			l:        []string{"a"},
			expected: []string{"a"},
		},
		{
			name:     "multiple elements",
			l:        []string{"a", "b", "a", "c", "b"},
			expected: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uniqueStrings := util.BuildUniqueStringSlice(tt.l)

			if len(uniqueStrings) != len(tt.expected) {
				t.Errorf("expected %d unique strings, got %d", len(tt.expected), len(uniqueStrings))
			}

			for i, s := range uniqueStrings {
				if s != tt.expected[i] {
					t.Errorf("expected %s at index %d, got %s", tt.expected[i], i, s)
				}
			}
		})
	}
}
