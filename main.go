package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	datasetUrl := flag.String("datasetUrl", "", "The url of the dataset.")
	meiliUrl := flag.String("meiliUrl", "", "The url of the MeiliSearch API.")
	flag.Parse()

	// Getting the secret key from the environment variables
	secretKey := os.Getenv("MEILI_SECRET_KEY")
	if secretKey == "" {
		panic("The secret key is missing")
	}

	// Checking if the URLs are here
	if *datasetUrl == "" || *meiliUrl == "" {
		panic("The -datasetUrl or -meiliUrl flag must be present.")
	}

	fmt.Println("Getting dataset content")
	data := GetURLContent(*datasetUrl)
	fmt.Println("Formatting the data")
	ParseData(*meiliUrl, secretKey, data)
}