package day05

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Ordering struct {
	left  int
	right int
}

func parse(in io.Reader) ([]Ordering, [][]int) {
	scanner := bufio.NewScanner(in)
	orderings := make([]Ordering, 0, 1200)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		left, _ := strconv.Atoi(line[:2])
		right, _ := strconv.Atoi(line[3:])
		orderings = append(orderings, Ordering{left, right})
	}
	updates := make([][]int, 0, 200)
	for scanner.Scan() {
		line := scanner.Text()
		update := make([]int, 0, (len(line)+1)/3)
		for len(line) > 2 {
			page, _ := strconv.Atoi(line[:2])
			update = append(update, page)
			line = line[3:]
		}
		page, _ := strconv.Atoi(line[:2])
		update = append(update, page)
		updates = append(updates, update)
	}
	return orderings, updates
}

type Graph = [100][]int

func buildGraph(orderings []Ordering) Graph {
	var graph [100][]int
	for i := 0; i < 100; i += 1 {
		count := 0
		for _, order := range orderings {
			if order.left == i {
				count += 1
			}
		}
		node := make([]int, 0, count)
		for _, order := range orderings {
			if order.left == i {
				node = append(node, order.right)
			}
		}
		graph[i] = node
	}
	return graph
}

func isSorted(update []int, graph Graph) bool {
	// This tracks how many other nodes have references to the index node
	var references [100]int
	// for every possible node i,
	for _, i := range update {
		// check every other node j and count
		for _, j := range update {
			// which js point to i
			for _, ref := range graph[j] {
				if ref == i {
					references[i] += 1
				}
			}
		}
	}
	for _, i := range update {
		fmt.Println("thinkin on", update, i, references[i])
		// If another node points to the node in the update, then we are out of order
		if references[i] != 0 {
			return false
		}
		// Delete i by decrementing every node it references
		for _, j := range graph[i] {
			references[j] -= 1
		}
	}
	return true
}

func Level1(in io.Reader) string {
	orderings, updates := parse(in)
	graph := buildGraph(orderings)
	total := 0
	for _, update := range updates {
		if isSorted(update, graph) {
			total += update[len(update)/2]
		}
	}
	return fmt.Sprint(total)
}

func Level2(in io.Reader) string {
	return ""
}
