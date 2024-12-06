package pull

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

func GetPuzzleInput(day int) []byte {
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)
	req, err := http.NewRequest("GET", url, nil)
	util.Check(err, "Failed to set up request")

	resp := util.SendRequest(req)
	defer resp.Body.Close()
	input, err := io.ReadAll(resp.Body)
	util.Check(err, "Failed to read response when fetching endpoint")
	if string(input[:6]) == "Please" {
		util.Panic("Day", day, "has not yet been released")
	}

	return input
}

func WritePuzzleInput(day int, input []byte) {
	filename := util.DefaultInputFilePath(day)
	err := os.WriteFile(filename, input, 0666)
	util.Check(err, "Failed to write input to file")
}

func Pull(day int) {
	input := GetPuzzleInput(day)
	WritePuzzleInput(day, input)
}
