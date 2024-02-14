package config

import (
	"fmt"
	"log"
	"os"

	dotenv "github.com/joho/godotenv"
)

type IpWatchConfig struct {
	CheckEndpoint string
	Postmark      envPostmark
}

type envPostmark struct {
	Token        string
	EmailTo      string
	EmailFrom    string
	TestLocation string
}

func GetConfig() IpWatchConfig {
	err := dotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file. Err: %s", err)
	}
	config := IpWatchConfig{}
	config.CheckEndpoint = getEnvValue("ENDPOINT_CHK", true)

	return config
}

func getEnvValue(envVar string, isRequired bool) string {
	value := os.Getenv(envVar)

	if value == "" && isRequired {
		log.Fatalf("value %s is required and not found", envVar)
	}

	return value
}

func DumpConfigToStdOut() {
	fmt.Printf("ipwatch current config:\n%+v\n", GetConfig())
}
