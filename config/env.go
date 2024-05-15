package config

import (
	"fmt"

	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
}

var Envs = initCong() //this creates a singleton

func initCong() Config {
	godotenv.Load() //Load environ variables into context
	godotenv.Read()
	return Config{
		PublicHost: getEnvProperty("PUBLIC_HOST", "http://localhost"),
		Port:       getEnvProperty("PUBLIC_PORT", "8080"),
		DBUser:     getEnvProperty("DB_USER", "root"),
		DBPassword: getEnvProperty("DB_PASSWORD", "password2020"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnvProperty("DB_HOST", "127.0.0.1"), getEnvProperty("DB_PORT", "3306")),
		DBName:     getEnvProperty("DB_NAME", "bigshop"),
	}

}

func getEnvProperty(key string, fallBack string) string {
	value, ok := os.LookupEnv(key)
	if ok {
		return value
	}
	return fallBack

}
