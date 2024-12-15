package day14

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

type Bot struct {
	px int
	py int
	vx int
	vy int
}

func parse(in io.Reader) []Bot {
	scanner := bufio.NewScanner(in)
	bots := make([]Bot, 0, 500)
	for scanner.Scan() {
		line := scanner.Bytes()
		i := 0
		i, px := util.ChompInt(line, i)
		i, py := util.ChompInt(line, i)
		i, vx := util.ChompInt(line, i)
		i, vy := util.ChompInt(line, i)
		bots = append(bots, Bot{px, py, vx, vy})
	}
	return bots
}

const w = 101
const h = 103

func Level1(in io.Reader) string {
	bots := parse(in)
	fmt.Println(bots)
	q1 := 0
	q2 := 0
	q3 := 0
	q4 := 0
	t := 100
	for _, bot := range bots {
		x := (((bot.px + (t * bot.vx)) % w) + w) % w
		y := (((bot.py + (t * bot.vy)) % h) + h) % h
		fmt.Println(x, y)
		if x < w/2 && y < h/2 {
			q1 += 1
		}
		if x > w/2 && y < h/2 {
			q2 += 1
		}
		if x < w/2 && y > h/2 {
			q3 += 1
		}
		if x > w/2 && y > h/2 {
			q4 += 1
		}
	}
	safetyFactor := q1 * q2 * q3 * q4
	return fmt.Sprint(safetyFactor)
}

func floodFill(botGrid []int, visited *[]bool, x int, y int) int {
	i := y*w + x
	if x < 0 || x >= w || y < 0 || y >= h || (*visited)[i] || botGrid[i] > 0 {
		return 0
	}
	(*visited)[i] = true
	return 1 + floodFill(botGrid, visited, x-1, y) + floodFill(botGrid, visited, x+1, y) + floodFill(botGrid, visited, x, y-1) + floodFill(botGrid, visited, x, y+1)
}

func hasCavity(botGrid []int) bool {
	// Start at 1,1 to avoid getting an abnormally small area by being boxed in a corner
	j := w + 2
	for botGrid[j] > 0 {
		j += 1
	}
	x := j % w
	y := j / w
	visited := make([]bool, w*h)
	outerArea := floodFill(botGrid[:], &visited, x, y)
	return outerArea > 100 && outerArea < 9900
}

// This iterates over every cell in the grid, O(10000)
func hasStreak(botGrid []int) bool {
	streak := 0
	for i := 0; i < w*h; i += 1 {
		if i%w == 0 || botGrid[i] == 0 {
			streak = 0
		} else {
			streak += 1
		}
		if streak > 10 {
			return true
		}
	}
	return false
}

// This starts at each bots coordinate to look for a streak, O(500*10) worst case
func hasStreak2(bots []Bot, botGrid []int) bool {
	requiredStreak := 10
	for _, bot := range bots {
		streak := 0
		if bot.px > w-requiredStreak {
			continue
		}
		for i := bot.px + (w * bot.py); i%w != 0; i += 1 {
			if botGrid[i] == 0 {
				break
			} else {
				streak += 1
			}
			if streak > requiredStreak {
				return true
			}
		}
	}
	return false
}

func Level2(in io.Reader) string {
	bots := parse(in)
	botGrid := [w * h]int{}
	for _, bot := range bots {
		botGrid[w*bot.py+bot.px] += 1
	}
	for i := 1; true; i += 1 {
		for b, bot := range bots {
			botGrid[w*bot.py+bot.px] -= 1
			bots[b].px = (bot.px + bot.vx + w) % w
			bots[b].py = (bot.py + bot.vy + h) % h
			botGrid[w*bots[b].py+bots[b].px] += 1
		}
		if hasStreak2(bots, botGrid[:]) && hasCavity(botGrid[:]) {
			// for y := 0; y < h; y += 1 {
			// 	for x := 0; x < w; x += 1 {
			// 		fmt.Print(botGrid[w*x+y])
			// 	}
			// 	fmt.Println()
			// }
			return fmt.Sprint(i)
		}
	}

	return ""
}
