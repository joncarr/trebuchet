package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "trebuchet",
	Short: "Trebuchet does some heavy lifting to launch productivity",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute is where it all begins
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
