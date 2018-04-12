package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	ProjectName string
	HTML        bool
	Vue         bool
)

var newCmd = &cobra.Command{
	Use:   "new [flags] [project name]",
	Short: "Creates a project directory for the supplied project name",
	Long: "\n\t" + `Creates a project directory for the supplied project name 
	following the 'new' command (if no flags are provided) in your current 
	working directory. Your current working directory MUST be within
	your $GOPATH`,
	Example: "\n$ trebuchet new project_name\n" +
		"\t - or - \n" +
		"$ trebuchet new --vue project_name\n" +
		"... project generation executing...\n" +
		"\n" +
		"> Your new project can be found at $GOPATH/src/github.com/[you]/project_name",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			ProjectName = args[0]
		} else {
			msg := "Please provide a project name."
			msg += "\nThis will create a directory within your $GOPATH/src."
			return fmt.Errorf("%s", msg)
		}
		// return newProjectInDir(projectName)
		return nil
	},
}

func init() {
	// newCmd.Flags().StringVar(&fork, "fork", "", "modify repo source for Ponzu core development")
	newCmd.Flags().BoolVar(&Vue, "vue", false, "Generate a Vue.js front-end for this CMS")
	newCmd.Flags().BoolVar(&HTML, "html", false, "Generate a HTML/CSS front-end for this CMS")
	// newCmd.Flags().BoolVar(&Vue, "vue", false, "Generate a Vue.js head for this cms")

	rootCmd.AddCommand(newCmd)
}
