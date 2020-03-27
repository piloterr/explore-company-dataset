package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// GetURLContent returns the content of a file in the URL
func GetURLContent(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(data)
}

// ParseData takes the json data and parse it as an array
func ParseData(data string) string {
	lines := strings.Split(data, "\n")
	return "[" + strings.Join(lines, ", ") + "]"
}