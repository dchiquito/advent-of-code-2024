package day16

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
)

func parse(in io.Reader) ([][]byte, int, int, int, int) {
	grid := make([][]byte, 0, 150)
	sx := 0
	sy := 0
	ex := 0
	ey := 0
	scanner := bufio.NewScanner(in)
	y := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			break
		}
		arr := make([]byte, 150)
		copy(arr, line)
		for x, b := range arr {
			if b == 'S' {
				arr[x] = '.'
				sx = x
				sy = y
			} else if b == 'E' {
				arr[x] = '.'
				ex = x
				ey = y
			}
		}
		grid = append(grid, arr)
		y += 1
	}
	return grid, sx, sy, ex, ey
}

func pack(x int, y int, d int) int {
	return (y * 150 * 4) + (x * 4) + d
}

func unpack(p int) (int, int, int) {
	d := p % 4
	p /= 4
	x := p % 150
	y := p / 150
	return x, y, d
}

// Heap implementation from the Go Docs
// An Item is something we manage in a priority queue.
type Item struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	// I tweaked this since I want lowest to highest
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// End heap implementation

const Up int = 0
const Right int = 1
const Down int = 2
const Left int = 3

func step(x int, y int, d int) (int, int) {
	switch d {
	case Up:
		return x, y - 1
	case Right:
		return x + 1, y
	case Down:
		return x, y + 1
	case Left:
		return x - 1, y
	}
	return -1, -1
}

func walk(grid [][]byte, visited *[][]int, q *PriorityQueue) {
	item := q.Pop().(*Item)
	x, y, d := unpack(item.value)
	if (*visited)[y][x] == 0 || (*visited)[y][x] > item.priority {
		(*visited)[y][x] = item.priority
		// Walk forward
		nx, ny := step(x, y, d)
		if grid[ny][nx] == '.' {
			q.Push(&Item{value: pack(nx, ny, d), priority: item.priority + 1})
		}
		// Walk right
		nx, ny = step(x, y, (d+1)%4)
		if grid[ny][nx] == '.' {
			q.Push(&Item{value: pack(nx, ny, (d+1)%4), priority: item.priority + 1001})
		}
		// Walk left
		nx, ny = step(x, y, (d+3)%4)
		if grid[ny][nx] == '.' {
			q.Push(&Item{value: pack(nx, ny, (d+3)%4), priority: item.priority + 1001})
		}
		heap.Init(q)
	}

}

func Level1(in io.Reader) string {
	grid, sx, sy, ex, ey := parse(in)
	visited := make([][]int, len(grid))
	q := make(PriorityQueue, 0, 150*150)
	for i := range visited {
		visited[i] = make([]int, len(grid[i]))
	}
	q.Push(&Item{value: pack(sx, sy, 1), priority: 0})
	heap.Init(&q)
	// TODO heap not sorting :( Gotta plow through the whole thing
	for q.Len() > 0 {
		// for visited[ey][ex] == 0 {
		// for i := 0; i < 300; i += 1 {
		walk(grid, &visited, &q)
	}
	// for y, row := range visited {
	// 	for x, p := range row {
	// 		if grid[y][x] == '#' {
	// 			fmt.Print("#")
	// 		} else {
	// 			if p == 0 {
	// 				fmt.Print(".")
	// 			} else {
	// 				fmt.Print("X")
	// 			}
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// for q.Len() > 0 {
	// 	fmt.Println(q.Len())
	// 	fmt.Println(q.Pop())
	// }
	return fmt.Sprint(visited[ey][ex])
}

func walkBack(weights [][]int, visited *[][]bool, x int, y int, prev int, sx int, sy int) bool {
	if x == sx && y == sy {
		(*visited)[y][x] = true
		return true
	}
	if x < 0 || x >= 140 || y < 0 || y >= 140 {
		return false
	}
	w := weights[y][x]
	if (w == prev-1 || w == prev-1001 || w == prev+999) && w > 0 {
		a := walkBack(weights, visited, x-1, y, w, sx, sy)
		b := walkBack(weights, visited, x+1, y, w, sx, sy)
		c := walkBack(weights, visited, x, y-1, w, sx, sy)
		d := walkBack(weights, visited, x, y+1, w, sx, sy)
		if a || b || c || d {
			(*visited)[y][x] = true
			return true
		}
	}
	return false
}

func Level2(in io.Reader) string {
	grid, sx, sy, ex, ey := parse(in)
	weights := make([][]int, len(grid))
	visited := make([][]bool, len(grid))
	q := make(PriorityQueue, 0, 150*150)
	for i := range weights {
		weights[i] = make([]int, len(grid[i]))
		visited[i] = make([]bool, len(grid[i]))
	}
	q.Push(&Item{value: pack(sx, sy, 1), priority: 0})
	heap.Init(&q)
	// TODO heap not sorting :( Gotta plow through the whole thing
	for q.Len() > 0 {
		// for visited[ey][ex] == 0 {
		// for i := 0; i < 300; i += 1 {
		walk(grid, &weights, &q)
	}
	walkBack(weights, &visited, ex, ey, weights[ey][ex]+1, sx, sy)

	// for y, row := range visited {
	// 	for x, p := range row {
	// 		// if grid[y][x] == '#' {
	// 		// 	fmt.Print(".")
	// 		// } else {
	// 		// 	if p {
	// 		// 		fmt.Print("X")
	// 		// 	} else {
	// 		// 		fmt.Print(" ")
	// 		// 	}
	// 		// }
	// 		if p {
	// 			fmt.Print("X")
	// 		} else if grid[y][x] == '#' {
	// 			fmt.Print(".")
	// 		} else {
	// 			fmt.Print(" ")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	tt := 0
	for _, row := range visited {
		for _, b := range row {
			if b {
				tt += 1
			}
		}
	}
	// TODO technically wrong on the second example
	// Because it takes a slower path back from the end initially
	return fmt.Sprint(tt)
}
