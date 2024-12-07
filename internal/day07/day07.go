package day07

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func parse(in io.Reader) [][]int {
	equations := make([][]int, 0, 1000)
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")
		equation := make([]int, len(tokens))
		equation[0], _ = strconv.Atoi(tokens[0][:len(tokens[0])-1])
		for i := 1; i < len(tokens); i += 1 {
			equation[i], _ = strconv.Atoi(tokens[i])
		}
		equations = append(equations, equation)
	}
	return equations
}

func solve1(desiredTotal int, soFar int, eq []int) bool {
	// fmt.Println(desiredTotal, soFar, eq)
	if len(eq) == 0 {
		return soFar == desiredTotal
	}
	// TODO assess this optimization with a benchmark
	if soFar > desiredTotal {
		return false
	}
	if solve1(desiredTotal, soFar+eq[0], eq[1:]) {
		return true
	}
	if solve1(desiredTotal, soFar*eq[0], eq[1:]) {
		return true
	}
	return false
}

func canSolve1(equation []int) bool {
	// TODO how much worse is bit twiddling?
	// limit := 1 << (len(equation) - 1)
	// for bits:=0; bits<limit; bits += 1 {}
	// The first value in the equation is the desired total, trim it off
	return solve1(equation[0], equation[1], equation[2:])
}

func Level1(in io.Reader) string {
	equations := parse(in)
	total := 0
	for _, equation := range equations {
		if canSolve1(equation) {
			total += equation[0]
		}
	}
	return fmt.Sprint(total)
}

func solve2(desiredTotal int, soFar int, eq []int) bool {
	if len(eq) == 0 {
		return soFar == desiredTotal
	}
	// TODO assess this optimization with a benchmark
	if soFar > desiredTotal {
		return false
	}
	if solve2(desiredTotal, soFar+eq[0], eq[1:]) {
		return true
	}
	if solve2(desiredTotal, soFar*eq[0], eq[1:]) {
		return true
	}
	// There are no terms >= 1000
	if eq[0] < 10 {
		if solve2(desiredTotal, (soFar*10)+eq[0], eq[1:]) {
			return true
		}
	} else if eq[0] < 100 {
		if solve2(desiredTotal, (soFar*100)+eq[0], eq[1:]) {
			return true
		}
	} else {
		if solve2(desiredTotal, (soFar*1000)+eq[0], eq[1:]) {
			return true
		}
	}
	return false
}

func canSolve2(equation []int) bool {
	return solve2(equation[0], equation[1], equation[2:])
}

func Level2(in io.Reader) string {
	equations := parse(in)
	total := 0
	for _, equation := range equations {
		if canSolve2(equation) {
			total += equation[0]
		}
	}
	return fmt.Sprint(total)
}
