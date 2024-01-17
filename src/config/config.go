package config

import (
	"errors"
	"os"

	"github.com/subosito/gotenv"
)

var (
	Api struct {
		Port string
	}
	Database struct {
		Host     string
		Port     string
		Name     string
		User     string
		Password string
	}
)

func LoadEnvironmentVariables(f string) (err error) {
	err = gotenv.Load(f)
	if err != nil {
		return err
	}

	if value, found := os.LookupEnv("API_PORT"); found {
		Api.Port = value
	} else {
		return errors.New("environment variable 'API_PORT' is required")
	}

	if value, found := os.LookupEnv("DATABASE_HOST"); found {
		Database.Host = value
	} else {
		return errors.New("environment variable 'DATABASE_HOST' is required")
	}

	if value, found := os.LookupEnv("DATABASE_PORT"); found {
		Database.Port = value
	} else {
		return errors.New("environment variable 'DATABASE_PORT' is required")
	}

	if value, found := os.LookupEnv("DATABASE_NAME"); found {
		Database.Name = value
	} else {
		return errors.New("environment variable 'DATABASE_NAME' is required")
	}

	if value, found := os.LookupEnv("DATABASE_USER"); found {
		Database.User = value
	} else {
		return errors.New("environment variable 'DATABASE_USER' is required")
	}

	if value, found := os.LookupEnv("DATABASE_PASSWORD"); found {
		Database.Password = value
	} else {
		return errors.New("environment variable 'DATABASE_PASSWORD' is required")
	}

	return nil
}
