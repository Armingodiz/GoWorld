package dataHandler

import (
	//"github.com/Armingodiz/go-stuff/apisaver/encryptor"
	"encoding/json"
"fmt"
"io/ioutil"
)

func SaveData(name, api, key string) {
//	encrypted := encryptor.Encrypt("Hello Encrypt", key)
//	fmt.Printf("encrypted : %s\n", encrypted)

	apiSecret := ApiSecret{
		Name: name,
		Secret:  api,
	}
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!1")
	fmt.Println(ApiSecrets.Secrets)
	fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!1")
	ApiSecrets.Secrets = append(ApiSecrets.Secrets, apiSecret)
	fmt.Println(ApiSecrets.Secrets)
	fmt.Println("11111111111111111111111111111")
	data := Secrets{
		Secrets: ApiSecrets.Secrets,
	}
	fmt.Println(data)
	jsonString, _ := json.Marshal(data)
	err := ioutil.WriteFile("secrets.json", jsonString, 0644)
	if err != nil {
		fmt.Println("errrrrrrrrrrrrrrrrrrrr")
	}
	fmt.Println("yessssssssssssssss")
}
