package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dchiquito/advent-of-code-2024/internal/run"
	"github.com/dchiquito/advent-of-code-2024/internal/submit"
	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

func main() {
	day := util.GetDayArg()
	level := util.GetLevelArg()
	in := run.GetInput(day)
	defer in.Close()
	var solution string
	if level == 1 {
		solution = run.RunPart1(day, in)
	} else {
		solution = run.RunPart2(day, in)
	}
	fmt.Println(solution)
	// Assume any 0,1, or 2 digit solutions are wrong
	if len(solution) < 3 || util.HasInputArg() {
		return
	}
	fmt.Print("Submit this? (y/n): ")
	reader := bufio.NewReader(os.Stdin)
	yesNo, _ := reader.ReadString('\n')
	if yesNo == "y\n" {
		submit.SendAnswer(day, level, solution)
	}
}
