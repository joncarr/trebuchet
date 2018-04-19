package app

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joncarr/trebuchet/app/configuration"
	"github.com/joncarr/trebuchet/app/webportal"
)

func main() {
	file, err := os.Open("configuration/config.json")
	if err != nil {
		// TODO: Handle gracefully
		log.Fatal(err)
	}

	config := new(configuration.Configuration)
	json.NewDecoder(file).Decode(config)

	// Checks the config file to see if we are developing or ready for production
	// then sets the appropriate value in the server variable.
	// Good for the time being but I'm not sure how well this will work if planning
	// to host on something like Heroku, where you can just set the env variable
	// TODO: Address this ^^^
	var server string
	switch config.Env {
	case "development":
		server = fmt.Sprintf("%s:%s", config.Dev.AppDomain, config.Dev.AppPort)
	case "production":
		// TODO: Finalize the production server info...
		server = fmt.Sprintf("%s:%s", config.Prod.AppDomain, config.Prod.AppPort)
	}

	// head over to webportal/webportal.go to see what's happening here.
	webportal.RunWebPortal(server)
}
