package playground

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/joncarr/trebuchet/app/configuration"
)

type animal struct {
	gorm.Model
	AnimalType string `gorm:"type:VARCHAR(36)"`
	Nickname   string `gorm:"type:VARCHAR(36);UNIQUE_INDEX"`
	Zone       int    `gorm:"type:INT"`
	Age        int    `gorm:"type:INT"`
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

	}

	//*****************************************************************************************************************************************************
	//* Make the connection to the DB
	db, err := gorm.Open(config.DBType, connString)
	if err != nil {
		fmt.Println("DB Connection Error: Connection was not made to the database")
		log.Fatal(err)
	}
	defer db.Close()

	db.DropTableIfExists(&animal{})
	db.AutoMigrate(&animal{})

	a := animal{
		AnimalType: "Tyrannosaurus Rex",
		Nickname:   "Biff",
		Zone:       1,
		Age:        43,
	}

	db.Save(&a)

	// db.Table("animals").Where("nickname = ? and zone = ?", "Biff", 1).Update("age", 29)

	animals := []animal{}
	db.Find(&animals, "age > ?", 12)
	fmt.Println(animals)

}
