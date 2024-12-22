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

	maxDelta := ((10 + 9) * 20 * 20 * 20) + ((10) * 20 * 20) + ((10) * 20) + (10)
	fmt.Println(maxDelta)
	checked := make([]bool, maxDelta+1)

	bestTotal := 0
	for xx, iDeltas := range deltas {
		fmt.Println(xx)
		for _, expectedDelta := range iDeltas {
			// if expectedDelta > maxDelta {
			// 	d := expectedDelta
			// 	fmt.Println(d, d/(20*20*20), (d/(20*20))%20, (d/20)%20, d%20)
			// }
			if checked[expectedDelta] {
				continue
			}
			checked[expectedDelta] = true
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

	xxx := 0
	for _, d := range checked {
		if d {
			xxx += 1
		}
	}
	fmt.Println("Checked", xxx)
	// 80 seconds, more than twice as fast
	return fmt.Sprint(bestTotal)
}
