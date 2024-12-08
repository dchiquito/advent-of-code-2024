package day08

import (
	"bufio"
	"fmt"
	"io"
)

const SIZE = 50

func parse(in io.Reader) []string {
	scanner := bufio.NewScanner(in)
	lines := make([]string, 0, SIZE)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func Level1(in io.Reader) string {
	grid := parse(in)
	var antinodes [SIZE][SIZE]bool
	for y1, line := range grid {
		for x1 := range line {
			c := line[x1]
			if c == '.' {
				continue
			}
			for y2 := (y1 + 1) / 2; y2 < y1+((SIZE+1-y1)/2); y2 += 1 {
				for x2 := (x1 + 1) / 2; x2 < x1+((SIZE+1-x1)/2); x2 += 1 {
					if grid[y2][x2] == c && !(y1 == y2 && x1 == x2) {
						antinodes[y2+y2-y1][x2+x2-x1] = true
					}
				}
			}
			// extra work for no reason, might as well benchmark it
			// for y2 := 0; y2 < SIZE; y2 += 1 {
			// 	for x2 := 0; x2 < SIZE; x2 += 1 {
			// 		if grid[y2][x2] == c && !(y1 == y2 && x1 == x2) {
			// 			ay := y2 + y2 - y1
			// 			ax := x2 + x2 - x1
			// 			if ay >= 0 && ay < SIZE && ax >= 0 && ax < SIZE {
			// 				antinodes[ay][ax] = true
			// 			}
			// 		}
			// 	}
			// }

			// For debugging
			// for yy, line := range antinodes {
			// 	for xx, an := range line {
			// 		if an {
			// 			if yy == y1 && xx == x1 {
			// 				fmt.Print("O")
			// 			} else {
			// 				fmt.Print("X")
			// 			}
			// 		} else {
			// 			fmt.Print(".")
			// 		}
			// 	}
			// 	fmt.Println()
			// }
			// return ""
		}
	}
	total := 0
	for _, line := range antinodes {
		for _, an := range line {
			if an {
				total += 1
			}
			// if an {
			// 	fmt.Print("X")
			// } else {
			// 	fmt.Print(".")
			// }
		}
		// fmt.Println()
	}
	return fmt.Sprint(total)
}

func Level2(in io.Reader) string {
	return ""
}
