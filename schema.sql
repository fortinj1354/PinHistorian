CREATE TABLE `messages` (
  `eventID`     TEXT PRIMARY KEY,
  `teamID`      TEXT      NOT NULL,
  `channelID`   TEXT      NOT NULL,
  `channelName` TEXT      NOT NULL,
  `userID`      TEXT      NOT NULL,
  `userDisplay` TEXT      NOT NULL,
  `messageText` TEXT      NOT NULL,
  `messageTime` TIMESTAMP NOT NULL
);

CREATE INDEX team_channel_time
  ON `messages` (`teamId`, `channelId`, `messageTime`);

CREATE TABLE `users` (
  `teamID`      TEXT NOT NULL,
  `userID`      TEXT NOT NULL,
  `userDisplay` TEXT NOT NULL,
  PRIMARY KEY (`teamID`, `userID`)
);

CREATE TABLE `channels` (
  `teamID`      TEXT NOT NULL,
  `channelID`   TEXT NOT NULL,
  `channelName` TEXT NOT NULL,
  PRIMARY KEY (`teamID`, `channelID`)
);