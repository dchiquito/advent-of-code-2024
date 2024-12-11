package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
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

func fmtNs(ns float64) string {
	unitsIndex := 0
	units := [...]string{"ns", "µs", "ms", "s"}
	for unitsIndex < len(units)-1 && ns >= 1000 {
		ns /= 1000
		unitsIndex += 1
	}
	return fmt.Sprintf("%0.3f%v", ns, units[unitsIndex])
}
func calculateMean(arr []float64) float64 {
	var total float64 = 0
	for _, x := range arr {
		total += float64(x)
	}
	return total / float64(len(arr))
}
func calculateVariance(mean float64, arr []float64) float64 {
	var total float64 = 0.0
	for _, x := range arr {
		deviation := mean - x
		total += deviation * deviation
	}
	return total / float64(len(arr))
}
func calculateStdDev(mean float64, arr []float64) float64 {
	return math.Sqrt(float64(calculateVariance(mean, arr)))
}
func calculateStdErr(mean float64, arr []float64) float64 {
	return calculateStdDev(mean, arr) / math.Sqrt(float64(len(arr)))
}
func printTimingInfo(timers []Timer) {
	labels := timers[0].labels

	// Invert the timers matrix
	times := make([][]float64, len(labels))
	for i := range times {
		times[i] = make([]float64, len(timers))
		for j, timer := range timers {
			times[i][j] = float64(timer.times[i])
		}
	}
	elapseds := make([]float64, len(timers))
	var elapsedSum float64 = 0
	for i, timer := range timers {
		elapseds[i] = float64(timer.elapsed)
		elapsedSum += elapseds[i]
	}

	means := make([]float64, len(times))
	for i := range means {
		means[i] = calculateMean(times[i])
	}
	elapsedMean := calculateMean(elapseds)

	// stdDevs := make([]float64, len(times))
	// for i := range stdDevs {
	// 	stdDevs[i] = calculateStdDev(means[i], times[i])
	// }
	// elapsedStdDev := calculateStdDev(elapsedMean, elapseds)

	stdErrs := make([]float64, len(times))
	for i := range stdErrs {
		stdErrs[i] = calculateStdErr(means[i], times[i])
	}
	elapsedStdErr := calculateStdErr(elapsedMean, elapseds)

	fmt.Printf("Solved %d times in %v\n", len(timers), fmtNs(elapsedSum))
	if len(labels) > 1 {
		maxLabelLen := 0
		for _, label := range labels {
			if len(label) > maxLabelLen {
				maxLabelLen = len(label)
			}
		}
		for i, label := range labels {
			fmt.Printf("  %[1]*v: %v (±%v)\n", maxLabelLen, label, fmtNs(means[i]), fmtNs(stdErrs[i]))
		}
	}
	fmt.Printf("average time elapsed: %v (±%v)\n", fmtNs(elapsedMean), fmtNs(elapsedStdErr))
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
	printTimingInfo(timers)
}
