package common

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
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

var (
	TextFailedToReadConfig = "failed to read config, load default config instead"
	TextLoadDefaultConfig  = "failed to load config, load default config instead"
	defaultConfigurer      = &Configurer{
		ServerHost:               "localhost",
		ServerPort:               "8081",
		DatabaseHost:             "localhost",
		DatabasePort:             "5433",
		DatabaseUser:             "postgres",
		DatabsePassword:          "postgres",
		DatabaseConnectionString: "postgres://postgres:postgres@localhost:5433/postgres",
	}
)

func NewConfigurer(fileName string, fileType string, filePath string, logger *zerolog.Logger) *Configurer {
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	viper.AddConfigPath(filePath)
	logger.Info().Str("service", serviceName).Msgf("config file used: %v", viper.GetViper().ConfigFileUsed())
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			logger.Error().Str("service", serviceName).Err(err).Msg(TextLoadDefaultConfig)
			return defaultConfigurer
		} else {
			logger.Error().Str("service", serviceName).Err(err).Msg(TextFailedToReadConfig)
			return defaultConfigurer
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
