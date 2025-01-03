package util

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Check(err any, msg string) {
	if err != nil {
		Panic(msg)
	}
}
func Panic(a ...any) {
	fmt.Println(a...)
	os.Exit(1)
}

func ToInt(s string) int {
	total := 0
	for i := 0; i < len(s); i += 1 {
		total *= 10
		total += int(s[i]) - 48
	}
	return total
}

// Takes a line of input and a starting index, searches for the next integer, and returns the next index and that integer.
func ChompInt(line []byte, i int) (int, int) {
	for line[i] < 48 || line[i] > 57 {
		i += 1
	}
	val := int(line[i]) - 48
	neg := i-1 >= 0 && line[i-1] == 45
	i += 1
	for i < len(line) && line[i] >= 48 && line[i] <= 57 {
		val = (val * 10) + int(line[i]) - 48
		i += 1
	}
	if neg {
		val = -val
	}
	return i, val
}

func GetSessionCookie() *http.Cookie {
	content, err := os.ReadFile(".session")
	Check(err, "Please log in to https://adventofcode.com, grab the session cookie, and save it in a .session file")
	value := strings.TrimSpace(string(content))
	cookie := &http.Cookie{
		Name:  "session",
		Value: value,
	}
	return cookie
}

func Client(url *url.URL, cookie *http.Cookie) *http.Client {
	jar, err := cookiejar.New(&cookiejar.Options{})
	Check(err, "Failed to create cookie jar")
	jar.SetCookies(url, []*http.Cookie{cookie})

	client := &http.Client{Jar: jar}
	return client
}

func SendRequest(req *http.Request) *http.Response {
	client := Client(req.URL, GetSessionCookie())
	resp, err := client.Do(req)
	Check(err, "Failed to send request")
	return resp
}

var argInput string

func prepFlags() {
	if flag.Parsed() {
		return
	}

	flag.StringVar(&argInput, "input", "", "foo bar")
	flag.Parse()
}

func GetDayArg() int {
	prepFlags()
	if len(flag.Args()) < 1 {
		Panic("Please specify the day")
	}
	day, err := strconv.Atoi(flag.Arg(0))
	Check(err, "Please specify a number for the day")
	if day < 1 || day > 25 {
		Check("naughty", "Please specify a day between 1 and 25")
	}
	return day
}

func GetLevelArg() int {
	prepFlags()
	if len(flag.Args()) < 2 {
		Panic("Please specify the level")
	}
	level, err := strconv.Atoi(flag.Arg(1))
	Check(err, "Please specify a number for the level")
	if level != 1 && level != 2 {
		Panic("Please specify either level 1 or 2")
	}
	return level
}

func DefaultInputFilePath(day int) string {
	return fmt.Sprintf("data/%02d.txt", day)
}

func GetInputFilePath(day int) string {
	prepFlags()
	if HasInputArg() {
		return argInput
	} else {
		return DefaultInputFilePath(day)
	}
}
func HasInputArg() bool {
	return argInput != ""
}

var stopwatch func(string)

func StartStopwatch(f func(string)) {
	stopwatch = f
}
func Stopwatch(label string) {
	if stopwatch != nil {
		stopwatch(label)
	}
}
