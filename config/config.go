package config

import (
	"os"

	"github.com/spf13/cast"
)

//Config ...
type Config struct {
	Environment     string // develop, staging, production
	LogLevel        string
	HTTPPort        string
	PostServiceHost string
	PostServicePort int
	FirstServiceHost string
	FirstServicePort int
	CtxTimeout      int
}

// Load loads environment vars and inflates Config
func Load() Config {
	c := Config{}

	c.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	c.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8000"))
	c.PostServiceHost = cast.ToString(getOrReturnDefault("POST_SERVICE_HOST", "localhost"))
	c.PostServicePort = cast.ToInt(getOrReturnDefault("POST_SERVICE_PORT", 8001))
	c.FirstServiceHost = cast.ToString(getOrReturnDefault("FIRST_SERVICE_HOST", "localhost"))
	c.FirstServicePort = cast.ToInt(getOrReturnDefault("FIRST_SERVICE_PORT", 8002))
	c.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	c.CtxTimeout = cast.ToInt(7)

	return c
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
