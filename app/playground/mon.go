package playground

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joncarr/trebuchet/app/configuration"
	mgo "gopkg.in/mgo.v2"
)

type animal struct {
	ID         int    `json:"id" bson:"id"`
	AnimalType string `json:"animal_type" bson:"animal_type"`
	Nickname   string `json:"nickname" bson:"nickname"`
	Zone       int    `json:"zone" bson:"zone"`
	Age        int    `json:"age" bson:"age"`
}

func main() {

	//*****************************************************************************************************************************************************
	//* Open file holding configuration settings and assign to a variable
	file, err := os.Open("configuration/config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	config := new(configuration.Configuration)
	json.NewDecoder(file).Decode(config)

	//*****************************************************************************************************************************************************
	//* Construct the DB connection string 'connString'

	var connString string

	switch config.DBType {
	case "mysql":
		connString = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", config.Dev.DBUserName, config.Dev.DBPassword, config.Dev.DBName)
	case "postgres":
		connString = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.Dev.DBUserName, config.Dev.DBPassword, config.Dev.DBName)
	case "sqlite":
		connString = fmt.Sprintf("$s.db", config.AppName)
	case "mongo":
		connString = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", config.Dev.DBUserName, config.Dev.DBPassword, config.Dev.MongoConnectionURI, config.Dev.DBPort, config.Dev.DBName)

	}

	//*****************************************************************************************************************************************************
	//* Make the connection to the DB

	if config.DBType == "mongo" {
		db, err := mgo.Dial(connString)
	} else {
		db, err := gorm.Open(config.DBType, connString)
	}
	if err != nil {
		fmt.Println("DB Connection Error: Connection was not made to the database")
		log.Fatal(err)
	}
	defer mongoSession.Close()

	fmt.Println("WE CONNECTED!")

}
