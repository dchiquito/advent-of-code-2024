package day19

import (
	"bufio"
	"fmt"
	"io"
	"sort"
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

var colors = [...]byte{'b', 'g', 'r', 'u', 'w'}

func charIndex(c byte) int {
	var i int
	switch c {
	case 'b':
		i = 0
	case 'g':
		i = 1
	case 'r':
		i = 2
	case 'u':
		i = 3
	case 'w':
		i = 4
	}
	return i
}
func buildIndex(towels []string) []int {
	index := make([]int, 5)
	t := 0
	for i1, c1 := range colors {
		index[i1] = t
		for t < len(towels) && towels[t][0] == c1 {
			t += 1
		}
	}
	return index
}

func recur(towels []string, pattern string, index []int) bool {
	if len(pattern) == 0 {
		return true
	}
	ind := charIndex(pattern[0])
	start := index[ind]
	end := len(towels) - 1
	if ind+1 < len(index) {
		end = index[ind+1]
	}

	for _, towel := range towels[start:end] {
		if len(pattern) >= len(towel) && pattern[:len(towel)] == towel {
			if recur(towels, pattern[len(towel):], index) {
				return true
			}
		}
	}
	return false
}

func Level1(in io.Reader) string {
	towels, patterns := parse(in)
	sort.Strings(towels)
	index := buildIndex(towels)
	total := 0
	for _, pattern := range patterns {
		if recur(towels, pattern, index) {
			total += 1
		}
	}
	return fmt.Sprint(total)
}

func recur2(towels []string, pattern string, cache *map[int]int, index []int) int {
	if len(pattern) == 0 {
		return 1
	}
	if (*cache)[len(pattern)] > 0 {
		return (*cache)[len(pattern)]
	}
	ind := charIndex(pattern[0])
	start := index[ind]
	end := len(towels) - 1
	if ind+1 < len(index) {
		end = index[ind+1]
	}
	total := 0
	for _, towel := range towels[start:end] {
		if len(pattern) >= len(towel) && pattern[:len(towel)] == towel {
			total += recur2(towels, pattern[len(towel):], cache, index)
		}
	}
	(*cache)[len(pattern)] = total
	return total
}

func Level2(in io.Reader) string {
	towels, patterns := parse(in)
	sort.Strings(towels)
	index := buildIndex(towels)
	total := 0
	for _, pattern := range patterns {
		total += recur2(towels, pattern, &map[int]int{}, index)
	}
	return fmt.Sprint(total)
}
