package configuration

import (
	"os"
)

type DatabaseConfiguration struct {
	Host         string `yaml:"host" json:"host"`
	Port         string `yaml:"port" json:"port"`
	DatabaseName string `yaml:"databaseName" json:"databaseName"`
	Username     string `yaml:"username" json:"username"`
	Password     string `yaml:"password" json:"password"`
}

func newDatabaseConfiguration(host string, port string, databaseName string, username string, password string) *DatabaseConfiguration {
	return &DatabaseConfiguration{
		Host:         host,
		Port:         port,
		DatabaseName: databaseName,
		Username:     username,
		Password:     password,
	}
}

func GetDatabaseConfiguration() *DatabaseConfiguration {
	// Might use ozzo-validation here, instead of manually checking each parameter
	host := os.Getenv("DB_HOST")

	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")

	if port == "" {
		port = "5432"
	}

	name := os.Getenv("DB_NAME")

	if name == "" {
		name = "postgres"
	}

	username := os.Getenv("DB_USERNAME")

	if username == "" {
		username = "postgres"
	}

	password := os.Getenv("DB_PASSWORD")

	if password == "" {
		password = ""
	}

	return newDatabaseConfiguration(host, port, name, username, password)
}
