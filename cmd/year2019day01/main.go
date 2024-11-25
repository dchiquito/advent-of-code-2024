package main

// This is the day 1 solution for 2019, which I solved to test the API early before starting 2024.

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

func level1() {
	scanner := bufio.NewScanner(os.Stdin)
	total_fuel := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		util.Check(err, "malformed input line")
		total_fuel += (mass / 3) - 2
	}
	fmt.Println(total_fuel)
}

func fuel_from_mass(mass int) int {
	fuel := (mass / 3) - 2
	if fuel <= 0 {
		return 0
	} else {
		return fuel + fuel_from_mass(fuel)
	}
}

func level2() {
	scanner := bufio.NewScanner(os.Stdin)
	total_fuel := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		util.Check(err, "malformed input line")
		total_fuel += fuel_from_mass(mass)
	}
	fmt.Println(total_fuel)
}

func main() {
	if util.GetLevelArg() == 1 {
		level1()
	} else {
		level2()
	}
}
