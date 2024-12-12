package day12

import (
	"bufio"
	"fmt"
	"io"
)

func parse(in io.Reader) [][]byte {
	scanner := bufio.NewScanner(in)
	grid := make([][]byte, 0, 1000)
	for scanner.Scan() {
		line := scanner.Bytes()
		grid = append(grid, line)
	}
	return grid
}

func floodFill(grid [][]byte, visited *[]bool, x int, y int, c byte) (int, int) {
	size := len(grid)
	if x < 0 || x >= size || y < 0 || y >= size {
		// We have left the bounds of the garden, so we need to increment the perimeter
		return 0, 1
	}
	if grid[y][x] != c {
		// We have left the region, so we need to increment the perimeter
		return 0, 1
	}
	if (*visited)[len(grid)*y+x] {
		// We've already checked this square
		return 0, 0
	}
	(*visited)[len(grid)*y+x] = true
	aLeft, pLeft := floodFill(grid, visited, x-1, y, c)
	aRight, pRight := floodFill(grid, visited, x+1, y, c)
	aUp, pUp := floodFill(grid, visited, x, y-1, c)
	aDown, pDown := floodFill(grid, visited, x, y+1, c)
	area := 1 + aLeft + aRight + aUp + aDown
	perimeter := pLeft + pRight + pUp + pDown
	return area, perimeter
}

func Level1(in io.Reader) string {
	grid := parse(in)
	size := len(grid)
	visited := make([]bool, size*size)
	total := 0
	for y, line := range grid {
		for x, c := range line {
			if visited[y*size+x] {
				continue
			}
			area, perimeter := floodFill(grid, &visited, x, y, c)
			fmt.Println(area, perimeter)
			total += area * perimeter
		}
	}
	return fmt.Sprint(total)
}

func Level2(in io.Reader) string {
	return ""
}
