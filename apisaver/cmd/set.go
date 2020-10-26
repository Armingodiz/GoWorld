package cmd

import (
	"fmt"

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
    fi()
	},
}

func init() {
	RootCmd.AddCommand(setCmd)
}
func fi(){
  fmt.Println("fuck me")
}
