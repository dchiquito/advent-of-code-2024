package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"

	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

func parse(in io.Reader) {}

func level1(in io.Reader) string {
	input, _ := io.ReadAll(in)
	re := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\))|`)
	muls := re.FindAllStringSubmatch(string(input), -1)
	total := 0
	for _, mul := range muls {
		left, _ := strconv.Atoi(mul[1])
		right, _ := strconv.Atoi(mul[2])
		total += left * right
	}
	return fmt.Sprint(total)
}

func level2(in io.Reader) string {
	input, _ := io.ReadAll(in)
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)`)
	ops := re.FindAllStringSubmatch(string(input), -1)
	total := 0
	do := true
	for _, op := range ops {
		if op[0] == "do()" {
			do = true
		} else if op[0] == "don't()" {
			do = false
		} else if do {
			left, _ := strconv.Atoi(op[1])
			right, _ := strconv.Atoi(op[2])
			total += left * right
		}
	}
	return fmt.Sprint(total)
}

func main() {
	if util.GetLevelArg() == 1 {
		fmt.Println(level1(os.Stdin))
	} else {
		fmt.Println(level2(os.Stdin))
	}
}
