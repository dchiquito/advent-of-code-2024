package util

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
)

func Check(err any, msg string) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
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
	Check(err, "Failed to get the input")
	return resp
}
