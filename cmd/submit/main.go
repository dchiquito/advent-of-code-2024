package main

import (
	"fmt"
	"io"
	"os"

	"github.com/dchiquito/advent-of-code-2024/internal/submit"
	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

func readAnswerFromStdin() string {
	answer, err := io.ReadAll(os.Stdin)
	util.Check(err, "Failed to read answer from stdin")
	return string(answer)
}

func main() {
	day := util.GetDayArg()
	level := util.GetLevelArg()
	answer := readAnswerFromStdin()
	fmt.Println("Submitting", answer)
	submit.SendAnswer(day, level, answer)
}
