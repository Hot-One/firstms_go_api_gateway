package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"

	TimeExpiredAt = time.Hour * 720
)

type Config struct {
	ServiceName string
	Environment string // debug, test, release
	Version     string

	ServiceHost string
	HTTPPort    string
	HTTPScheme  string
	Domain      string

	DefaultOffset string
	DefaultLimit  string

	// firstms_go_order_service host and port for connection to service
	OrderServiceHost string
	OrderGRPCPort    string

	PostgresMaxConnections int32

	SecretKey string
}

// Load ...
func Load() Config {
	if err := godotenv.Load("./.env"); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "ibron_go_user_service"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", ReleaseMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.ServiceHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "localhost"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8001"))
	config.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_Scheme", "http"))
	config.Domain = cast.ToString(getOrReturnDefaultValue("DOMAIN", "localhost:8001"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("LIMIT", "10"))

	config.OrderServiceHost = cast.ToString(getOrReturnDefaultValue("USER_SERVICE_HOST", "localhost"))
	config.OrderGRPCPort = cast.ToString(getOrReturnDefaultValue("USER_GRPC_PORT", ":9101"))

	config.SecretKey = cast.ToString(getOrReturnDefaultValue("SECRET_KEY", "ibron"))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
