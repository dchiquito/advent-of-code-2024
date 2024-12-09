package day09

import (
	"fmt"
	"io"
)

func parse(in io.Reader) []byte {
	input, _ := io.ReadAll(in)
	input = input[:len(input)-1]
	for i := range input {
		input[i] -= 48
	}
	return input
}

// A more performant, but wrong, solution
func Level1_wrong(in io.Reader) string {
	bytes := parse(in)
	// fmt.Println(bytes)
	var leftIndex int64 = 0
	var leftSubIndex byte = 0
	var rightIndex int64 = int64(len(bytes) - 1)
	var rightSubIndex byte = 0
	var total int64 = 0
	var blockId int64 = 0
	for leftIndex < rightIndex || rightSubIndex+leftSubIndex < bytes[leftIndex] {
		fmt.Println(leftIndex, leftSubIndex, bytes[leftIndex], "     ", rightIndex, rightSubIndex, bytes[rightIndex], "      ", leftIndex < rightIndex, rightSubIndex+leftSubIndex <= bytes[leftIndex])
		if leftIndex%2 == 0 {
			// Even left indices indicate file blocks
			fileId := leftIndex / 2
			total += fileId * blockId
			// fmt.Println("in place", fileId, "*", blockId, "=", fileId*blockId, "=>", total)
			leftSubIndex += 1
			if leftSubIndex >= bytes[leftIndex] {
				leftIndex += 1
				leftSubIndex = 0
			}
		} else {
			fileId := rightIndex / 2
			total += fileId * blockId
			// fmt.Println("copied  ", fileId, "*", blockId, "=", fileId*blockId, "=>", total)
			leftSubIndex += 1
			for leftSubIndex >= bytes[leftIndex] {
				leftIndex += 1
				leftSubIndex = 0
			}
			// fmt.Println("before", rightIndex, rightSubIndex, bytes[rightIndex])
			rightSubIndex += 1
			for rightSubIndex >= bytes[rightIndex] {
				// decrement by two to skip the empty sections
				rightIndex -= 2
				rightSubIndex = 0
			}
			// fmt.Println("after", rightIndex, rightSubIndex, bytes[rightIndex])
		}
		blockId += 1
		fmt.Println(blockId, total)
	}
	return fmt.Sprint(total)
}

// A dumb, naive solution that allocates and fills the entire array instead of doing it intelligently
func Level1(in io.Reader) string {
	bytes := parse(in)
	length := 0
	for _, b := range bytes {
		length += int(b)
	}
	arr := make([]int, length)
	arrIndex := 0
	for i, b := range bytes {
		var filler int
		if i%2 == 0 {
			filler = i / 2
		} else {
			filler = 0
		}
		for j := 0; j < int(b); j += 1 {
			arr[arrIndex] = filler
			arrIndex += 1
		}
	}
	leftIndex := int(bytes[0])
	rightIndex := len(arr) - 1
	for leftIndex < rightIndex {
		if arr[leftIndex] == 0 {
			arr[leftIndex] = arr[rightIndex]
			rightIndex -= 1
			for arr[rightIndex] == 0 {
				rightIndex -= 1
			}
		}
		leftIndex += 1
	}
	checksum := 0
	for i, a := range arr[:leftIndex+1] {
		checksum += i * a
	}
	return fmt.Sprint(checksum)
}

func Level2(in io.Reader) string {
	return ""
}
