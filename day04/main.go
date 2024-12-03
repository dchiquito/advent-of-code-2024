package main

import (
	"fmt"
	"io"
	"os"

	"github.com/dchiquito/advent-of-code-2024/internal/util"
)

func parse(in io.Reader) {}

func level1(in io.Reader) string {
	return ""
}

func level2(in io.Reader) string {
	return ""
}

func main() {
	if util.GetLevelArg() == 1 {
		fmt.Println(level1(os.Stdin))
	} else {
		fmt.Println(level2(os.Stdin))
	}
}
