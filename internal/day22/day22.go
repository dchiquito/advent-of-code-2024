package day22

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

func parse(in io.Reader) []int {
	scanner := bufio.NewScanner(in)
	lines := []int{}
	for scanner.Scan() {
		lines = append(lines, util.ToInt(scanner.Text()))
	}
	return lines
}

func nextSecret(prev int) int {
	mod := 16777216
	a := ((prev << 6) ^ prev) % mod
	b := ((a >> 5) ^ a) % mod // TODO remove mod
	c := ((b << 11) ^ b) % mod
	return c
}

func Level1(in io.Reader) string {
	lines := parse(in)
	total := 0
	for _, line := range lines {
		s := line
		for i := 0; i < 2000; i += 1 {
			s = nextSecret(s)
		}
		total += s
	}
	return fmt.Sprint(total)
}

func Level2(in io.Reader) string {
	return ""
}
