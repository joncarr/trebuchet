package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	appName  = "trebuchet"
	bornDate = 2018
	err      error
)

var rootCmd = &cobra.Command{
	Use:  "trebuchet [command]",
	Long: renderLongDesc(),
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

func renderLongDesc() string {
	var copyright string
	appName := strings.Title(appName)

	if time.Now().Year() == bornDate {
		copyright = fmt.Sprintf("(c) %d Jade Cricket Design Co.", bornDate)
	} else {
		copyright = fmt.Sprintf("(c) %d - %d Jade Cricket Design Co.", bornDate, time.Now().Year())
	}

	return "\n" + appName + " is an open-source, rapid development toolset for creating headless content \n" +
		"management systems and is heavily inspired by the likes of Ponzu CMS, Buffalo, \n" +
		"and Fragmenta CMS. " + appName + " is released under the BSD-3-Clause license.\n" + copyright

}
