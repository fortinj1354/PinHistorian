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
	selectStmt, err := db.Prepare("SELECT teamId, channelId, channelName FROM channels WHERE teamId = ? AND channelID = ?")
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
