package day10

import (
	"bufio"
	"fmt"
	"io"
)

func parse(in io.Reader) [][]byte {
	scanner := bufio.NewScanner(in)
	grid := make([][]byte, 0, 50)
	for scanner.Scan() {
		line := scanner.Bytes()
		for i := range line {
			line[i] -= 48
		}
		grid = append(grid, line)
	}
	return grid
}

func ascend(grid [][]byte, x int, y int, visited []int) (int, []int) {
	height := grid[y][x]
	if height == 9 {
		for _, v := range visited {
			if v == y*10000+x {
				return 0, visited
			}
		}
		visited = append(visited, y*10000+x)
		return 1, visited
	}
	total := 0
	var hikes int
	if x > 0 && grid[y][x-1] == height+1 {
		hikes, visited = ascend(grid, x-1, y, visited)
		total += hikes
	}
	if x < len(grid[y])-1 && grid[y][x+1] == height+1 {
		hikes, visited = ascend(grid, x+1, y, visited)
		total += hikes
	}
	if y > 0 && grid[y-1][x] == height+1 {
		hikes, visited = ascend(grid, x, y-1, visited)
		total += hikes
	}
	if y < len(grid)-1 && grid[y+1][x] == height+1 {
		hikes, visited = ascend(grid, x, y+1, visited)
		total += hikes
	}
	return total, visited
}

func Level1(in io.Reader) string {
	grid := parse(in)
	total := 0
	for y, line := range grid {
		for x, h := range line {
			if h == 0 {
				// TODO does preallocating with a capacity speed things up?
				// visited := make([]int, 0)
				var visited []int
				hikes, _ := ascend(grid, x, y, visited)
				total += hikes
			}
		}
	}
	return fmt.Sprint(total)
}

func Level2(in io.Reader) string {
	return ""
}
