package main

import (
	"bytes"
	"fmt"
	"io"
	"time"

	"github.com/dchiquito/advent-of-code-2024/internal/run"
	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

type Timer struct {
	labels  []string
	times   []int64
	elapsed int64
}

func RunTimer(day int, level int, in io.Reader) Timer {
	labels := make([]string, 0, 100)
	times := make([]int64, 0, 100)
	start := time.Now()
	mark := start
	util.StartStopwatch(func(label string) {
		labels = append(labels, label)
		now := time.Now()
		times = append(times, now.Sub(mark).Nanoseconds())
		mark = now
	})
	if level == 1 {
		run.RunPart1(day, in)
	} else {
		run.RunPart2(day, in)
	}
	end := time.Now()
	elapsed := end.Sub(start).Nanoseconds()
	labels = append(labels, "end")
	times = append(times, end.Sub(mark).Nanoseconds())
	return Timer{labels, times, elapsed}
}

func main() {
	day := util.GetDayArg()
	level := util.GetLevelArg()
	file := run.GetInput(day)
	buff, _ := io.ReadAll(file)
	file.Close()

	var totalTime int64 = 0
	timers := make([]Timer, 0, 1000)
	for true {
		in := bytes.NewReader(buff)
		timer := RunTimer(day, level, in)
		timers = append(timers, timer)
		totalTime += timer.elapsed
		if totalTime > 5_000_000_000 {
			break
		}
		if totalTime > 1_000_000_000 && len(timers) > 10 {
			break
		}
	}
	labels := timers[0].labels
	n := int64(len(timers))
	stopwatchSums := make([]int64, n)
	var elapsedSum int64 = 0
	for _, timer := range timers {
		for i, t := range timer.times {
			stopwatchSums[i] += t
		}
		elapsedSum += timer.elapsed
	}
	fmt.Printf("Solved %d times in %v\n", n, fmtNs(elapsedSum))
	if len(labels) > 1 {
		for i, label := range labels {
			fmt.Printf("%v: %v\n", label, fmtNs(stopwatchSums[i]/n))
		}
	}
	fmt.Printf("average time elapsed: %v\n", fmtNs(elapsedSum/n))
}

func fmtNs(ns int64) string {
	unitsIndex := 0
	units := [...]string{"ns", "µs", "ms", "s"}
	for unitsIndex < len(units)-1 && ns >= 10000 {
		ns /= 1000
		unitsIndex += 1
	}
	return fmt.Sprintf("%d%v", ns, units[unitsIndex])
}
