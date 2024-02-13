package main

import (
	"fmt"
	"log"
	"os"

	dotenv "github.com/joho/godotenv"
)

type envConfig struct {
	checkEndpoint string
	postmark      envPostmark
}

type envPostmark struct {
	token        string
	emailTo      string
	emailFrom    string
	testLocation string
}

func getEnvConfig() envConfig {
	err := dotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file. Err: %s", err)
	}
	config := envConfig{}
	config.checkEndpoint = getEnvValue("ENDPOINT_CHK", true)

	return config
}

func getEnvValue(envVar string, isRequired bool) string {
	value := os.Getenv(envVar)

	if value == "" && isRequired {
		log.Fatalf("value %s is required and not found", envVar)
	}

	return value
}

func dumpConfigToStdOut() {
	fmt.Printf("%+v", getEnvConfig())
}
