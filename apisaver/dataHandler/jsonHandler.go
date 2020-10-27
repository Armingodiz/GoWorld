package dataHandler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var ApiSecrets *Secrets

type ApiSecret struct {
	Name   string
	Secret string
}

type Secrets struct {
	Secrets []ApiSecret `json:"apiSecrets"`
}

func init() {
	ApiSecrets = readApiSecrets()
}
func readApiSecrets() *Secrets {
	jsonFile, err := os.Open("secrets.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened secrets.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// we initialize our Users array
	ApiSecrets2 := &Secrets{}
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, ApiSecrets2)
	if ApiSecrets2.Secrets == nil {
		fmt.Println("#################################################################33")
		ApiSecrets2 = &Secrets{
			Secrets: []ApiSecret{},
		}
	}
	fmt.Println(ApiSecrets2.Secrets)
	return ApiSecrets2
}
