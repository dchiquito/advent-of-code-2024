package main

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/dchiquito/advent-of-code-2024/internal/run"
	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

func RunTimer(day int, level int, in io.Reader) int64 {
	before := time.Now()
	if level == 1 {
		run.RunPart1(day, in)
	} else {
		run.RunPart2(day, in)
	}
	after := time.Now()
	elapsed := after.Sub(before).Nanoseconds()
	return elapsed
}

func main() {
	day := util.GetDayArg()
	level := util.GetLevelArg()
	file := run.GetInput(day)
	buff, _ := io.ReadAll(file)
	file.Close()

	totalTime := 0
	n := 0
	start := time.Now()
	for true {
		in := bytes.NewReader(buff)
		totalTime += int(RunTimer(day, level, in))
		n += 1
		elapsed := time.Now().Sub(start).Seconds()
		if elapsed > 5.0 {
			break
		}
		if elapsed > 1.0 && n > 10 {
			break
		}
	}
	meanTime := totalTime / n
	unitsIndex := 0
	units := [...]string{"ns", "µs", "ms", "s"}
	for unitsIndex < len(units)-1 && meanTime >= 10000 {
		meanTime /= 1000
		unitsIndex += 1
	}
	fmt.Println(meanTime, units[unitsIndex], "  ", n)
}
