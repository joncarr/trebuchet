package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v", "ver"},
	Short:   "Print the version number of Trebuchet",
	Long:    `All software has versions. This is Trebuchet's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Trebuchet v0.1.0 -- HEAD")
	},
}
