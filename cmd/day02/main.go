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
	ascending := report[0] < report[1]
	for i := 1; i < len(report); i += 1 {
		delta := report[i] - report[i-1]
		if !ascending {
			delta = -delta
		}
		if delta < 1 || delta > 3 {
			return false
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

func isDeltaSafe(a int, b int, ascending bool) bool {
	delta := b - a
	if !ascending {
		delta = -delta
	}
	return delta >= 1 && delta <= 3
}

func isSafeDamped(report []int) bool {
	fmt.Println(report)
	ascending := report[0] < report[1]
	damped := false
	for i := 1; i < len(report); i += 1 {
		if !isDeltaSafe(report[i-1], report[i], ascending) {
			fmt.Println("unsafe", report[i-1], report[i], i)
			if damped {
				return false
			} else {
				// The problematic level may be report[i] or report[i-1]
				// i-2 i-1  i  i+1
				//  a   b   c   d
				if i >= 2 && isDeltaSafe(report[i-2], report[i], ascending) &&
					(i+1 >= len(report) || isDeltaSafe(report[i], report[i+1], ascending)) {
					fmt.Println("removing i-1 solves it")
					// First, check if removing report[i-1] (b) solves the problem
					// (a,c) and (c,d) must both be safe
					damped = true
				} else if i+1 < len(report) && isDeltaSafe(report[i-1], report[i+1], ascending) {
					fmt.Println("removing i solves it")
					// If that didn't work, check if removing report[i] (c) solves the problem
					// The last loop iteration already checked that (a,b) is safe,
					// so we only need to check (b,d)
					damped = true
					// The next step in the loop would check (c,d), but we just removed c, so we replace c with b to simulate removing c from the list
					report[i] = report[i-1]
				} else {
					fmt.Println("neither :(")
					// Netiher removal solved the problem
					return false
				}
			}
		}
	}
	return true
}

func level2wrong(in io.Reader) string {
	reports := parse(in)
	safeReports := 0
	for _, report := range reports {
		if isSafeDamped(report) {
			safeReports += 1
		}
	}
	return fmt.Sprint(safeReports)
}
func level2(in io.Reader) string {
	reports := parse(in)
	safeReports := 0
	for _, report := range reports {
		trimmedReport := make([]int, len(report)-1)
		for i := range report {
			copy(trimmedReport, report[:i])
			copy(trimmedReport[i:], report[i+1:])
			if isSafe(trimmedReport) {
				safeReports += 1
				break
			}
		}
	}
	return fmt.Sprint(safeReports)
}

func main() {
	if util.GetLevelArg() == 1 {
		fmt.Println(level1(os.Stdin))
	} else {
		fmt.Println(level2(os.Stdin))
	}
}
