package day19

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func parse(in io.Reader) ([]string, []string) {
	scanner := bufio.NewScanner(in)
	scanner.Scan()
	towels := strings.Split(scanner.Text(), ", ")
	scanner.Scan()
	patterns := []string{}
	for scanner.Scan() {
		patterns = append(patterns, scanner.Text())
	}
	return towels, patterns
}

func recur(towels []string, pattern string) bool {
	if len(pattern) == 0 {
		return true
	}
	for _, towel := range towels {
		if len(pattern) >= len(towel) && pattern[:len(towel)] == towel {
			if recur(towels, pattern[len(towel):]) {
				return true
			}
		}
	}
	return false
}

func Level1(in io.Reader) string {
	towels, patterns := parse(in)
	total := 0
	for _, pattern := range patterns {
		if recur(towels, pattern) {
			total += 1
		}
	}
	return fmt.Sprint(total)
}

func recur2(towels []string, pattern string, cache *map[int]int) int {
	if len(pattern) == 0 {
		return 1
	}
	if (*cache)[len(pattern)] > 0 {
		return (*cache)[len(pattern)]
	}
	total := 0
	for _, towel := range towels {
		if len(pattern) >= len(towel) && pattern[:len(towel)] == towel {
			total += recur2(towels, pattern[len(towel):], cache)
		}
	}
	(*cache)[len(pattern)] = total
	return total
}

func Level2(in io.Reader) string {
	towels, patterns := parse(in)
	total := 0
	for _, pattern := range patterns {
		total += recur2(towels, pattern, &map[int]int{})
	}
	return fmt.Sprint(total)
}
