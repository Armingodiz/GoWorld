package cmd

import (
	"fmt"

	"github.com/Armingodiz/go-stuff/apisaver/dataHandler"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "this command will get you decryoted api secret with inputted key and name",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("	received key :")
		fmt.Println(key)
		fmt.Print("	received name :  ")
		fmt.Println(args[0])
		getApi(args[0], key)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getApi(name, key string) string {
	apiSecret := dataHandler.LoadApi(name, key)
	return apiSecret
}
