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

type Section struct {
	FileID       *int
	FileLength   byte
	SuffixLength byte
	Suffix       *Section
}

func (s Section) String() string {
	str := ""
	var i byte
	var fileid string
	if s.FileID != nil {
		fileid = fmt.Sprint(*s.FileID)
	} else {
		fileid = "."
	}
	for i = 0; i < s.FileLength; i += 1 {
		str += fileid
	}
	if s.Suffix != nil {
		str += s.Suffix.String()
	} else {
		for i = 0; i < s.SuffixLength; i += 1 {
			str += "."
		}
	}
	return str
}
func Print(sections []Section) {
	for _, section := range sections {
		fmt.Print(section.String())
	}
	fmt.Println()
}

func (s Section) Checksum(index int) (int, int) {
	checksum := 0
	if s.FileID != nil {
		for i := 0; i < int(s.FileLength); i += 1 {
			checksum += index * (*s.FileID)
			index += 1
		}
	} else {
		index += int(s.FileLength)
	}
	if s.Suffix != nil {
		i, c := s.Suffix.Checksum(index)
		index = i
		checksum += c
	} else {
		index += int(s.SuffixLength)
	}
	return index, checksum
}

func Level2(in io.Reader) string {
	bytes := parse(in)
	sections := make([]Section, (len(bytes)+1)/2)
	for i := 0; i < len(bytes)/2; i += 1 {
		fileLength := bytes[2*i]
		suffixLength := bytes[(2*i)+1]
		sections[i] = Section{&i, fileLength, suffixLength, nil}
	}
	// The last section must be handled specially
	lastIndex := len(sections) - 1
	sections[lastIndex] = Section{&lastIndex, bytes[lastIndex*2], 0, nil}
	var memos [10]int
	// r is the index of the file we are trying to move
	for r := len(sections) - 1; r > 0; r -= 1 {
		sectionToMove := &sections[r]
		for ; memos[sectionToMove.FileLength] < r; memos[sectionToMove.FileLength] += 1 {
			found := false
			sectionWithGap := &sections[memos[sectionToMove.FileLength]]
			for sectionWithGap.SuffixLength >= sectionToMove.FileLength {
				if sectionWithGap.Suffix == nil {
					// Copy sectionToMove into the suffix
					sectionWithGap.Suffix = &Section{
						sectionToMove.FileID,
						sectionToMove.FileLength,
						sectionWithGap.SuffixLength - sectionToMove.FileLength,
						nil,
					}
					// Delete the original by niling the file ID
					sectionToMove.FileID = nil
					found = true
					break
				} else {
					sectionWithGap = sectionWithGap.Suffix
				}
			}
			if found {
				// We found somewhere to cram it, stop looking
				break
			}
		}
	}
	index := 0
	checksum := 0
	for _, section := range sections {
		i, c := section.Checksum(index)
		index = i
		checksum += c
	}
	return fmt.Sprint(checksum)
}
