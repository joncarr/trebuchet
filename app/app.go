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

	var server string
	switch config.Env {
	case "development":
		server = fmt.Sprintf("%s:%s", config.Dev.AppDomain, config.Dev.AppPort)
	case "production":
		// TODO: Finalize the production server info...
		server = fmt.Sprintf("%s:%s", config.Prod.AppDomain, config.Prod.AppPort)
	}

	webportal.RunWebPortal(server)
}
