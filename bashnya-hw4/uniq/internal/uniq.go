package uniq

import (
	"fmt"
	"strings"
)

type Options struct {
	Count      bool // -c
	Duplicates bool // -d
	UniqueOnly bool // -u
	IgnoreCase bool // -i
	SkipFields int  // -f num
	SkipChars  int  // -s num
}

func normalize(line string, opts Options) string {
	if opts.IgnoreCase {
		line = strings.ToLower(line)
	}

	// Пропуск первых num_fields
	fields := strings.Fields(line)
	if opts.SkipFields > 0 && len(fields) > opts.SkipFields {
		line = strings.Join(fields[opts.SkipFields:], " ")
	}

	// Пропуск первых num_chars
	if opts.SkipChars > 0 && len(line) > opts.SkipChars {
		line = line[opts.SkipChars:]
	}

	return line
}

// возвращает уникальные строки в зависимости от опций
func Uniq(lines []string, opts Options) []string {
	counts := make(map[string]int)
	order := []string{}

	for _, line := range lines {
		key := normalize(line, opts)
		if _, exists := counts[key]; !exists {
			order = append(order, line)
		}
		counts[key]++
	}

	var result []string
	for _, line := range order {
		key := normalize(line, opts)
		cnt := counts[key]

		switch {
		case opts.Duplicates && cnt > 1:
			result = append(result, line)
		case opts.UniqueOnly && cnt == 1:
			result = append(result, line)
		case !opts.Duplicates && !opts.UniqueOnly:
			if opts.Count {
				result = append(result, fmt.Sprintf("%d %s", cnt, line))
			} else {
				result = append(result, line)
			}
		}
	}
	return result
}
