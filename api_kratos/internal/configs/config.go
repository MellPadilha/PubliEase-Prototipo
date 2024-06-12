package configs

import (
	"errors"

	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API      APIConfig
	Hannibal HannibalConfig
	EuRobo   EuRoboConfig
}

type APIConfig struct {
	Port string
}

type HannibalConfig struct {
	Address string
	Route   string
}

type EuRoboConfig struct {
	Address string
	Route   string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("hannibal.address", "127.0.0.1:8080")
	viper.SetDefault("hannibal.route", "/api/v1/form/")
	viper.SetDefault("eurobo.address", "127.0.0.1:7000")
	viper.SetDefault("eurobo.route", "/motivo")
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.Hannibal = HannibalConfig{
		Address: viper.GetString("hannibal.address"),
		Route:   viper.GetString("hannibal.route"),
	}

	cfg.EuRobo = EuRoboConfig{
		Address: viper.GetString("eurobo.address"),
		Route:   viper.GetString("eurobo.route"),
	}

	return nil
}

func GetServerPort() string {
	return cfg.API.Port
}

func GetHannibalHost() string {
	return cfg.Hannibal.Address
}

func GetHannibalRoute() string {
	return cfg.Hannibal.Route
}

func GetEuRoboHost() string {
	return cfg.EuRobo.Address
}

func GetEuRoboRoute() string {
	return cfg.EuRobo.Route
}
