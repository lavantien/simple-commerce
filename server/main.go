package main

import "simple-commerce/common"

const (
	serviceName = "server"
)

func main() {
	logger := common.GetLogger()
	configurer := common.NewConfigurer(logger)
	logger.Info().Str("service", serviceName).Msgf("server starting with config: %#v", configurer)
}
