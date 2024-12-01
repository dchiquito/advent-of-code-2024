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

func parse() ([]int, []int) {
	scanner := bufio.NewScanner(os.Stdin)
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
	sort.Slice(lefts, func(i, j int) bool { return lefts[i] < lefts[j] })
	sort.Slice(rights, func(i, j int) bool { return rights[i] < rights[j] })
	return lefts, rights
}

func level1() {
	lefts, rights := parse()

	totalDifferences := 0
	for i := range lefts {
		left := lefts[i]
		right := rights[i]
		if left > right {
			totalDifferences += left - right
		} else {
			totalDifferences += right - left
		}
	}
	fmt.Println(totalDifferences)
}

func level2() {
	lefts, rights := parse()
	rightIndex := 0
	similarityScore := 0
	for _, left := range lefts {
		dupes := 0
		for rightIndex < len(rights) && rights[rightIndex] <= left {
			if rights[rightIndex] == left {
				dupes += 1
			}
			rightIndex += 1
		}
		similarityScore += left * dupes
	}
	fmt.Println(similarityScore)
}

func main() {
	if util.GetLevelArg() == 1 {
		level1()
	} else {
		level2()
	}
}
