package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var Set = &cobra.Command{
	Use:   "set",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}

func init() {
	RootCmd.AddCommand(Set)
}
