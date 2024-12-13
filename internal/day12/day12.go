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
		arr := make([]byte, len(line))
		copy(arr, line)
		grid = append(grid, arr)
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
	for y, line := range grid {
		for x, c := range line {
			if visited[y*size+x] {
				continue
			}
			area, perimeter := floodFill(grid, &visited, x, y, c)
			total += area * perimeter
			// if area != calcArea(grid, x, y, c) {
			// 	fmt.Println(area, perimeter, calcPerimeter(grid, x, y, c))
			// }
			// if perimeter != calcPerimeter(grid, x, y, c) {
			// 	fmt.Println(area, perimeter, calcPerimeter(grid, x, y, c))
			// }
		}
	}
	return fmt.Sprint(total)
}

func floodFill2(grid [][]byte, visited *[]bool, edges *map[int]bool, x int, y int, c byte, d int) int {
	size := len(grid)
	// check if we have left the bounds of the garden or region and need to increment the perimeter
	if x < 0 || x >= size || y < 0 || y >= size || grid[y][x] != c {
		if d == 0 {
			(*edges)[toEdge(size, x, y+1, d)] = true
		}
		if d == 1 {
			(*edges)[toEdge(size, x-1, y, d)] = true
		}
		if d == 2 {
			(*edges)[toEdge(size, x, y-1, d)] = true
		}
		if d == 3 {
			(*edges)[toEdge(size, x+1, y, d)] = true
		}
		return 0
	}
	if (*visited)[(len(grid)*y)+x] {
		// We've already checked this square
		return 0
	}
	(*visited)[len(grid)*y+x] = true
	//  0
	// 3 1
	//  2
	aLeft := floodFill2(grid, visited, edges, x-1, y, c, 3)
	aRight := floodFill2(grid, visited, edges, x+1, y, c, 1)
	aUp := floodFill2(grid, visited, edges, x, y-1, c, 0)
	aDown := floodFill2(grid, visited, edges, x, y+1, c, 2)
	area := 1 + aLeft + aRight + aUp + aDown
	return area
}

func toEdge(size int, x int, y int, d int) int {
	return (y * size * 4) + (x * 4) + d
}
func fromEdge(size int, e int) (int, int, int) {
	d := e % 4
	coord := e / 4
	x := coord % size
	y := coord / size
	return x, y, d
}
func countEdges(size int, edges map[int]bool) int {
	edgeNum := 0
	for len(edges) > 0 {
		edgeNum += 1
		var edge int
		for edge = range edges {
			break
		}
		delete(edges, edge)
		x, y, d := fromEdge(size, edge)
		//  0
		// 3 1
		//  2
		if d == 0 || d == 2 {
			// Facing up/down, the edge will extend left/right
			for dx := 1; x+dx < size; dx += 1 {
				adjEdge := toEdge(size, x+dx, y, d)
				if !edges[adjEdge] {
					break
				}
				delete(edges, adjEdge)
			}
			for dx := -1; x+dx >= 0; dx -= 1 {
				adjEdge := toEdge(size, x+dx, y, d)
				if !edges[adjEdge] {
					break
				}
				delete(edges, adjEdge)
			}
		} else {
			// Facing left/right, the edge will extend up/down
			for dy := 1; y+dy < size; dy += 1 {
				adjEdge := toEdge(size, x, y+dy, d)
				if !edges[adjEdge] {
					break
				}
				delete(edges, adjEdge)
			}
			for dy := -1; y+dy >= 0; dy -= 1 {
				adjEdge := toEdge(size, x, y+dy, d)
				if !edges[adjEdge] {
					break
				}
				delete(edges, adjEdge)
			}
		}
	}
	return edgeNum
}

func Level2(in io.Reader) string {
	grid := parse(in)
	size := len(grid)
	visited := make([]bool, size*size)
	total := 0
	for y, line := range grid {
		for x, c := range line {
			if visited[y*size+x] {
				continue
			}
			edges := map[int]bool{}
			area := floodFill2(grid, &visited, &edges, x, y, c, 2)
			numEdges := countEdges(size, edges)
			total += area * numEdges
		}
	}
	return fmt.Sprint(total)
}
