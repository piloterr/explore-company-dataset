package main

import (
	"flag"
	"fmt"
)

func main() {
	url := flag.String("url", "", "The url of the dataset.")
	flag.Parse()

	if *url == "" {
		panic("The -url flag must be present.")
	}

	data := GetURLContent(*url)
	fmt.Println(ParseData(data))
}