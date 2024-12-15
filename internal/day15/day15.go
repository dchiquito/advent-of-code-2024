package day15

import (
	"bufio"
	"fmt"
	"io"
)

const Space = 0
const Wall = 1
const Box = 2

const Up = 0
const Right = 1
const Down = 2
const Left = 3

func parse(in io.Reader) ([][]int, []int, int, int) {
	grid := make([][]int, 0, 50)
	moves := make([]int, 0, 20000)
	rx := 0
	ry := 0
	scanner := bufio.NewScanner(in)
	y := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			break
		}
		arr := make([]int, 0, 50)
		for x, b := range line {
			if b == '.' {
				arr = append(arr, Space)
			} else if b == '#' {
				arr = append(arr, Wall)
			} else if b == 'O' {
				arr = append(arr, Box)
			} else if b == '@' {
				arr = append(arr, Space)
				rx = x
				ry = y
			}
		}
		grid = append(grid, arr)
		y += 1
	}
	for scanner.Scan() {
		line := scanner.Bytes()
		for _, c := range line {
			if c == '^' {
				moves = append(moves, Up)
			} else if c == '>' {
				moves = append(moves, Right)
			} else if c == 'v' {
				moves = append(moves, Down)
			} else if c == '<' {
				moves = append(moves, Left)
			}
		}
	}
	return grid, moves, rx, ry
}

func printGrid(grid [][]int, rx int, ry int) {
	for y, row := range grid {
		for x, c := range row {
			if x == rx && y == ry {
				fmt.Print("@")
			} else if c == Space {
				fmt.Print(".")
			} else if c == Wall {
				fmt.Print("#")
			} else if c == Box {
				fmt.Print("O")
			}
		}
		fmt.Println()
	}
}

func DoMove(grid *[][]int, rx int, ry int, d int) (int, int) {
	if d == Up {
		y := ry - 1
		for (*grid)[y][rx] == Box {
			y -= 1
		}
		if (*grid)[y][rx] == Space {
			(*grid)[y][rx] = Box
			(*grid)[ry-1][rx] = Space
			return rx, ry - 1
		} else {
			return rx, ry
		}
	}
	if d == Down {
		y := ry + 1
		for (*grid)[y][rx] == Box {
			y += 1
		}
		if (*grid)[y][rx] == Space {
			(*grid)[y][rx] = Box
			(*grid)[ry+1][rx] = Space
			return rx, ry + 1
		} else {
			return rx, ry
		}
	}
	if d == Left {
		x := rx - 1
		for (*grid)[ry][x] == Box {
			x -= 1
		}
		if (*grid)[ry][x] == Space {
			(*grid)[ry][x] = Box
			(*grid)[ry][rx-1] = Space
			return rx - 1, ry
		} else {
			return rx, ry
		}
	}
	if d == Right {
		x := rx + 1
		for (*grid)[ry][x] == Box {
			x += 1
		}
		if (*grid)[ry][x] == Space {
			(*grid)[ry][x] = Box
			(*grid)[ry][rx+1] = Space
			return rx + 1, ry
		} else {
			return rx, ry
		}
	}
	return 0, 0
}

func Level1(in io.Reader) string {
	grid, moves, x, y := parse(in)
	fmt.Println(moves)
	printGrid(grid, x, y)
	for _, move := range moves {
		x, y = DoMove(&grid, x, y, move)
	}
	printGrid(grid, x, y)
	total := 0
	for y, row := range grid {
		for x, c := range row {
			if c == Box {
				total += y*100 + x
			}
		}
	}
	return fmt.Sprint(total)
}

func Level2(in io.Reader) string {
	return ""
}
