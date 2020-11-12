# Pin Historian

Records and stores pins for Slack, currently only supports one server per instance.

## Running

Environment Variables:

- SLACK_SECRET - The signing secret for Slack HMAC verification
- SLACK_SIGNING_VERSION - (Optional) The version number for Slack HMAC verification, defaults to `v0`
- SLACK_OAUTH - OAuth bearer token for using the Slack API
- QUERY_API_KEY - Bearer token for querying stored pins
- DATABASE_NAME - (Optional) Name for the SQLite database file, defaults to `PinHistorian.sqlite`

Sentry Configuration:

Leaving both environment variables blank will disable Sentry. To enable Sentry, provide values for both variables below.

- SENTRY_DSN - Endpoint for [Sentry](https://sentry.io/) error reporting
- SENTRY_ENVIRONMENT - Environment for Sentry error reporting

## Slack App Configuration

Set up the following event subscriptions to point to the root URL that the app is running at.

The bot only receives pin events for channels it has been added to, and will ignore any non-public channels.

Bot Event Subscriptions:

- channel_rename
- pin_added
- user_change

Bot Token OAuth Scopes:

- channels:read
- pins:read
- users.profile:read
- users:read