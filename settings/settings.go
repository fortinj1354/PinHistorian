package settings

import (
	"os"
)

type Settings struct {
	SlackSecret         string
	SlackSigningVersion string
	SlackOAuth          string
	DatabaseName        string
	QueryAPIKey         string
	SentryDSN           string
	SentryEnvironment   string
}

var settings Settings

func LoadSettings() {
	settings = Settings{}

	if value, found := os.LookupEnv("SLACK_SECRET"); found {
		settings.SlackSecret = value
	} else {
		panic("No Slack OAuth")
	}

	if value, found := os.LookupEnv("SLACK_SIGNING_VERSION"); found {
		settings.SlackSigningVersion = value
	} else {
		settings.SlackSigningVersion = "v0"
	}

	if value, found := os.LookupEnv("SLACK_OAUTH"); found {
		settings.SlackOAuth = value
	} else {
		panic("No Slack OAuth")
	}

	if value, found := os.LookupEnv("QUERY_API_KEY"); found {
		settings.QueryAPIKey = value
	} else {
		panic("No Query API Key")
	}

	if value, found := os.LookupEnv("DATABASE_NAME"); found {
		settings.DatabaseName = value
	} else {
		settings.DatabaseName = "PinHistorian.sqlite"
	}

	if value, found := os.LookupEnv("SENTRY_DSN"); found {
		settings.SentryDSN = value
	}

	if value, found := os.LookupEnv("SENTRY_ENVIRONMENT"); found {
		settings.SentryEnvironment = value
	}
}

func GetSlackSecret() string {
	return settings.SlackSecret
}

func GetSlackSigningVersion() string {
	return settings.SlackSigningVersion
}

func GetSlackOAuth() string {
	return settings.SlackOAuth
}

func GetDatabaseName() string {
	return settings.DatabaseName
}

func GetQueryAPIKey() string {
	return settings.QueryAPIKey
}

func GetSentryDSN() string {
	return settings.SentryDSN
}

func GetSentryEnvironment() string {
	return settings.SentryEnvironment
}
