package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

func getPuzzleInput(day int) []byte {
	url := fmt.Sprintf("https://adventofcode.com/2019/day/%d/input", day)
	req, err := http.NewRequest("GET", url, nil)
	util.Check(err, "Failed to set up request")

	resp := util.SendRequest(req)
	defer resp.Body.Close()
	input, err := io.ReadAll(resp.Body)
	util.Check(err, "Failed to read response when fetching endpoint")

	return input
}

func writePuzzleInput(day int, input []byte) {
	filename := fmt.Sprintf("data/%02d.txt", day)
	err := os.WriteFile(filename, input, 0666)
	util.Check(err, "Failed to write input to file")
}

func pull(day int) {
	input := getPuzzleInput(day)
	writePuzzleInput(day, input)
}

func main() {
	day := util.GetDayArg()
	pull(day)
}
