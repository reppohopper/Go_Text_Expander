package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"log"
)

type Abbreviation struct {
	Abbreviation string `json:"abbreviation"`
	Expansion    string `json:"expansion"`
}

var ExpansionsMap map[string][]rune

func init() {
	filePath := os.Getenv("EXPANSION_FILE_PATH")
    if filePath == "" {
        log.Fatal("EXPANSION_FILE_PATH is not set. Please create an " + 
	        "environment variable called 'EXPANSION_FILE_PATH' that stores " + 
		    "the path to your expansions file")
    }

	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var abbreviations []Abbreviation
	err = json.Unmarshal(byteValue, &abbreviations)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	ExpansionsMap = make(map[string][]rune)
	for _, abbr := range abbreviations {
		ExpansionsMap[abbr.Abbreviation] = []rune(abbr.Expansion)
	}

	// Print the map.
	fmt.Println("Successfully aquired expansion map. ")
}
