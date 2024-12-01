package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

func Quicksort(arr []int) {
	length := len(arr)
	if length < 2 {
		return
	}
	// p a b c d e f
	left := 0
	right := length - 1
	for left < right {
		if arr[left] < arr[left+1] {
			// TODO what swap here
			t := arr[right]
			arr[right] = arr[left+1]
			arr[left] = t
		}
		right -= 1
	}
}

func level1() {
	scanner := bufio.NewScanner(os.Stdin)
	total_differences := 0
	lefts := make([]int, 0, 1000)
	rights := make([]int, 0, 1000)
	for scanner.Scan() {
		line := scanner.Text()
		left, err := strconv.Atoi(line[:5])
		// left, err := strconv.Atoi(line[:1])
		util.Check(err, "malformed left input")
		lefts = append(lefts, left)
		right, err := strconv.Atoi(line[8:])
		// right, err := strconv.Atoi(line[4:])
		util.Check(err, "malformed right input")
		rights = append(rights, right)
	}
	sort.Slice(lefts, func(i, j int) bool { return lefts[i] > lefts[j] })
	sort.Slice(rights, func(i, j int) bool { return rights[i] > rights[j] })

	for i := range lefts {
		left := lefts[i]
		right := rights[i]
		if left > right {
			total_differences += left - right
		} else {
			total_differences += right - left
		}
	}
	fmt.Println(total_differences)
}

func level2() {
}

func main() {
	if util.GetLevelArg() == 1 {
		level1()
	} else {
		level2()
	}
}
