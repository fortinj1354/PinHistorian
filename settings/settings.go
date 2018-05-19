package settings

import (
	"os"
)

type Settings struct {
	SlackToken   string
	SlackOAuth   string
	DatabaseName string
	QueryAPIKey  string
}

var settings Settings

func LoadSettings() {
	settings = Settings{}

	if value, found := os.LookupEnv("SLACK_TOKEN"); found {
		settings.SlackToken = value
	} else {
		panic("No Slack Token")
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
}

func GetSlackToken() string {
	return settings.SlackToken
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
