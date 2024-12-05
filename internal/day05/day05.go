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

func Level1(in io.Reader) string {
	orderings, updates := parse(in)
	fmt.Println(orderings, updates)
	return ""
}

func Level2(in io.Reader) string {
	return ""
}
