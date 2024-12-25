package day25

import (
	"bufio"
	"fmt"
	"io"
)

func parse(in io.Reader) ([][]int, [][]int) {
	scanner := bufio.NewScanner(in)
	keys := [][]int{}
	locks := [][]int{}
	for scanner.Scan() {
		if scanner.Bytes()[0] == '#' {
			// lock
			// These are the height of the columns
			lock := []int{5, 5, 5, 5, 5}
			for i := 0; i < 5; i += 1 {
				scanner.Scan()
				line := scanner.Bytes()
				for j, l := range lock {
					if l == 5 && line[j] == '.' {
						lock[j] = i
					}
				}
			}
			locks = append(locks, lock)
		} else {
			// key
			// These are the gaps between the top and the cut
			key := []int{0, 0, 0, 0, 0}
			for i := 0; i < 5; i += 1 {
				scanner.Scan()
				line := scanner.Bytes()
				for j, k := range key {
					if k == 0 && line[j] == '#' {
						key[j] = 5 - i
					}
				}
			}
			keys = append(keys, key)
		}
		scanner.Scan() // bottom row contains no new information
		scanner.Scan() // skip the newline
	}
	return locks, keys
}

func Level1(in io.Reader) string {
	// I think there's a better solution with a radix sort but I'm not doing that, O(n^2), n=500 is fine
	locks, keys := parse(in)
	total := 0
	for _, lock := range locks {
		for _, key := range keys {
			// fmt.Println(lock, key)
			fit := true
			for i := 0; i < 5; i += 1 {
				if key[i]+lock[i] > 5 {
					fit = false
					break
				}
			}
			if fit {
				total += 1
			}
		}
	}
	return fmt.Sprint(total)
}

func Level2(in io.Reader) string {
	return ""
}
