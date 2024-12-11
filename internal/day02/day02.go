package day02

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

func parse(in io.Reader) [][]int {
	scanner := bufio.NewScanner(in)
	reports := make([][]int, 0, 1000)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Fields(line)
		report := make([]int, len(lineSplit))
		for i, levelString := range lineSplit {
			level, err := strconv.Atoi(levelString)
			util.Check(err, "Invalid level string")
			report[i] = level
		}
		reports = append(reports, report)
	}
	return reports
}

func isDeltaSafe(a int, b int, ascending bool) bool {
	var delta int
	if ascending {
		delta = b - a
	} else {
		delta = a - b
	}
	return delta >= 1 && delta <= 3
}

func isSafe(report []int) bool {
	ascending := report[0] < report[1]
	for i := 1; i < len(report); i += 1 {
		if !isDeltaSafe(report[i-1], report[i], ascending) {
			return false
		}
	}
	return true
}

func Level1(in io.Reader) string {
	reports := parse(in)
	safeReports := 0
	for _, report := range reports {
		if isSafe(report) {
			safeReports += 1
		}
	}
	return fmt.Sprint(safeReports)
}

func isSafeDamped(report []int) bool {
	// trimmedReport := make([]int, len(report)-1)
	// for i := range report {
	// 	copy(trimmedReport, report[:i])
	// 	copy(trimmedReport[i:], report[i+1:])
	// 	if isSafe(trimmedReport) {
	// 		return true
	// 	}
	// }
	// return false
	// Try deleting the first element
	if isSafe(report[1:]) {
		return true
	}
	// Try deleting the last element
	if isSafe(report[:len(report)-1]) {
		return true
	}
	// Try deleting the second element
	ascending := report[0] < report[2]
	if isDeltaSafe(report[0], report[2], ascending) {
		safe := true
		for i := 3; i < len(report); i += 1 {
			if !isDeltaSafe(report[i-1], report[i], ascending) {
				safe = false
				break
			}
		}
		if safe {
			return true
		}
	}
	// Try deleting the remaining elements
	ascending = report[0] < report[1]
	if isDeltaSafe(report[0], report[1], ascending) {
		for x := 2; x < len(report)-1; x += 1 {
			if !isDeltaSafe(report[x-1], report[x+1], ascending) {
				continue
			}
			safe := true
			for i := 2; i < x; i += 1 {
				if !isDeltaSafe(report[i-1], report[i], ascending) {
					safe = false
					break
				}
			}
			if !safe {
				continue
			}
			for i := x + 2; i < len(report); i += 1 {
				if !isDeltaSafe(report[i-1], report[i], ascending) {
					safe = false
					break
				}
			}
			if safe {
				return true
			}
		}
	}
	return false
}

func Level2(in io.Reader) string {
	reports := parse(in)
	util.Stopwatch("parse")
	safeReports := 0
	for _, report := range reports {
		if isSafeDamped(report) {
			safeReports += 1
		}
	}
	return fmt.Sprint(safeReports)
}
