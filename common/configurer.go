package common

import (
	"errors"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

var (
	ErrorFailedToReadConfig = errors.New("failed to read config")
	ErrorLoadDefaultConfig  = errors.New("failed to load config, load default config instead")
)

type Configurer struct {
	ServerHost               string
	ServerPort               string
	DatabaseHost             string
	DatabasePort             string
	DatabaseUser             string
	DatabsePassword          string
	DatabaseConnectionString string
}

func NewConfigurer(logger *zerolog.Logger) *Configurer {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../.")
	viper.AddConfigPath(".")
	logger.Info().Str("service", serviceName).Msgf("config file used: %v", viper.GetViper().ConfigFileUsed())
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Error().Str("service", serviceName).Err(err).Msg(ErrorLoadDefaultConfig.Error())
			return &Configurer{
				ServerHost:               "localhost",
				ServerPort:               "8081",
				DatabaseHost:             "localhost",
				DatabasePort:             "5433",
				DatabaseUser:             "postgres",
				DatabsePassword:          "postgres",
				DatabaseConnectionString: "postgres://postgres:postgres@localhost:5433/postgres",
			}
		} else {
			logger.Fatal().Str("service", serviceName).Err(err).Msg(ErrorFailedToReadConfig.Error())
		}
	}
	connectionString := "postgres://" + viper.GetString("database.user") + ":" + viper.GetString("database.password") + "@" + viper.GetString("database.host") + ":" + viper.GetString("database.port") + "/" + viper.GetString("database.name")
	logger.Info().Str("service", serviceName).Msg("configurer initialized")
	return &Configurer{
		ServerHost:               viper.GetString("server.host"),
		ServerPort:               viper.GetString("server.port"),
		DatabaseHost:             viper.GetString("database.host"),
		DatabasePort:             viper.GetString("database.port"),
		DatabaseUser:             viper.GetString("database.user"),
		DatabsePassword:          viper.GetString("database.password"),
		DatabaseConnectionString: connectionString,
	}
}
