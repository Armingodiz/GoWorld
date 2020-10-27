package dataHandler

import (
	"fmt"
	"github.com/Armingodiz/go-stuff/apisaver/encryptor"
)

func SaveData(name , api , key string){
	encrypted := encryptor.Encrypt("Hello Encrypt", key)
	fmt.Printf("encrypted : %s\n", encrypted)
}
