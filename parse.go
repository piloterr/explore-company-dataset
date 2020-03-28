package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

type CompanyWithID struct {
	CompanyID int `json:"company_id"`
	NewCompany
}

type Company struct {
	Name                    string `json:"name"`
	Domain                  string `json:"domain"`
	YearFounded             string `json:"year founded"`
	Industry                string `json:"industry"`
	SizeRange               string `json:"size range"`
	Locality                string `json:"locality"`
	Country                 string `json:"country"`
	LinkedinUrl             string `json:"linkedin url"`
	CurrentEmployeeEstimate string `json:"current employee estimate"`
	TotalEmployeeEstimate   string `json:"total employee estimate"`
}

type NewCompany Company

// GetURLContent returns the content of a file in the URL
func GetURLContent(url string) string {
	bytes, err := ioutil.ReadFile(url)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

// ParseData takes the json data and parse it as an array
func ParseData(url, key, data string) {
	lines := strings.Split(data, "\n")

	for id, line := range lines {
		fmt.Printf("Parsing the %dth line.\n", id)

		// Serialize the json
		var company Company
		err := json.Unmarshal([]byte(line), &company)
		if err != nil {
			panic(err)
		}

		// Add the company id field to the json
		bytes, err := json.Marshal(CompanyWithID{
			CompanyID:  id,
			NewCompany: NewCompany(company),
		})
		if err != nil {
			panic(err)
		}

		PostData(url, "[" + string(bytes) + "]", key)
	}
}
