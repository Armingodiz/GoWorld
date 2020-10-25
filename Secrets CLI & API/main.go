package main

import (
	"encoding/hex"
	"fmt"
	"gi"
)

func main() {
	bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	key := hex.EncodeToString(bytes) //encode key in bytes to string and keep as secret, put in a vault
	fmt.Printf("key to encrypt/decrypt : %s\n", key)

	encrypted :=  encrypt("Hello Encrypt", key)
	fmt.Printf("encrypted : %s\n", encrypted)

	decrypted := decrypt(encrypted, key)
	fmt.Printf("decrypted : %s\n", decrypted)
}
