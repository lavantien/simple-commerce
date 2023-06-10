package main

import "simple-commerce/common"

const (
	serviceName    = "server"
	configFileName = "config"
	configFileType = "yaml"
	configFilePath = "../."
)

func main() {
	logger := common.GetLogger()
	configurer := common.NewConfigurer(configFileName, configFileType, configFilePath, logger)
	logger.Info().Str("service", serviceName).Msgf("server starting with config: %#v", configurer)
}
