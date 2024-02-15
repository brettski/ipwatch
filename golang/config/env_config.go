package config

import (
	"fmt"
	"log"
	"os"
	"strings"

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
	// postmark
	config.Postmark.Token = getEnvValue("POSTMARK_TOKEN", false)
	config.Postmark.EmailFrom = getEnvValue("EMAIL_FROM", false)
	config.Postmark.EmailTo = getEnvValue("EMAIL_TO", false)
	config.Postmark.TestLocation = getEnvValue("TEST_LOCATION", false)

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

// Validate Postmark values are set
func (appConfig IpWatchConfig) ValidatePostmark() (bool, string) {
	log.Print("Validating Postmark values")
	var issues []string
	if appConfig.Postmark.Token == "" {
		issues = append(issues, "POSTMARK_TOKEN")
	}
	if appConfig.Postmark.EmailFrom == "" {
		issues = append(issues, "EMAIL_FROM")
	}
	if appConfig.Postmark.EmailTo == "" {
		issues = append(issues, "EMAIL_TO")
	}

	if len(issues) > 0 {
		out := strings.Join(issues, `, `)
		return false, out
	}

	return true, ""
}
