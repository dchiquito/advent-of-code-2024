package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strconv"
	"strings"
)

func _check(err any, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}

func getSessionCookie() string {
	content, err := os.ReadFile(".session")
	_check(err, "Please log in to https://adventofcode.com/2024, grab the session cookie, and save it in a .session file")
	return strings.TrimSpace(string(content))
}

func getPuzzleInput(day int) []byte {
	url := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)
	cookie := &http.Cookie{
		Name:  "session",
		Value: getSessionCookie(),
	}

	jar, err := cookiejar.New(&cookiejar.Options{})
	_check(err, "Failed to create cookie jar")

	client := &http.Client{
		Jar: jar,
	}
	req, err := http.NewRequest("GET", url, nil)
	_check(err, "Failed to set up request")

	client.Jar.SetCookies(req.URL, []*http.Cookie{cookie})

	resp, err := client.Do(req)
	_check(err, "Failed to get the input")
	defer resp.Body.Close()

	input, err := io.ReadAll(resp.Body)
	_check(err, "Failed to read response when fetching endpoint")

	return input
}

func writePuzzleInput(day int, input []byte) {
	filename := fmt.Sprintf("data/%02d.txt", day)
	err := os.WriteFile(filename, input, 0666)
	_check(err, "Failed to write input to file")
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
	_check(err, "Please specify a number for the day")
	if day < 1 || day > 25 {
		_check("naughty", "Please specify a day between 1 and 25")
	}
	pull(day)
}
