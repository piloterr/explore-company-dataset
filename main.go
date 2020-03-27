package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := flag.String("url", "", "The url of the dataset.")
	flag.Parse()

	if *url == "" {
		panic("The -url flag must be present.")
	}

	data := GetURLContent(*url)

	fmt.Println(data)
}

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
