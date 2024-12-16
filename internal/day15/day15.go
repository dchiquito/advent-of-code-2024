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
	for _, move := range moves {
		x, y = DoMove(&grid, x, y, move)
	}
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

func parse2(in io.Reader) ([][]int, []int, int, int) {
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
		arr := make([]int, 0, 100)
		for x, b := range line {
			if b == '.' {
				arr = append(arr, Space)
				arr = append(arr, Space)
			} else if b == '#' {
				arr = append(arr, Wall)
				arr = append(arr, Wall)
			} else if b == 'O' {
				arr = append(arr, Box)
				arr = append(arr, Space)
			} else if b == '@' {
				arr = append(arr, Space)
				arr = append(arr, Space)
				rx = 2 * x
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

func DoMove2(grid *[][]int, rx int, ry int, d int) (int, int) {
	if d == Left && PushLeft(grid, rx, ry) {
		return rx - 1, ry
	}
	if d == Right && PushRight(grid, rx, ry) {
		return rx + 1, ry
	}
	if d == Up && PushUp(grid, rx, ry) {
		return rx, ry - 1
	}
	if d == Down && PushDown(grid, rx, ry) {
		return rx, ry + 1
	}
	return rx, ry
}

// Take the thing at (x,y) and shift it left if possible.
func PushLeft(grid *[][]int, x int, y int) bool {
	if (*grid)[y][x-1] == Wall {
		return false
	}
	if x > 1 && (*grid)[y][x-2] == Box {
		// There is a box to the left. Try to push it
		if !PushLeft(grid, x-2, y) {
			return false
		}
	}
	// There was either no box, or a pushable box. Move left
	(*grid)[y][x-1] = (*grid)[y][x]
	(*grid)[y][x] = Space
	return true
}

// Take the thing at (x,y) and shift it right if possible.
func PushRight(grid *[][]int, x int, y int) bool {
	var delta int
	if (*grid)[y][x] == Box {
		// We are moving a box, so the hitbox is larger
		delta = 2
	} else {
		// We are moving the robot, smaller hitbox
		delta = 1
	}
	if (*grid)[y][x+delta] == Wall {
		return false
	}
	if x < 100-delta && (*grid)[y][x+delta] == Box {
		// There is a box to the right. Try to push it
		if !PushRight(grid, x+delta, y) {
			return false
		}
	}
	// There was either no box, or a pushable box. Move right
	(*grid)[y][x+1] = (*grid)[y][x]
	(*grid)[y][x] = Space
	return true
}

func PushDown(grid *[][]int, x int, y int) bool {
	if CanPushDown(grid, x, y) {
		if (*grid)[y+1][x-1] == Box {
			PushDown(grid, x-1, y+1)
		}
		if (*grid)[y+1][x] == Box {
			PushDown(grid, x, y+1)
		}
		if (*grid)[y][x] == Box && (*grid)[y+1][x+1] == Box {
			PushDown(grid, x+1, y+1)
		}
		(*grid)[y+1][x] = (*grid)[y][x]
		(*grid)[y][x] = Space
		return true
	} else {
		return false
	}
}

func CanPushDown(grid *[][]int, x int, y int) bool {
	if (*grid)[y][x] != Box {
		if (*grid)[y+1][x] == Wall {
			return false
		}
		return ((*grid)[y+1][x-1] != Box || CanPushDown(grid, x-1, y+1)) && ((*grid)[y+1][x] != Box || CanPushDown(grid, x, y+1))
	} else {
		if (*grid)[y+1][x] == Wall || (*grid)[y+1][x+1] == Wall {
			return false
		}
		return ((*grid)[y+1][x-1] != Box || CanPushDown(grid, x-1, y+1)) && ((*grid)[y+1][x] != Box || CanPushDown(grid, x, y+1)) && ((*grid)[y+1][x+1] != Box || CanPushDown(grid, x+1, y+1))
	}
}
func PushUp(grid *[][]int, x int, y int) bool {
	if CanPushUp(grid, x, y) {
		if (*grid)[y-1][x-1] == Box {
			PushUp(grid, x-1, y-1)
		}
		if (*grid)[y-1][x] == Box {
			PushUp(grid, x, y-1)
		}
		if (*grid)[y][x] == Box && (*grid)[y-1][x+1] == Box {
			PushUp(grid, x+1, y-1)
		}
		(*grid)[y-1][x] = (*grid)[y][x]
		(*grid)[y][x] = Space
		return true
	} else {
		return false
	}
}

func CanPushUp(grid *[][]int, x int, y int) bool {
	if (*grid)[y][x] != Box {
		if (*grid)[y-1][x] == Wall {
			return false
		}
		return ((*grid)[y-1][x-1] != Box || CanPushUp(grid, x-1, y-1)) && ((*grid)[y-1][x] != Box || CanPushUp(grid, x, y-1))
	} else {
		if (*grid)[y-1][x] == Wall || (*grid)[y-1][x+1] == Wall {
			return false
		}
		return ((*grid)[y-1][x-1] != Box || CanPushUp(grid, x-1, y-1)) && ((*grid)[y-1][x] != Box || CanPushUp(grid, x, y-1)) && ((*grid)[y-1][x+1] != Box || CanPushUp(grid, x+1, y-1))
	}
}

func Level2(in io.Reader) string {
	grid, moves, x, y := parse2(in)
	for _, move := range moves {
		x, y = DoMove2(&grid, x, y, move)
	}
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
