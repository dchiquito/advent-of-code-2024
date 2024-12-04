package day04

import (
	"fmt"
	"io"
)

const SIZE = 140

func parse(in io.Reader) [SIZE][SIZE]byte {
	var grid [SIZE][SIZE]byte
	var newlineBuffer [1]byte
	for row := 0; row < SIZE; row += 1 {
		in.Read(grid[row][:])
		in.Read(newlineBuffer[:])
	}
	return grid
}

const X = 88
const M = 77
const A = 65
const S = 83

func Level1(in io.Reader) string {
	grid := parse(in)
	total := 0
	// Horizontal
	for x := 0; x < SIZE-3; x += 1 {
		for y := 0; y < SIZE; y += 1 {
			if grid[y][x] == X && grid[y][x+1] == M && grid[y][x+2] == A && grid[y][x+3] == S {
				total += 1
			}
			if grid[y][x] == S && grid[y][x+1] == A && grid[y][x+2] == M && grid[y][x+3] == X {
				total += 1
			}
		}
	}
	// Vertical
	for x := 0; x < SIZE; x += 1 {
		for y := 0; y < SIZE-3; y += 1 {
			if grid[y][x] == X && grid[y+1][x] == M && grid[y+2][x] == A && grid[y+3][x] == S {
				total += 1
			}
			if grid[y][x] == S && grid[y+1][x] == A && grid[y+2][x] == M && grid[y+3][x] == X {
				total += 1
			}
		}
	}
	// Diagonal /
	for x := 0; x < SIZE-3; x += 1 {
		for y := 3; y < SIZE; y += 1 {
			if grid[y][x] == X && grid[y-1][x+1] == M && grid[y-2][x+2] == A && grid[y-3][x+3] == S {
				total += 1
			}
			if grid[y][x] == S && grid[y-1][x+1] == A && grid[y-2][x+2] == M && grid[y-3][x+3] == X {
				total += 1
			}
		}
	}
	// Diagonal \
	for x := 0; x < SIZE-3; x += 1 {
		for y := 0; y < SIZE-3; y += 1 {
			if grid[y][x] == X && grid[y+1][x+1] == M && grid[y+2][x+2] == A && grid[y+3][x+3] == S {
				total += 1
			}
			if grid[y][x] == S && grid[y+1][x+1] == A && grid[y+2][x+2] == M && grid[y+3][x+3] == X {
				total += 1
			}
		}
	}
	return fmt.Sprint(total)
}

func Level2(in io.Reader) string {
	grid := parse(in)
	total := 0
	for x := 0; x < SIZE-2; x += 1 {
		for y := 0; y < SIZE-2; y += 1 {
			// More explicit way
			// if ((grid[y][x] == M && grid[y+1][x+1] == A && grid[y+2][x+2] == S) ||
			// 	(grid[y][x] == S && grid[y+1][x+1] == A && grid[y+2][x+2] == M)) &&
			// 	((grid[y][x+2] == M && grid[y+1][x+1] == A && grid[y+2][x] == S) ||
			// 		(grid[y][x+2] == S && grid[y+1][x+1] == A && grid[y+2][x] == M)) {
			// 	total += 1
			// }
			// No redundant checks
			if grid[y+1][x+1] == A &&
				((grid[y][x] == M && grid[y+2][x+2] == S) || (grid[y][x] == S && grid[y+2][x+2] == M)) &&
				((grid[y][x+2] == M && grid[y+2][x] == S) || (grid[y][x+2] == S && grid[y+2][x] == M)) {
				total += 1
			}
		}
	}
	return fmt.Sprint(total)
}
