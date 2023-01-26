package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type StoryArc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

func parseJson(fileName string) map[string]StoryArc {

	// Open our jsonFile
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	var parsedData map[string]StoryArc

	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Loads the data
	errUnmarshal := json.Unmarshal([]byte(byteValue), &parsedData)
	if errUnmarshal != nil {
		fmt.Println(err)
	}
	// Printing to test that this works
	// fmt.Printf("%v", parsedData)

	return parsedData
}

func main() {
	startingArc := "intro"
	data := parseJson("gopher.json")
	generateTemplate(startingArc, data[startingArc])

	// http.ListenAndServe(":8080", mapHandler)
}
