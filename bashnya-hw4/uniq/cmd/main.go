package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"uniq/internal"
)

func main() {
	// инициализируем флаги
	countFlag := flag.Bool("c", false, "count occurrences of lines")
	dupFlag := flag.Bool("d", false, "show only duplicate lines")
	uniqFlag := flag.Bool("u", false, "show only unique lines")
	ignoreCase := flag.Bool("i", false, "ignore case")
	skipFields := flag.Int("f", 0, "skip num fields")
	skipChars := flag.Int("s", 0, "skip num chars")

	flag.Parse()

	// проверка корректности флагов
	if (*dupFlag && *uniqFlag) || (*dupFlag && *countFlag) || (*uniqFlag && *countFlag) {
		fmt.Fprintln(os.Stderr, "Error: -c, -d, and -u are mutually exclusive")
		flag.Usage()
		os.Exit(1)
	}

	// определяем файлы ввода и вывода
	args := flag.Args()
	var input io.Reader = os.Stdin
	var output io.Writer = os.Stdout

	if len(args) >= 1 {
		file, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening input file:", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	}

	if len(args) >= 2 {
		file, err := os.Create(args[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error creating output file:", err)
			os.Exit(1)
		}
		defer file.Close()
		output = file
	}

	// читаем входные данные
	scanner := bufio.NewScanner(input)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		os.Exit(1)
	}

	opts := uniq.Options{
		Count:      *countFlag,
		Duplicates: *dupFlag,
		UniqueOnly: *uniqFlag,
		IgnoreCase: *ignoreCase,
		SkipFields: *skipFields,
		SkipChars:  *skipChars,
	}

	result := uniq.Uniq(lines, opts)

	for _, line := range result {
		fmt.Fprintln(output, strings.TrimRight(line, "\n"))
	}
}
