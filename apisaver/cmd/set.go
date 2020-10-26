package cmd

import (
	"fmt"
  "crypto/rand"
	"encoding/hex"
	encryptor "github.com/Armingodiz/go-stuff/apisaver/encryptor"
	"github.com/spf13/cobra"
)
var key string

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "this Command will save your api secret key in encrypted mode",
	Run: func(cmd *cobra.Command, args []string) {
    //TODO completing set func
    fmt.Print("name of api secret : ")
    fmt.Println(args[0])
    fmt.Print("received api secret :")
    fmt.Println(args[1])
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
func setApi(){
  bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
  if _, err := rand.Read(bytes); err != nil {
    panic(err.Error())
  }
  key := hex.EncodeToString(bytes) //encode key in bytes to string and keep as secret, put in a vault
  fmt.Printf("key to encrypt/decrypt : %s\n", key)
  fmt.Println("warning ! --------------- > save this key , you wont be able to encrypt/decrypt your saved api secret without it !!")

  encrypted := encryptor.Encrypt("Hello Encrypt", key)
  fmt.Printf("encrypted : %s\n", encrypted)
}
