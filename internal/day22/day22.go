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
	b := (a >> 5) ^ a
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
	lines := parse(in)
	maxDelta := ((10 + 9) * 20 * 20 * 20) + ((10) * 20 * 20) + ((10) * 20) + (10)
	sums := make([]int, maxDelta+1)
	lastVisit := make([]int, maxDelta+1)
	for i, line := range lines {
		lastPrice := line % 10
		lastDelta := 0
		secret := nextSecret(line)
		for j := 1; j < 2000; j += 1 {
			secret = nextSecret(secret)
			price := secret % 10
			delta := 10 + price - lastPrice
			delta = ((lastDelta * 20) % (20 * 20 * 20 * 20)) + delta
			if lastVisit[delta] != i {
				lastVisit[delta] = i
				sums[delta] += price
			}
			lastPrice = price
			lastDelta = delta
		}
	}

	bestTotal := 0
	for _, total := range sums {
		if total > bestTotal {
			bestTotal = total
		}
	}

	// 80 seconds, more than twice as fast
	return fmt.Sprint(bestTotal)
}
