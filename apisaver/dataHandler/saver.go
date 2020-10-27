package dataHandler

import (
	"encoding/json"
	"fmt"
	"github.com/Armingodiz/go-stuff/apisaver/encryptor"
	"io/ioutil"
)

func SaveData(name, api, key string) {
	encrypted := encryptor.Encrypt(api, key)
	//fmt.Printf("encrypted : %s\n", encrypted)
	apiSecret := ApiSecret{
		Name:   name,
		Secret: encrypted,
	}
	ApiSecrets.Secrets = append(ApiSecrets.Secrets, apiSecret)
	data := Secrets{
		Secrets: ApiSecrets.Secrets,
	}
	jsonString, _ := json.Marshal(data)
	err := ioutil.WriteFile("secrets.json", jsonString, 0644)
	if err != nil {
		fmt.Println("error in writing on json !")
	}
}
