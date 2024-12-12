package day12

import (
	"bufio"
	"fmt"
	"io"
	"slices"
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
	if (*visited)[(len(grid)*y)+x] {
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

func calcArea(grid [][]byte, x int, y int, c byte) int {
	size := len(grid)
	visited := []int{}
	toVisit := []int{(y * size) + x}
	for len(toVisit) > 0 {
		v := toVisit[0]
		toVisit = toVisit[1:]
		if slices.Contains(visited, v) {
			continue
		}
		x := v % size
		y := v / size
		if grid[y][x] != c {
			continue
		}
		visited = append(visited, v)
		if x > 0 {
			toVisit = append(toVisit, (y*size)+x-1)
		}
		if x < size-1 {
			toVisit = append(toVisit, (y*size)+x+1)
		}
		if y > 0 {
			toVisit = append(toVisit, ((y-1)*size)+x)
		}
		if y < size-1 {
			toVisit = append(toVisit, ((y+1)*size)+x)
		}
	}
	return len(visited)
}

func calcPerimeter(grid [][]byte, x int, y int, c byte) int {
	size := len(grid)
	queue := make([]int, 1)
	queue[0] = (y * size) + x
	perim := 0
	for i := 0; i < len(queue); i += 1 {
		index := queue[i]
		x = index % size
		y = index / size
		if x == 0 || grid[y][x-1] != c {
			perim += 1
		} else {
			idx := (y * size) + (x - 1)
			if !slices.Contains(queue, idx) {
				queue = append(queue, idx)
			}
		}
		if x == size-1 || grid[y][x+1] != c {
			perim += 1
		} else {
			idx := (y * size) + (x + 1)
			if !slices.Contains(queue, idx) {
				queue = append(queue, idx)
			}
		}
		if y == 0 || grid[y-1][x] != c {
			perim += 1
		} else {
			idx := ((y - 1) * size) + x
			if !slices.Contains(queue, idx) {
				queue = append(queue, idx)
			}
		}
		if y == size-1 || grid[y+1][x] != c {
			perim += 1
		} else {
			idx := ((y + 1) * size) + x
			if !slices.Contains(queue, idx) {
				queue = append(queue, idx)
			}
		}
	}
	return perim
}

func Level1(in io.Reader) string {
	grid := parse(in)
	size := len(grid)
	visited := make([]bool, size*size)
	total := 0
	regions := 0
	totalArea := 0
	for y, line := range grid {
		for x, c := range line {
			if visited[y*size+x] {
				continue
			}
			area, perimeter := floodFill(grid, &visited, x, y, c)
			total += area * perimeter
			if area != calcArea(grid, x, y, c) {
				fmt.Println(area, perimeter, calcPerimeter(grid, x, y, c))
			}
			if perimeter != calcPerimeter(grid, x, y, c) {
				fmt.Println(area, perimeter, calcPerimeter(grid, x, y, c))
			}
			fmt.Println(area, perimeter)
			fmt.Println(calcArea(grid, x, y, c), calcPerimeter(grid, x, y, c))
			regions += 1
			totalArea += area
		}
	}
	// Check for missed plots
	for i, v := range visited {
		if !v {
			fmt.Println("MISSED ", i, v)
		}
	}
	fmt.Println(totalArea, "=", size, "x", size, "=", size*size)
	return fmt.Sprint(total)
}

func Level2(in io.Reader) string {
	return ""
}
