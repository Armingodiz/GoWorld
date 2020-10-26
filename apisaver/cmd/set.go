package cmd

import (
	"fmt"


	"github.com/spf13/cobra"
)
var key string

var Set = &cobra.Command{
	Use:   "set",
	Short: "this Command will save your api secret key in encrypted mode",
	Run: func(cmd *cobra.Command, args []string) {
    //TODO completing set func
		fmt.Println(args)
    fmt.Println("ADDED")
    fmt.Println(key)
	},
}

func init() {
	RootCmd.AddCommand(Set)
  Set.PersistentFlags().StringVar(&key, "config", "", "config file (default is $HOME/.cobra.yaml)")
}
