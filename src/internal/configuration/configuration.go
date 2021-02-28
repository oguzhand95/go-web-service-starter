package configuration

import "flag"

type Configuration struct {
	DatabaseConfiguration *DatabaseConfiguration
	SessionConfiguration *SessionConfiguration
}

func NewConfiguration() (*Configuration, error) {
	var err error

	configuration := &Configuration{
		DatabaseConfiguration: nil,
		SessionConfiguration: nil,
	}

	databaseConfiguration := GetDatabaseConfiguration()
	sessionConfiguration := GetSessionConfiguration()

	configuration.DatabaseConfiguration = databaseConfiguration
	configuration.SessionConfiguration = sessionConfiguration

	flag.Parse()

	if err = databaseConfiguration.Validate(); err != nil {
		return nil, err
	}

	if err = sessionConfiguration.Validate(); err != nil {
		return nil, err
	}

	return configuration, nil
}