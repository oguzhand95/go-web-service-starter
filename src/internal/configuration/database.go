package configuration

import (
	"flag"
	validation "github.com/go-ozzo/ozzo-validation"
)

type DatabaseConfiguration struct {
	Host         *string `yaml:"host" json:"host"`
	Port         *string `yaml:"port" json:"port"`
	DatabaseName *string `yaml:"databaseName" json:"databaseName"`
	Username     *string `yaml:"username" json:"username"`
	Password     *string `yaml:"password" json:"password"`
}

func (dc DatabaseConfiguration) Validate() error {
	return validation.ValidateStruct(&dc,
		validation.Field(&dc.Host, validation.NilOrNotEmpty),
		validation.Field(&dc.Port, validation.NilOrNotEmpty),
		validation.Field(&dc.DatabaseName, validation.NilOrNotEmpty),
		validation.Field(&dc.Username, validation.NilOrNotEmpty),
		validation.Field(&dc.Password, validation.NilOrNotEmpty),
	)
}

func newDatabaseConfiguration(host *string, port *string, databaseName *string, username *string, password *string) *DatabaseConfiguration {
	return &DatabaseConfiguration{
		Host:         host,
		Port:         port,
		DatabaseName: databaseName,
		Username:     username,
		Password:     password,
	}
}

func GetDatabaseConfiguration() *DatabaseConfiguration {
	host := flag.String("db-host", "localhost", "postgres database host")
	port := flag.String("db-port", "5432", "postgres database port")
	name := flag.String("db-name", "postgres", "postgres database name")
	username := flag.String("db-username", "postgres", "postgres database username")
	password := flag.String("db-password", "", "postgres database password")

	return newDatabaseConfiguration(host, port, name, username, password)
}
