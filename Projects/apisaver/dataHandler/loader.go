package dataHandler

import (
	"fmt"
	encryptor "github.com/Armingodiz/go-stuff/apisaver/encryptor"
)

func LoadApi(name, key string) string {
	encrypted, ok := getEncrypted(name)
	if !ok {
		fmt.Println("	error in loading api key !!")
		return ""
	} else {
		decrypted := encryptor.Decrypt(encrypted, key)
		fmt.Printf("	decrypted : %s\n", decrypted)
		return decrypted
	}
}

func getEncrypted(name string) (string, bool) {
	for _, api := range ApiSecrets.Secrets {
		if api.Name == name {
			return api.Secret, true
		}
	}
	return "", false
}
