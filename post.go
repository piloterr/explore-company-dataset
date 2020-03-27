package main

import (
	"bytes"
	"net/http"
)

// PostData posts the companies data to the MeiliSearch API
func PostData(meiliUrl, payload, secretKey string) {
	req, err := http.NewRequest(
		"POST",
		meiliUrl,
		bytes.NewBuffer([]byte(payload)),
	)
	req.Header.Set("X-Meili-API-Key", secretKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
}