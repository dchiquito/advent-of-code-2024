package day06

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

const Size = 130

// const Size = 10

type Cell int

const (
	Empty Cell = 0
	Block Cell = 1
)

func parse(in io.Reader) ([]Cell, int, int) {
	grid := make([]Cell, Size*Size)
	scanner := bufio.NewScanner(in)
	startX := 0
	startY := 0
	y := 0
	for scanner.Scan() {
		for x, c := range scanner.Text() {
			if c == '.' {
				grid[y*Size+x] = Empty
			} else if c == '#' {
				grid[y*Size+x] = Block
			} else if c == '^' {
				grid[y*Size+x] = Empty
				startX = x
				startY = y
			}
		}
		y += 1
	}
	return grid, startX, startY
}

func Level1(in io.Reader) string {
	grid, x, y := parse(in)
	var visited [Size * Size]bool
	visited[y*Size+x] = true
	dx := 0
	dy := -1
	for true {
		nx := x + dx
		ny := y + dy
		if nx < 0 || nx >= Size || ny < 0 || ny >= Size {
			break
		}
		if grid[ny*Size+nx] == Block {
			// Turn right
			t := dx
			dx = -dy
			dy = t
		} else {
			x = nx
			y = ny
			visited[y*Size+x] = true
		}
	}
	count := 0
	for _, v := range visited {
		if v {
			count += 1
		}
	}
	return fmt.Sprint(count)
}

type Dir int

const (
	Up    Dir = 1
	Down  Dir = 2
	Left  Dir = 4
	Right Dir = 8
)

func dir(dx int, dy int) Dir {
	if dx == -1 {
		return Left
	} else if dx == 1 {
		return Right
	} else if dy == -1 {
		return Up
	} else if dy == 1 {
		return Down
	}
	util.Panic("Undefined direction", dx, dy)
	return 0
}

func hasCycle(grid []Cell, x int, y int) bool {
	var visited [Size * Size]Dir
	visited[y*Size+x] = 1
	dx := 0
	dy := -1
	for true {
		nx := x + dx
		ny := y + dy
		if nx < 0 || nx >= Size || ny < 0 || ny >= Size {
			return false
		}
		if visited[ny*Size+nx]&dir(dx, dy) > 0 {
			// display(grid, visited[:])
			return true
		}
		if grid[ny*Size+nx] == Block {
			// Turn right
			t := dx
			dx = -dy
			dy = t
		} else {
			x = nx
			y = ny
		}
		visited[y*Size+x] |= dir(dx, dy)
	}
	util.Panic("unreachable")
	return false
}

func display(grid []Cell, visited []Dir) {
	// for y := 0; y < Size; y += 1 {
	// 	for x := 0; x < Size; x += 1 {
	// 		fmt.Print(visited[y*Size+x], " ")
	// 	}
	// 	fmt.Println()
	// }
	for y := 0; y < Size; y += 1 {
		for x := 0; x < Size; x += 1 {
			i := y*Size + x
			if grid[i] == Block {
				fmt.Print("#")
			} else {
				v := visited[i]&(Up|Down) > 0
				h := visited[i]&(Left|Right) > 0
				if v && h {
					fmt.Print("+")
				} else if v {
					fmt.Print("|")
				} else if h {
					fmt.Print("-")
				} else {
					fmt.Print(".")
				}
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func Level2(in io.Reader) string {
	grid, x, y := parse(in)
	count := 0
	for i := 0; i < Size*Size; i += 1 {
		if grid[i] != Empty || i == y*Size+x {
			continue
		}
		grid[i] = Block
		if hasCycle(grid, x, y) {
			count += 1
		}
		grid[i] = Empty
	}
	return fmt.Sprint(count)
}
