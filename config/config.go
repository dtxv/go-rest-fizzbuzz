package config

import (
	"fmt"
	"strconv"

	"github.com/rs/zerolog/log"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

// utils constant
const (
	Colon = ":"
	Fmt3  = "%s%s%s"
)

// AppConfig struct
type AppConfig struct {
	HTTPPort     int
	HostAddress  string
	LoggerConfig *LoggerConfig
}

// Key string enum
const (
	hostAddress  string = "HOST_ADDRESS"
	hostPort     string = "HTTP_PORT"
	loggingJSON  string = "LOGGING_JSON"
	loggingLevel string = "LOGGING_LEVEL"
)

// GlobalConfig global object
var GlobalConfig *AppConfig

// Load env vars
func LoadFromEnv() *AppConfig {
	viper.AutomaticEnv()
	viper.SetDefault(loggingLevel, zerolog.DebugLevel)
	viper.SetDefault(hostAddress, "0.0.0.0")
	viper.SetDefault(hostPort, "5000")

	loggingLevel, err := zerolog.ParseLevel(viper.GetString(loggingLevel))
	if err != nil || loggingLevel == zerolog.NoLevel {
		loggingLevel = zerolog.InfoLevel
	}

	return &AppConfig{
		HostAddress: viper.GetString(hostAddress),
		HTTPPort:    viper.GetInt(hostPort),

		LoggerConfig: &LoggerConfig{
			Level: loggingLevel,
			JSON:  viper.GetBool(loggingJSON),
		},
	}
}

// Host url stringifier
func (c *AppConfig) HostUrl() string {
	return fmt.Sprintf(Fmt3, c.HostAddress, Colon, strconv.Itoa(c.HTTPPort))
}

// Print configurations
func (c *AppConfig) PrintConfig() {
	log.
		Debug().
		Interface("config", c).
		Msg("Configuration")
}

// Init configurations
func init() {
	GlobalConfig = LoadFromEnv()
}
