package main

import (
	"fmt"
	"github.com/dchiquito/advent-of-code-2024/internal/util"
	"io"
	"net/http"
	"os"
	"strconv"
)

func getPuzzleInput(day int) []byte {
	url := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)
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
	if len(os.Args) != 2 {
		fmt.Println("Please specify the day")
		os.Exit(1)
	}
	fmt.Println("Hello World", os.Args[1])
	day, err := strconv.Atoi(os.Args[1])
	util.Check(err, "Please specify a number for the day")
	if day < 1 || day > 25 {
		util.Check("naughty", "Please specify a day between 1 and 25")
	}
	pull(day)
}
