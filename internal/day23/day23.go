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

func largestSubgroup(graph Graph, n int) int {
	adjs := graph[n]
	notIncluded := make([]bool, len(adjs))
	numNotIncluded := 0
	var clumps []int
	for numNotIncluded < len(adjs) {
		clumps = make([]int, len(adjs))
		for i, ni := range adjs {
			if notIncluded[i] {
				continue
			}
			for j := i + 1; j < len(adjs); j += 1 {
				if notIncluded[j] {
					continue
				}
				nj := adjs[j]
				if isAdjacent(graph, ni, nj) {
					clumps[i] += 1
					clumps[j] += 1
				}
			}
		}
		smallest := 9999
		for i, c := range clumps {
			if !notIncluded[i] && c < smallest {
				smallest = c
			}
		}
		for i, c := range clumps {
			if !notIncluded[i] && c == smallest {
				notIncluded[i] = true
				numNotIncluded += 1
			}
		}
	}
	xxx := 1
	for _, c := range clumps {
		if c > 0 {
			xxx += 1
		}
	}
	return xxx
}

func Level2(in io.Reader) string {
	graph := parse(in)
	// Every node has 13 neighbors
	// The largest possible connected subgroup is 14
	largestGroupSize := 0
	var group []int
	for n := range graph {
		if graph[n] != nil {
			groupSize := largestSubgroup(graph, n)
			if groupSize > largestGroupSize {
				group = make([]int, 0, 14)
				largestGroupSize = groupSize
			}
			if groupSize == largestGroupSize {
				group = append(group, n)
			}
		}
	}
	// Conveniently, the list of nodes is already sorted
	// Build the string from the list of nodes
	arr := make([]byte, len(group)*3-1)
	arr[0] = byte(group[0]/26) + 'a'
	arr[1] = byte(group[0]%26) + 'a'
	for i := 1; i < len(group); i += 1 {
		arr[i*3-1] = ','
		arr[i*3] = byte(group[i]/26) + 'a'
		arr[i*3+1] = byte(group[i]%26) + 'a'
	}
	return string(arr)
}
