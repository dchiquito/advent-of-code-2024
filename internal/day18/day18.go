package day18

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

const Size int = 71

func parse1(in io.Reader) []bool {
	scanner := bufio.NewScanner(in)
	grid := make([]bool, Size*Size)
	for i := 0; i < 1024; i += 1 {
		scanner.Scan()
		line := scanner.Bytes()
		_, x := util.ChompInt(line, 0)
		_, y := util.ChompInt(line, 2)
		grid[(y*Size)+x] = true
	}
	return grid
}

func shortestPath(grid []bool) int {
	weights := make([]int, Size*Size)
	toVisit := []int{0}
	for len(toVisit) > 0 {
		i := toVisit[0]
		toVisit = toVisit[1:]
		if grid[i] {
			// We've tried to walk into a falling byte
			continue
		}
		if i == len(grid)-1 {
			break
		}
		x := i % Size
		y := i / Size
		// Left
		i2 := (y * Size) + x - 1
		if x > 0 && (weights[i2] == 0 || weights[i2] > weights[i]+1) {
			toVisit = append(toVisit, i2)
			weights[i2] = weights[i] + 1
		}
		// Up
		i2 = ((y - 1) * Size) + x
		if y > 0 && (weights[i2] == 0 || weights[i2] > weights[i]+1) {
			toVisit = append(toVisit, i2)
			weights[i2] = weights[i] + 1
		}
		// Right
		i2 = (y * Size) + x + 1
		if x < Size-1 && (weights[i2] == 0 || weights[i2] > weights[i]+1) {
			toVisit = append(toVisit, i2)
			weights[i2] = weights[i] + 1
		}
		// Down
		i2 = ((y + 1) * Size) + x
		if y < Size-1 && (weights[i2] == 0 || weights[i2] > weights[i]+1) {
			toVisit = append(toVisit, i2)
			weights[i2] = weights[i] + 1
		}
	}
	return weights[len(weights)-1]
}

func Level1(in io.Reader) string {
	grid := parse1(in)
	solution := shortestPath(grid)
	return fmt.Sprint(solution)
}

func parse2(in io.Reader) ([]bool, []int) {
	scanner := bufio.NewScanner(in)
	grid := make([]bool, Size*Size)
	for i := 0; i < 1024; i += 1 {
		scanner.Scan()
		line := scanner.Bytes()
		_, x := util.ChompInt(line, 0)
		_, y := util.ChompInt(line, 2)
		grid[(y*Size)+x] = true
	}
	additional := []int{}
	for scanner.Scan() {
		line := scanner.Bytes()
		_, x := util.ChompInt(line, 0)
		_, y := util.ChompInt(line, 2)
		additional = append(additional, (y*Size)+x)
	}
	return grid, additional
}

func Level2(in io.Reader) string {
	grid, additional := parse2(in)
	for _, i := range additional {
		grid[i] = true
		if shortestPath(grid) == 0 {
			x := i % Size
			y := i / Size
			return fmt.Sprintf("%d,%d", x, y)
		}
	}
	return ""
}
