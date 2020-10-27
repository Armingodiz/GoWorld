package dataHandler

import (
	"fmt"
	encryptor "github.com/Armingodiz/go-stuff/apisaver/encryptor"
)

func LoadApi(name, key string) string {
	encrypted := " " //this value should be read from json
	decrypted := encryptor.Decrypt(encrypted, key)
	fmt.Printf("decrypted : %s\n", decrypted)
	return decrypted
}
