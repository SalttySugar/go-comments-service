package service

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// global singleton variable
var config *Config

// global configuration object
type Config struct {
	Server   *ServerConfig
	Database *DatabaseConfiguration
	Caching cache.Config
}

// application entry point configuration object
type AppConfig struct {
	Server      *ServerConfig
	Controllers *[]Controller
	Middlewares *[]fiber.Handler
}

// server configuration object
type ServerConfig struct {
	Host  string
	Port  int
	Debug bool
}

// ==== database configuration START ==== //
type DatabaseDriverType string

const (
	POSTGRES DatabaseDriverType = "postgresql"
	MARIADB  DatabaseDriverType = "mariadb"
	SQLITE   DatabaseDriverType = "sqlite"
	MYSQL    DatabaseDriverType = "mysql"
)

type DatabaseConnectionConfiguration struct {
	Open     int
	Idle     int
	Lifetime time.Duration
}

type DatabaseConfiguration struct {
	Driver     DatabaseDriverType
	Name       string
	User       string
	Password   string
	Host       string
	Port       int
	Path       string
	Connection *DatabaseConnectionConfiguration
}

// ==== database configuration END ==== //

func LoadConfig() *Config {
	if config != nil {
		return config
	}
	var file string

	// flags
	pflag.String("file", "", "--file <file_name>.yaml|yml")

	// defaults
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("server.debug", true)
	viper.SetDefault("database.driver", "sqlite")
	viper.SetDefault("database.path", "db.sqlite")
	viper.SetDefault("database.connection.open", 24)
	viper.SetDefault("database.connection.idle", 10)
	viper.SetDefault("database.connection.lifetime", time.Hour)

	viper.AutomaticEnv()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}

	file = viper.Get("file").(string)

	if len(file) != 0 {
		viper.SetConfigFile(file)
		if err := viper.ReadInConfig(); err != nil {
			panic("Error reading configuration file: " + err.Error())
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}

	return config
}
