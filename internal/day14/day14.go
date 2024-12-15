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

func floodFill(bots []int, visited *[]bool, x int, y int) int {
	i := y*w + x
	if x < 0 || x >= w || y < 0 || y >= h || (*visited)[i] || bots[i] > 0 {
		return 0
	}
	(*visited)[i] = true
	return 1 + floodFill(bots, visited, x-1, y) + floodFill(bots, visited, x+1, y) + floodFill(bots, visited, x, y-1) + floodFill(bots, visited, x, y+1)
}

func Level2(in io.Reader) string {
	bots := parse(in)
	botGrid := [w * h]int{}
	for _, bot := range bots {
		botGrid[w*bot.py+bot.px] += 1
	}
	visited := make([]bool, w*h)
	for i := 1; true; i += 1 {
		for b, bot := range bots {
			botGrid[w*bot.py+bot.px] -= 1
			bots[b].px = (bot.px + bot.vx + w) % w
			bots[b].py = (bot.py + bot.vy + h) % h
			botGrid[w*bots[b].py+bots[b].px] += 1
		}
		// Start at 1,1 to avoid getting an abnormally small area by being boxed in a corner
		j := w + 2
		for botGrid[j] > 0 {
			j += 1
		}
		x := j % w
		y := j / w
		visited = make([]bool, w*h)
		outerArea := floodFill(botGrid[:], &visited, x, y)
		if outerArea > 100 && outerArea < 9900 {
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
