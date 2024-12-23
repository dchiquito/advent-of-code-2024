package day23

import (
	"bufio"
	"fmt"
	"io"
)

type Graph [][]int

func parse(in io.Reader) Graph {
	scanner := bufio.NewScanner(in)
	graph := make(Graph, 26*26)
	for scanner.Scan() {
		line := scanner.Bytes()
		a := int(line[0] - 'a')
		b := int(line[1] - 'a')
		c := int(line[3] - 'a')
		d := int(line[4] - 'a')
		left := a*26 + b
		right := c*26 + d
		if graph[left] == nil {
			graph[left] = []int{}
		}
		graph[left] = append(graph[left], right)
		if graph[right] == nil {
			graph[right] = []int{}
		}
		graph[right] = append(graph[right], left)
	}
	return graph
}

func isAdjacent(graph Graph, n1 int, n2 int) bool {
	for _, adj := range graph[n1] {
		if adj == n2 {
			return true
		}
	}
	return false
}

func Level1(in io.Reader) string {
	graph := parse(in)
	total := 0
	for i := 26 * 19; i < 26*20; i += 1 {
		if graph[i] != nil {
			for j := 0; j < len(graph[i]); j += 1 {
				nj := graph[i][j]
				if nj/26 == 19 && nj < i {
					// Avoid double counting
					continue
				}
				for k := j + 1; k < len(graph[i]); k += 1 {
					nk := graph[i][k]
					if nk/26 == 19 && nk < i {
						// Avoid double counting
						continue
					}
					if isAdjacent(graph, nj, nk) {
						total += 1
						break
					}
				}
			}
		}
	}
	return fmt.Sprint(total)
}

func Level2(in io.Reader) string {
	return ""
}
