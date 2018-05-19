#Pin Historian

Records and stores pins for Slack. Currently only supports one server per instance.

##Running

Required Environment Variables:

- SLACK_TOKEN - The verification token on Slack's POST requests
- SLACK_OAUTH - OAuth bearer token for using the Slack API
- QUERY_API_KEY - Bearer token for querying stored pins
- DATABASE_NAME - Name for the SQLite database file, defaults to PinHistorian.sqlite if not specified