package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func isSafe(report []int) bool {
	if report[0] < report[1] {
		for i := 1; i < len(report); i += 1 {
			delta := report[i] - report[i-1]
			if delta < 1 || delta > 3 {
				return false
			}
		}
	} else {
		for i := 1; i < len(report); i += 1 {
			delta := report[i] - report[i-1]
			if delta > -1 || delta < -3 {
				return false
			}
		}
	}
	return true
}

func level1(in io.Reader) string {
	reports := parse(in)
	safeReports := 0
	for _, report := range reports {
		if isSafe(report) {
			safeReports += 1
		}
	}
	return fmt.Sprint(safeReports)
}

func level2(in io.Reader) string {
	return ""
}

func main() {
	if util.GetLevelArg() == 1 {
		fmt.Println(level1(os.Stdin))
	} else {
		fmt.Println(level2(os.Stdin))
	}
}
