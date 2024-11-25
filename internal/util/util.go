package util

import (
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

func GetDayArg() int {
	if len(os.Args) < 2 {
		Panic("Please specify the day")
	}
	day, err := strconv.Atoi(os.Args[1])
	Check(err, "Please specify a number for the day")
	if day < 1 || day > 25 {
		Check("naughty", "Please specify a day between 1 and 25")
	}
	return day
}

func GetLevelArg() int {
	if len(os.Args) < 3 {
		Panic("Please specify the level")
	}
	level, err := strconv.Atoi(os.Args[2])
	Check(err, "Please specify a number for the level")
	if level != 1 && level != 2 {
		Panic("Please specify either level 1 or 2")
	}
	return level
}
