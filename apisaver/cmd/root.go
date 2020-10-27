package cmd

// after using this pattern you must run go install . in your dir to install your new cli , with every change you make
// you must run go install . again
import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "apisaver",
	Short: "Task is a CLI task manager",
}

func init() {
	RootCmd.PersistentFlags().StringVar(&key, "key", "", "this key will be used to encrypt your api secrets")
}
