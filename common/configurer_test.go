package common

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/rs/zerolog"
)

func TestNewConfigurer(t *testing.T) {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Caller().Logger()
	t.Run("not found", func(t *testing.T) {
		configurer := NewConfigurer("config", "yaml", ".", &logger)
		if !cmp.Equal(configurer, defaultConfigurer) {
			t.Errorf("configurer should be equal to defaultConfigurer %v", defaultConfigurer)
		}
	})
	t.Run("success", func(t *testing.T) {
		configurer := NewConfigurer("config", "yaml", "../.", &logger)
		if configurer == nil {
			t.Error("configurer should not be nil")
		} else {
			failedConfigurer := &Configurer{
				ServerHost:               "",
				ServerPort:               "",
				DatabaseHost:             "",
				DatabasePort:             "",
				DatabaseUser:             "",
				DatabsePassword:          "",
				DatabaseConnectionString: "",
			}
			if cmp.Equal(configurer, failedConfigurer) {
				t.Errorf("configurer should not be equal to failedConfigurer %v", failedConfigurer)
			}
		}
	})
	t.Run("not read", func(t *testing.T) {
		configurer := NewConfigurer("config", "json", "../.", &logger)
		if !cmp.Equal(configurer, defaultConfigurer) {
			t.Errorf("configurer should be equal to defaultConfigurer %v", defaultConfigurer)
		}
	})
}
