package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Rankings struct {
	Users []Ranking `json:"rankings"`
}
type Ranking struct {
	Name      string
	ApiSecret string
}

// an array of users

func main() {
	// writng on json
	ranks := []Ranking{}
	ranks = append(ranks, Ranking{
		Name:      "armin",
		ApiSecret: "gooooooooooooooooooooodiz",
	})
	ranks = append(ranks, Ranking{
		Name:      "armin2",
		ApiSecret: "gooooooooooooooooooooodiz2",
	})
	datas := Rankings{
		Users: ranks,
	}
	jsonString, _ := json.Marshal(datas)
	err := ioutil.WriteFile("test.json", jsonString, 0644)
	if err != nil {
		fmt.Println("errrrrrrrrrrrrrrrrrrrr")
	}

	// reading on json/*

	jsonFile, err := os.Open("test.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var rankings Rankings

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &rankings)
	fmt.Println(rankings.Users)
}
