package configuration

// Package configuration defines the structures used to hold
// the values stored inside config.json
//
// Todo: Finish writing doc for configuration

// Configuration holds the web app configuration settings
type Configuration struct {
	AppName string     `json:"app_name"`
	Env     string     `json:"environment"`
	DBType  string     `json:"db_type"`
	Dev     DevConfig  `json:"dev_config"`
	Prod    ProdConfig `json:"prod_config"`
}

// DevConfig holds the values for development environments
type DevConfig struct {
	AppDomain          string `json:"app_domain"`
	AppPort            string `json:"app_port"`
	DBUserName         string `json:"db_username"`
	DBPassword         string `json:"db_password"`
	DBName             string `json:"db_name"`
	DBSchema           string `json:"db_schema"`
	DBPort             string `json:"db_port"`
	MongoConnectionURI string `json:"mongo_db_connection_URI"`
}

// ProdConfig holds the values for production environments
type ProdConfig struct {
	AppDomain          string `json:"app_domain"`
	AppPort            string `json:"app_port"`
	DBUserName         string `json:"db_username"`
	DBPassword         string `json:"db_password"`
	DBName             string `json:"db_name"`
	DBSchema           string `json:"db_schema"`
	DBPort             string `json:"db_port"`
	MongoConnectionURI string `json:"mongo_db_connection_URI"`
}

// TestConfig holds the values for production environments
type TestConfig struct {
	AppDomain          string `json:"app_domain"`
	AppPort            string `json:"app_port"`
	DBUserName         string `json:"db_username"`
	DBPassword         string `json:"db_password"`
	DBName             string `json:"db_name"`
	DBSchema           string `json:"db_schema"`
	DBPort             string `json:"db_port"`
	MongoConnectionURI string `json:"mongo_db_connection_URI"`
}
