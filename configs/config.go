package config

import (
	"log"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseDriver   string `mapstructure:"DATABASE_DRIVER"`
	DatabaseHost     string `mapstructure:"DATABASE_HOST"`
	DatabasePort     string `mapstructure:"DATABASE_PORT"`
	DatabaseUser     string `mapstructure:"DATABASE_USER"`
	DatabasePassword string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseName     string `mapstructure:"DATABASE_NAME"`
	DatabaseSSLMode  string `mapstructure:"DATABASE_SSL_MODE"`
	WebServerPort    string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret        string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn     string `mapstructure:"JWT_EXPIRES_IN"`
	JWTAuth          *jwtauth.JWTAuth
}

var config Config

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading config file, %s", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}

	config.JWTAuth = jwtauth.New("HS256", []byte(config.JWTSecret), nil)

	return &config, nil
}

func BuildStringConnection() string {
	return "host=" + config.DatabaseHost +
		" port=" + config.DatabasePort +
		" user=" + config.DatabaseUser +
		" dbname=" + config.DatabaseName +
		" password=" + config.DatabasePassword +
		" sslmode=" + config.DatabaseSSLMode
}
