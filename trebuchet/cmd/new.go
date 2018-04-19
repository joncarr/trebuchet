package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	// ProjectName is the name of the name supplied from the CLI tool
	ProjectName string

	// HTML stores the value of the HTML flag
	HTML bool

	// Vue stores the value of the vue flag
	Vue bool
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

		// Get the current working directory the 'new' command was executed in
		// this path MUST be in users $GOPATH
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		env := os.Getenv("GOPATH")

		fmt.Println("Working Directory:", wd)
		fmt.Println("GOPATH:", env)

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
