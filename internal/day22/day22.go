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
	prices := make([][]int, len(lines))
	deltas := make([][]int, len(lines))
	for i, line := range lines {
		prices[i] = make([]int, 2000)
		deltas[i] = make([]int, 2000)
		prices[i][0] = line % 10
		lastPrice := prices[i][0]
		lastDelta := 0
		secret := nextSecret(line)
		for j := 1; j < 2000; j += 1 {
			secret = nextSecret(secret)
			price := secret % 10
			prices[i][j] = price
			delta := 10 + price - lastPrice
			delta = ((lastDelta * 20) % (20 * 20 * 20 * 20)) + delta
			deltas[i][j] = delta
			lastPrice = price
			lastDelta = delta
		}
	}
	bestTotal := 0
	// 192 seconds :(
	for a := 0; a < 10; a += 1 {
		for b := 0; b < 10; b += 1 {
			fmt.Println(a, b, bestTotal)
			for c := 0; c < 10; c += 1 {
				for d := 0; d < 10; d += 1 {
					for e := 0; e < 10; e += 1 {
						expectedDelta := ((10 + b - a) * 20 * 20 * 20) + ((10 + c - b) * 20 * 20) + ((10 + d - c) * 20) + (10 + e - d)
						total := 0
						for i, iDeltas := range deltas {
							for j, realDelta := range iDeltas {
								if realDelta == expectedDelta {
									total += prices[i][j]
									break
								}
							}
						}
						if total > bestTotal {
							bestTotal = total
						}
					}
				}
			}
		}
	}
	return fmt.Sprint(bestTotal)
}
