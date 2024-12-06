package main

import (
	"fmt"
	"io"
	"os"

	"github.com/dchiquito/advent-of-code-2024/internal/day01"
	"github.com/dchiquito/advent-of-code-2024/internal/day02"
	"github.com/dchiquito/advent-of-code-2024/internal/day03"
	"github.com/dchiquito/advent-of-code-2024/internal/day04"
	"github.com/dchiquito/advent-of-code-2024/internal/day05"
	"github.com/dchiquito/advent-of-code-2024/internal/day06"
	"github.com/dchiquito/advent-of-code-2024/internal/day07"
	"github.com/dchiquito/advent-of-code-2024/internal/day08"
	"github.com/dchiquito/advent-of-code-2024/internal/day09"
	"github.com/dchiquito/advent-of-code-2024/internal/day10"
	"github.com/dchiquito/advent-of-code-2024/internal/day11"
	"github.com/dchiquito/advent-of-code-2024/internal/day12"
	"github.com/dchiquito/advent-of-code-2024/internal/day13"
	"github.com/dchiquito/advent-of-code-2024/internal/day14"
	"github.com/dchiquito/advent-of-code-2024/internal/day15"
	"github.com/dchiquito/advent-of-code-2024/internal/day16"
	"github.com/dchiquito/advent-of-code-2024/internal/day17"
	"github.com/dchiquito/advent-of-code-2024/internal/day18"
	"github.com/dchiquito/advent-of-code-2024/internal/day19"
	"github.com/dchiquito/advent-of-code-2024/internal/day20"
	"github.com/dchiquito/advent-of-code-2024/internal/day21"
	"github.com/dchiquito/advent-of-code-2024/internal/day22"
	"github.com/dchiquito/advent-of-code-2024/internal/day23"
	"github.com/dchiquito/advent-of-code-2024/internal/day24"
	"github.com/dchiquito/advent-of-code-2024/internal/day25"
	"github.com/dchiquito/advent-of-code-2024/internal/pull"
	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

func getInput(day int) *os.File {
	path := fmt.Sprintf("data/%02d.txt", day)
	in, _ := os.Open(path)
	if in == nil {
		// Assume the input has not yet been pulled, try to pull it
		pull.Pull(day)
		in, _ = os.Open(path)
		if in == nil {
			util.Panic("Failed to fetch input for day", 6)
		}
	}
	return in
}

func RunPart1(day int, in io.Reader) string {
	switch day {
	case 1:
		return day01.Level1(in)
	case 2:
		return day02.Level1(in)
	case 3:
		return day03.Level1(in)
	case 4:
		return day04.Level1(in)
	case 5:
		return day05.Level1(in)
	case 6:
		return day06.Level1(in)
	case 7:
		return day07.Level1(in)
	case 8:
		return day08.Level1(in)
	case 9:
		return day09.Level1(in)
	case 10:
		return day10.Level1(in)
	case 11:
		return day11.Level1(in)
	case 12:
		return day12.Level1(in)
	case 13:
		return day13.Level1(in)
	case 14:
		return day14.Level1(in)
	case 15:
		return day15.Level1(in)
	case 16:
		return day16.Level1(in)
	case 17:
		return day17.Level1(in)
	case 18:
		return day18.Level1(in)
	case 19:
		return day19.Level1(in)
	case 20:
		return day20.Level1(in)
	case 21:
		return day21.Level1(in)
	case 22:
		return day22.Level1(in)
	case 23:
		return day23.Level1(in)
	case 24:
		return day24.Level1(in)
	case 25:
		return day25.Level1(in)
	}
	util.Panic("Invalid day", day)
	return ""
}

func RunPart2(day int, in io.Reader) string {
	switch day {
	case 1:
		return day01.Level2(in)
	case 2:
		return day02.Level2(in)
	case 3:
		return day03.Level2(in)
	case 4:
		return day04.Level2(in)
	case 5:
		return day05.Level2(in)
	case 6:
		return day06.Level2(in)
	case 7:
		return day07.Level2(in)
	case 8:
		return day08.Level2(in)
	case 9:
		return day09.Level2(in)
	case 10:
		return day10.Level2(in)
	case 11:
		return day11.Level2(in)
	case 12:
		return day12.Level2(in)
	case 13:
		return day13.Level2(in)
	case 14:
		return day14.Level2(in)
	case 15:
		return day15.Level2(in)
	case 16:
		return day16.Level2(in)
	case 17:
		return day17.Level2(in)
	case 18:
		return day18.Level2(in)
	case 19:
		return day19.Level2(in)
	case 20:
		return day20.Level2(in)
	case 21:
		return day21.Level2(in)
	case 22:
		return day22.Level2(in)
	case 23:
		return day23.Level2(in)
	case 24:
		return day24.Level2(in)
	case 25:
		return day25.Level2(in)
	}
	util.Panic("Invalid day", day)
	return ""
}

func main() {
	day := util.GetDayArg()
	level := util.GetLevelArg()
	in := getInput(day)
	defer in.Close()
	var solution string
	if level == 1 {
		solution = RunPart1(day, in)
	} else {
		solution = RunPart2(day, in)
	}
	fmt.Println(solution)
	fmt.Println("Submit this?")
}
