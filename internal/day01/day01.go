package day01

import (
	"bufio"
	"fmt"
	"io"
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

func parse(in io.Reader) ([]int, []int) {
	scanner := bufio.NewScanner(in)
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

func Level1(in io.Reader) string {
	lefts, rights := parse(in)

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
	return fmt.Sprint(totalDifferences)
}

func Level2(in io.Reader) string {
	lefts, rights := parse(in)
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
	return fmt.Sprint(similarityScore)
}
