package models

import "database/sql"

type Channel struct {
	TeamID      string
	ChannelID   string
	ChannelName string
}

func SaveChannel(channel *Channel) {
	insertStmt, err := db.Prepare("INSERT INTO channels(teamID, channelID, channelName) VALUES (?,?,?)")
	if err != nil {
		panic(err)
	}
	defer insertStmt.Close()
	insertStmt.Exec(channel.TeamID, channel.ChannelID, channel.ChannelName)
}

func GetChannel(teamId string, channelId string) *Channel {
	selectStmt, err := db.Prepare("SELECT teamID, channelID, channelName FROM channels WHERE teamID = ? AND channelID = ?")
	if err != nil {
		panic(err)
	}
	defer selectStmt.Close()
	channel := &Channel{}
	err = selectStmt.QueryRow(teamId, channelId).Scan(&channel.TeamID, &channel.ChannelID, &channel.ChannelName)
	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		panic(err)
	default:
	}
	return channel
}

func UpdateChannel(channel *Channel) {
	updateChannelStmt, err := db.Prepare("UPDATE channels SET channelName = ? WHERE teamID = ? AND channelID = ?")
	if err != nil {
		panic(err)
	}
	defer updateChannelStmt.Close()
	updateChannelStmt.Exec(channel.ChannelName, channel.TeamID, channel.ChannelID)

	updateMessagesStmt, err := db.Prepare("UPDATE messages SET channelName = ? WHERE teamID = ? AND channelID = ?")
	if err != nil {
		panic(err)
	}
	defer updateMessagesStmt.Close()
	updateMessagesStmt.Exec(channel.ChannelName, channel.TeamID, channel.ChannelID)
}

func GetChannels(teamId string) []Channel {
	selectStmt, err := db.Prepare("SELECT teamId, channelID, channelName FROM channels WHERE teamID = ?")
	if err != nil {
		panic(err)
	}
	defer selectStmt.Close()
	rows, selectErr := selectStmt.Query(teamId)
	var channels []Channel

	switch {
	case selectErr == sql.ErrNoRows:
		return channels
	case selectErr != nil:
		panic(err)
	default:
	}

	for rows.Next() {
		channel := Channel{}
		rows.Scan(&channel.TeamID, &channel.ChannelID, &channel.ChannelName)
		channels = append(channels, channel)
	}

	return channels
}
