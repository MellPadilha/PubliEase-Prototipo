package config

import (
	"github.com/spf13/viper"
)

// setupViper initializes Viper for configuration management.
// Returns an error if reading the configuration file fails.
func setupViper() error {
	viper.SetConfigFile("./config.yaml")
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

// getConfigValues populates the provided Configuration struct with values from Viper.
// It unmarshals the configuration data into the Configuration struct.
// Returns an error if unmarshaling fails.
func getConfigValues(configuration *Configuration) error {
	err := viper.Unmarshal(&configuration)
	if err != nil {
		return err
	}

	return nil
}
