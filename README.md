# Pin Historian

Records and stores pins for Slack. Currently only supports one server per instance.

## Running

Required Environment Variables:

- SLACK_TOKEN - The verification token on Slack's POST requests
- SLACK_OAUTH - OAuth bearer token for using the Slack API
- QUERY_API_KEY - Bearer token for querying stored pins
- SENTRY_DSN - Endpoint for [Sentry](https://sentry.io/) error reporting
- SENTRY_ENVIRONMENT - (Optional) Environment for Sentry error reporting
- DATABASE_NAME - (Optional) Name for the SQLite database file, defaults to PinHistorian.sqlite if not specified

## Slack App Configuration

Set up the following event subscriptions to point to the root URL that the app is running at.

Event Subscriptions:

- channel_rename
- pin_added
- user_change

OAuth Scopes:

- channels:read
- pins:read
- users.profile:read
- users:read