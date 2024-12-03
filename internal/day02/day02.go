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
	delta := b - a
	if !ascending {
		delta = -delta
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
	trimmedReport := make([]int, len(report)-1)
	for i := range report {
		copy(trimmedReport, report[:i])
		copy(trimmedReport[i:], report[i+1:])
		if isSafe(trimmedReport) {
			return true
		}
	}
	return false
}

func Level2(in io.Reader) string {
	reports := parse(in)
	safeReports := 0
	for _, report := range reports {
		if isSafeDamped(report) {
			safeReports += 1
		}
	}
	return fmt.Sprint(safeReports)
}
