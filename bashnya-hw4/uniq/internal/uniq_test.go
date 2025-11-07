package uniq

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUniq(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		opts   Options
		expect []string
	}{
		{
			name:   "simple unique",
			input:  []string{"a", "b", "a"},
			opts:   Options{},
			expect: []string{"a", "b"},
		},
		{
			name:   "count option",
			input:  []string{"a", "b", "a"},
			opts:   Options{Count: true},
			expect: []string{"2 a", "1 b"},
		},
		{
			name:   "duplicates only",
			input:  []string{"a", "b", "a"},
			opts:   Options{Duplicates: true},
			expect: []string{"a"},
		},
		{
			name:   "unique only",
			input:  []string{"a", "b", "a"},
			opts:   Options{UniqueOnly: true},
			expect: []string{"b"},
		},
		{
			name:   "ignore case",
			input:  []string{"Hello", "hello"},
			opts:   Options{IgnoreCase: true, Count: true},
			expect: []string{"2 Hello"},
		},
		{
			name:   "skip fields",
			input:  []string{"1 abc def", "2 abc def"},
			opts:   Options{SkipFields: 1, Count: true},
			expect: []string{"2 1 abc def"},
		},
		{
			name:   "skip chars",
			input:  []string{"xxabc", "yyabc"},
			opts:   Options{SkipChars: 2, Count: true},
			expect: []string{"2 xxabc"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Uniq(tt.input, tt.opts)
			require.Equal(t, tt.expect, result)
		})
	}
}
