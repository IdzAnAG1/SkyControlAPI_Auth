package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type AuthServerConfig struct {
	Port string `env:"AUTH_SERVER_PORT" envDefault:"8080"`
}
type GatewayConfig struct {
	Port    string `env:"GATEWAY_PORT" envDefault:"8081"`
	Timeout int    `env:"GATEWAY_TIMEOUT" envDefault:"5000"`
	NoCA    int    `env:"GATEWAY_NUMBER_OF_CONNECTION_ATTEMPTS" envDefault:"5"`
}

type DatabaseConfig struct {
	Host     string `env:"AUTH_DATABASE_HOST,required"`
	Port     string `env:"AUTH_DATABASE_PORT" envDefault:"5432"`
	Name     string `env:"AUTH_DATABASE_NAME" envDefault:"auth"`
	User     string `env:"AUTH_DATABASE_USER" envDefault:"postgres"`
	Password string `env:"AUTH_DATABASE_PASSWORD,required"`
	Timeout  int    `env:"AUTH_DATABASE_TIMEOUT" envDefault:"5000"`
	NoCA     int    `env:"AUTH_DATABASE_NUMBER_OF_CONNECTION_ATTEMPTS" envDefault:"5"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST,required"`
	Port     string `env:"REDIS_PORT" envDefault:"6379"`
	Name     string `env:"REDIS_NAME,required"`
	User     string `env:"REDIS_USER,required"`
	Password string `env:"REDIS_PASSWORD,required"`
	Timeout  int    `env:"REDIS_TIMEOUT" envDefault:"5000"`
	NoCA     int    `env:"REDIS_NUMBER_OF_CONNECTION_ATTEMPTS" envDefault:"5"`
}
type Config struct {
	// Configuration for Auth server
	AuthServer AuthServerConfig
	// Configuration for relation with Gateway
	Gateway GatewayConfig
	// Configuration for relation with Database
	Database DatabaseConfig
	// Configuration for relation with Redis
	Redis RedisConfig
}

func LoadAndGetConfig(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		fmt.Printf(
			"Error loading env file: %v \n"+
				"Trying to read environment variables\n", err,
		)
	}

	config := &Config{}

	err = env.Parse(config)
	if err != nil {
		fmt.Println("Failed to parse env variables")
		return nil, err
	}

	return config, nil
}
